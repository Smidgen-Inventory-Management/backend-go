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

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	api "smidgen-backend/src/api"
	models "smidgen-backend/src/models"
	service "smidgen-backend/src/services"
	utils "smidgen-backend/src/utils"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var log = utils.Log()

func main() {
	_, err := os.Stat("configs")
	dirExists := !os.IsNotExist(err)
	argsLengthMoreThanTwo := (len(os.Args) > 2)

	switch {
	case !dirExists || argsLengthMoreThanTwo:
		log.Error("Path to configuration files not found. Ensure you have a \"configs\" directory, or run the server with the appropriate arguments.")
		log.Error("Usage: go run main.go <path_to_server_configurations> <path_to_database_configurations>")
		return
	case argsLengthMoreThanTwo:
		utils.ServerConfigPath = os.Args[1]
		utils.DatabaseConfigPath = os.Args[2]
	default:
		utils.ServerConfigPath = "configs/server.yaml"
		utils.DatabaseConfigPath = "configs/db_conn.yaml"
	}

	serverConfig, err := LoadServerConfig(utils.ServerConfigPath)
	if err != nil {
		log.Errorf("Failed to load server config: %v", err)
	} else {
		log.Info("Loaded server configurations.")
	}

	envConfig, ok := serverConfig.Environments["Development"]
	if !ok {
		log.Fatal("Environment configuration not found")
	} else {
		log.Info("Loaded server environment configurations.")
	}

	log = utils.Log(envConfig.Debug)

	if err := retryDatabaseConnection(utils.DatabaseConfigPath); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Debug("Database connection successful!")
	}

	hostname := envConfig.Host + ":" + envConfig.Port
	router := loadRoutes(envConfig)

	log.Debug("Routes loaded.")
	log.Infof("Server starting on %s", hostname)
	server := &http.Server{
		Addr:         hostname,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Errorf("Failed to start server: %v", err)
	}
}

func retryDatabaseConnection(configPath string) error {
	const maxRetries = 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		if err := CheckDatabaseConnection(configPath); err != nil {
			if attempt < maxRetries {
				log.Warnf("Failed to connect to database. Retrying in 3 seconds... (Attempt %d/%d)", attempt, maxRetries)
				time.Sleep(3 * time.Second)
				continue
			}
			return fmt.Errorf(fmt.Sprintf("failed to connect to database after %d attempts. Verify the database is running then relaunch the server", maxRetries))
		}
		return nil
	}
	return nil
}

func CheckDatabaseConnection(configPath string) error {
	db, err := utils.NewDatabaseConnection(configPath, "read")
	if err != nil {
		return fmt.Errorf("failed to create database connection: %v", err)
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}

func LoadServerConfig(yamlFilePath string) (models.ServerConfig, error) {
	yamlFile, err := os.Open(yamlFilePath)
	if err != nil {
		return models.ServerConfig{}, fmt.Errorf("\nfailed to open YAML file: %v", err)
	}
	defer yamlFile.Close()

	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return models.ServerConfig{}, fmt.Errorf("\nfailed to read YAML file: %v", err)
	}
	var config models.ServerConfig
	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		return models.ServerConfig{}, fmt.Errorf("\nfailed to unmarshal YAML: %v", err)
	}

	return config, nil
}

func loadRoutes(environmentConfig struct {
	Host     string "yaml:\"host\""
	Port     string "yaml:\"port\""
	Debug    bool   "yaml:\"debug\""
	RootPath string "yaml:\"root_path\""
}) *mux.Router {

	DefaultAPIService := service.NewDefaultAPIService()
	BusinessUnitAPIService := service.NewBusinessUnitAPIService()
	EquipmentAPIService := service.NewEquipmentAPIService()
	EquipmentAssignmentAPIService := service.NewEquipmentAssignmentAPIService()
	UserAPIService := service.NewUserAPIService()
	log.Debug("loaded API cervices")

	DefaultAPIController := api.NewDefaultAPIController(DefaultAPIService)
	BusinessUnitAPIController := api.NewBusinessUnitAPIController(BusinessUnitAPIService)
	EquipmentAPIController := api.NewEquipmentAPIController(EquipmentAPIService)
	EquipmentAssignmentAPIController := api.NewEquipmentAssignmentAPIController(EquipmentAssignmentAPIService)
	UserAPIController := api.NewUserAPIController(UserAPIService)
	log.Debug("loaded API controllers")

	router := utils.NewRouter(environmentConfig.RootPath, BusinessUnitAPIController, DefaultAPIController, EquipmentAPIController, EquipmentAssignmentAPIController, UserAPIController)
	log.Debug("successfully created routers")
	return router
}
