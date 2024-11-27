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
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	api "smidgen-backend/src/api"
	models "smidgen-backend/src/models"
	utils "smidgen-backend/src/utils"
	"time"
)

// ManufacturerAPIService is a service that implements the logic for the ManufacturerAPIServicer
// This service should implement the business logic for every endpoint for the ManufacturerAPI API.
// Include any external packages or services that will be required by this service.
type ManufacturerAPIService struct {
}

// NewManufacturerAPIService creates a default api service
func NewManufacturerAPIService() api.ManufacturerAPIServicer {
	return &ManufacturerAPIService{}
}

// AddManufacturer - Create manufacturer
func (s *ManufacturerAPIService) AddManufacturer(ctx context.Context, manufacturer models.Manufacturer) (utils.ImplResponse, error) {
	privilege := "write"
	var uuid16 [2]byte
	_, err := rand.Read(uuid16[:])
	if err != nil {
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	uuid := int(binary.BigEndian.Uint16(uuid16[:]))

	logEntry := models.AuditLog{
		LogId:           uuid,
		ActionTimestamp: time.Now(),
		ActionStatus:    "Failed",
		Action:          "ADD_MANUFACTURER",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("manufacturers", manufacturer)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Error(err)
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}

// DeleteManufacturer - Delete manufacturer
func (s *ManufacturerAPIService) DeleteManufacturer(ctx context.Context, manufacturerId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	var uuid16 [2]byte

	_, err := rand.Read(uuid16[:])
	if err != nil {
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	uuid := int(binary.BigEndian.Uint16(uuid16[:]))

	logEntry := models.AuditLog{
		LogId:           uuid,
		ActionTimestamp: time.Now(),
		ActionStatus:    "Failed",
		Action:          "DELETE_MANUFACTURER",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("manufacturers", "ManufacturerID", manufacturerId)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Error: %v", err)
		return utils.Response(404, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, nil), nil
}

// GetManufacturers - Get manufacturers
func (s *ManufacturerAPIService) GetManufacturers(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	var uuid16 [2]byte

	_, err := rand.Read(uuid16[:])
	if err != nil {
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	uuid := int(binary.BigEndian.Uint16(uuid16[:]))

	logEntry := models.AuditLog{
		LogId:           uuid,
		ActionTimestamp: time.Now(),
		ActionStatus:    "Failed",
		Action:          "GET_MANUFACTURER",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.Manufacturer
	rows, err := dbConnection.GetRows("manufacturers", &dest)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Error: %v", err)
	}
	if len(rows) == 0 {
		logConnection.InsertRow("audit_log", logEntry)
		return utils.Response(200, nil), fmt.Errorf("no manufacturer was found in the database")
	}
	var Assets []models.Manufacturer
	for _, row := range rows {
		manufacturer, ok := row.(models.Manufacturer)
		if !ok {
			logEntry.ActionStatus = "WARN"
			logConnection.InsertRow("audit_log", logEntry)
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		Assets = append(Assets, manufacturer)
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, Assets), nil
}

// GetManufacturerById - Get manufacturer
func (s *ManufacturerAPIService) GetManufacturerById(ctx context.Context, manufacturerId int32) (utils.ImplResponse, error) {
	privilege := "read"
	var uuid16 [2]byte

	_, err := rand.Read(uuid16[:])
	if err != nil {
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	uuid := int(binary.BigEndian.Uint16(uuid16[:]))

	logEntry := models.AuditLog{
		LogId:           uuid,
		ActionTimestamp: time.Now(),
		ActionStatus:    "Failed",
		Action:          "GET_MANUFACTURER_BY_ID",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}
	var dest models.Manufacturer
	row, err := dbConnection.GetByID("manufacturers", "manufacturerId", manufacturerId, &dest)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	manufacturer, ok := row.(models.Manufacturer)
	if !ok {
		logEntry.ActionStatus = "WARN"
		logConnection.InsertRow("audit_log", logEntry)
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, manufacturer), nil
}

// UpdateManufacturer - Update manufacturer
func (s *ManufacturerAPIService) UpdateManufacturer(ctx context.Context, manufacturerId int32, manufacturer models.Manufacturer) (utils.ImplResponse, error) {
	privilege := "write"
	var uuid16 [2]byte

	_, err := rand.Read(uuid16[:])
	if err != nil {
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	uuid := int(binary.BigEndian.Uint16(uuid16[:]))

	logEntry := models.AuditLog{
		LogId:           uuid,
		ActionTimestamp: time.Now(),
		ActionStatus:    "Failed",
		Action:          "UPDATE_MANUFACTURER",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("manufacturers", "manufacturerId", manufacturerId, manufacturer)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Error(err)
		return utils.Response(400, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}
