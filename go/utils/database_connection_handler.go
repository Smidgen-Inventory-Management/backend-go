/*
 * Smidgen
 *
 * API for interacting with Smidgen.
 *
 *   Smidgen aims to simplify and automate common tasks that logisticians
 *   conduct on a daily basis so they can focus on the effective distribution
 *   of materiel, as well as maintain an accurate record keeping book of
 *   receiving, issuance, audits, surpluses, amongst other logistical tasks.
 *   Copyright (C) 2024  Jose Hernandez
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package smidgen

import (
	"database/sql"
	"fmt"
	"io"
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

var log = Log()

func NewDatabaseConnection(privilege string) (*DatabaseConnection, error) {
	instance := &DatabaseConnection{privilege: privilege}
	initErr := instance.initialize(privilege)
	if initErr != nil {
		return nil, fmt.Errorf("failed to initialize database connection: %v", initErr)
	}
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
	log.Info("Successfully loaded database configurations.")

	field := reflect.ValueOf(&config).Elem().FieldByName(strings.ToUpper(privilege[:1]) + privilege[1:])
	if !field.IsValid() {
		return fmt.Errorf("\ninvalid privilege level")
	}
	log.Info("Successfully loaded %v connection configurations.", privilege)
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

func (dao *DatabaseConnection) Ping() error {
	return dao.db.Ping()
}
