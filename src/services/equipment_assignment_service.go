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

// EquipmentAssignmentAPIService is a service that implements the logic for the EquipmentAssignmentAPIServicer
// This service should implement the business logic for every endpoint for the EquipmentAssignmentAPI API.
// Include any external packages or services that will be required by this service.
type EquipmentAssignmentAPIService struct {
}

// NewEquipmentAssignmentAPIService creates a default api service
func NewEquipmentAssignmentAPIService() api.EquipmentAssignmentAPIServicer {
	return &EquipmentAssignmentAPIService{}
}

// AddEquipmentAssignment - Create assignment
func (s *EquipmentAssignmentAPIService) AddEquipmentAssignment(ctx context.Context, equipmentAssignment models.EquipmentAssignment) (utils.ImplResponse, error) {
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
		Action:          "ADD_EQUIPMENT_ASSIGNMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("equipment_assignment", equipmentAssignment)
	if err != nil {
		log.Error(err)
		logConnection.InsertRow("audit_log", logEntry)
		return utils.Response(500, nil), errors.New("an error has occurred while adding new data")
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}

// DeleteEquipmentAssignment - Delete assignment
func (s *EquipmentAssignmentAPIService) DeleteEquipmentAssignment(ctx context.Context, assignmentId int32) (utils.ImplResponse, error) {
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
		Action:          "DELETE_EQUIPMENT_ASSIGNMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("equipment", "equipmentid", assignmentId)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Error: %v", err)
		return utils.Response(404, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, nil), nil
}

// GetEquipmentAssignments - Get assignments
func (s *EquipmentAssignmentAPIService) GetEquipmentAssignments(ctx context.Context) (utils.ImplResponse, error) {
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
		Action:          "GET_EQUIPMENT_ASSIGNMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.EquipmentAssignment
	rows, err := dbConnection.GetRows("equipment_assignment", &dest)

	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Error("Error: %v", err)
	}

	if len(rows) == 0 {
		logEntry.ActionStatus = "SUCCESS"
		logConnection.InsertRow("audit_log", logEntry)
		return utils.Response(404, nil), fmt.Errorf("no equipment assignments were found in the database")
	}

	var Assignments []models.EquipmentAssignment
	for _, row := range rows {
		Assignment, ok := row.(models.EquipmentAssignment)
		if !ok {
			logEntry.ActionStatus = "WARN"
			logConnection.InsertRow("audit_log", logEntry)
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		Assignments = append(Assignments, Assignment)
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, Assignments), nil
}

// GetEquipmentAssignmentById - Get assignment
func (s *EquipmentAssignmentAPIService) GetEquipmentAssignmentById(ctx context.Context, assignmentId int32) (utils.ImplResponse, error) {
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
		Action:          "GET_EQUIPMENT_ASSIGNMENT_BY_ID",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}
	var dest models.Equipment
	row, err := dbConnection.GetByID("equipment_assignment", "assignmentId", assignmentId, &dest)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	assignment, ok := row.(models.EquipmentAssignment)
	if !ok {
		logConnection.InsertRow("audit_log", logEntry)
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(200, assignment), nil
}

// UpdateEquipmentAssignment - Update assignment
func (s *EquipmentAssignmentAPIService) UpdateEquipmentAssignment(ctx context.Context, assignmentId int32, equipmentAssignment models.EquipmentAssignment) (utils.ImplResponse, error) {
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
		Action:          "UPDATE_EQUIPMENT_ASSIGNMENT",
	}
	logConnection, _ := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "write")
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("equipment_assignment", "assignmentid", assignmentId, equipmentAssignment)
	if err != nil {
		logConnection.InsertRow("audit_log", logEntry)
		log.Error(err)
		return utils.Response(400, nil), err
	}

	logEntry.ActionStatus = "SUCCESS"
	logConnection.InsertRow("audit_log", logEntry)
	return utils.Response(202, nil), nil
}
