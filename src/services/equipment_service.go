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
	"strconv"
	"time"
)

// EquipmentAPIService is a service that implements the logic for the EquipmentAPIServicer
// This service should implement the business logic for every endpoint for the EquipmentAPI API.
// Include any external packages or services that will be required by this service.
type EquipmentAPIService struct {
}

// NewEquipmentAPIService creates a default api service
func NewEquipmentAPIService() api.EquipmentAPIServicer {
	return &EquipmentAPIService{}
}

// AddEquipment - Create equipment
func (s *EquipmentAPIService) AddEquipment(ctx context.Context, equipment models.Equipment) (utils.ImplResponse, error) {
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
		Action:          "ADD_EQUIPMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("equipment", equipment)
	if err != nil {
		if err.Error() == "23503" {
			logConnection.InsertRow("audit_log", logEntry)
			log.Error("unable to insert new data. The data requires reference to another table (foreign key constraint)")
			return utils.Response(400, nil), errors.New("the BusinessUnit with ID of " + strconv.Itoa(int(equipment.BusinessUnitId)) + " does not exist")
		}
		logConnection.InsertRow("audit_log", logEntry)
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}

// DeleteEquipment - Delete equipment
func (s *EquipmentAPIService) DeleteEquipment(ctx context.Context, equipmentId int32) (utils.ImplResponse, error) {
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
		Action:          "DELETE_EQUIPMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("equipment", "EquipmentId", equipmentId)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Error: %v", err)
		return utils.Response(404, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, nil), nil
}

// GetEquipments - Get equipments
func (s *EquipmentAPIService) GetEquipments(ctx context.Context) (utils.ImplResponse, error) {
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
		Action:          "GET_EQUIPMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.Equipment
	rows, err := dbConnection.GetRows("equipment", &dest)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Error: %v", err)
	}
	if len(rows) == 0 {
		logConnection.InsertRow("audit_log", logEntry)
		return utils.Response(200, nil), fmt.Errorf("no equipment was found in the database")
	}
	var Assets []models.Equipment
	for _, row := range rows {
		equipment, ok := row.(models.Equipment)
		if !ok {
			logEntry.ActionStatus = "WARN"
			logConnection.InsertRow("audit_log", logEntry)
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		Assets = append(Assets, equipment)
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, Assets), nil
}

// GetEquipmentById - Get equipment
func (s *EquipmentAPIService) GetEquipmentById(ctx context.Context, equipmentId int32) (utils.ImplResponse, error) {
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
		Action:          "GET_EQUIPMENT_BY_ID",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}
	var dest models.Equipment
	row, err := dbConnection.GetByID("equipment", "equipmentId", equipmentId, &dest)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	equipment, ok := row.(models.Equipment)
	if !ok {
		logEntry.ActionStatus = "WARN"
		logConnection.InsertRow("audit_log", logEntry)
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, equipment), nil
}

// UpdateEquipment - Update equipment
func (s *EquipmentAPIService) UpdateEquipment(ctx context.Context, equipmentId int32, equipment models.Equipment) (utils.ImplResponse, error) {
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
		Action:          "UPDATE_EQUIPMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("equipment", "equipmentId", equipmentId, equipment)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Error(err)
		return utils.Response(400, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}
