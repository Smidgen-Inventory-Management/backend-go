package smidgen

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type DatabaseConnection struct {
	db        *sql.DB
	privilege string
	mu        sync.Mutex
}

type databaseConfig struct {
	Admin  databaseCredentials `yaml:"admin"`
	Read   databaseCredentials `yaml:"read"`
	Write  databaseCredentials `yaml:"write"`
	Delete databaseCredentials `yaml:"delete"`
}

type databaseCredentials struct {
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var (
	instance *DatabaseConnection
	once     sync.Once
)

func NewDatabaseConnection(privilege string) (*DatabaseConnection, error) {
	once.Do(func() {
		instance = &DatabaseConnection{privilege: privilege}
		if err := instance.initialize(privilege); err != nil {
			log.Fatalf("Failed to initialize database connection: %v", err)
		}
	})
	return instance, nil
}

func (dao *DatabaseConnection) initialize(privilege string) error {
	yamlFilePath := os.Args[2]
	yamlFile, err := os.Open(yamlFilePath)
	if err != nil {
		return fmt.Errorf("\nfailed to open YAML file: %v", err)
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return fmt.Errorf("\nfailed to read YAML file: %v", err)
	}
	var config databaseConfig

	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		return fmt.Errorf("\nfailed to unmarshal YAML: %v", err)
	}

	field := reflect.ValueOf(&config).Elem().FieldByName(strings.ToUpper(privilege[:1]) + privilege[1:])
	if !field.IsValid() {
		return fmt.Errorf("\ninvalid privilege level")
	}
	connectionConfig := field.Interface().(databaseCredentials)

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connectionConfig.Url, connectionConfig.Port, connectionConfig.User, connectionConfig.Password, connectionConfig.Database)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("\nfailed to open database connection: %v", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("\nfailed to ping database: %v", err)
	}
	dao.db = db
	return nil
}

func (dao *DatabaseConnection) Close() error {
	dao.mu.Lock()
	defer dao.mu.Unlock()
	if dao.db != nil {
		if err := dao.db.Close(); err != nil {
			return fmt.Errorf("\nfailed to close database connection: %v", err)
		}
		dao.db = nil
	}
	return nil
}
