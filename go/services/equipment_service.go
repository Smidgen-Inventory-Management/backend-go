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
	"errors"
	"fmt"
	"log"
	api "smidgen-backend/go/api"
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
	"strconv"
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

// AddEquipmentEquipmentPost - Create equipment
func (s *EquipmentAPIService) AddEquipmentEquipmentPost(ctx context.Context, equipment models.Equipment) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("equipment", equipment)
	if err != nil {
		if err.Error() == "23503" {
			log.Println("unable to insert new data. The data requires reference to another table (foreign key constraint)")
			return utils.Response(400, nil), errors.New("the BusinessUnit with ID of " + strconv.Itoa(int(equipment.BusinessUnitId)) + " does not exist")
		}
		return utils.Response(500, nil), errors.New("an error has occured while adding new data")
	}
	return utils.Response(202, nil), nil
}

// DeleteEquipmentEquipmentEquipmentIdDelete - Delete equipment
func (s *EquipmentAPIService) DeleteEquipmentEquipmentEquipmentIdDelete(ctx context.Context, equipmentId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("equipment", "EquipmentID", equipmentId)
	if err != nil {
		fmt.Println("Error:", err)
		return utils.Response(500, nil), err
	}

	return utils.Response(200, nil), nil
}

// GetEquipmentEquipmentGet - Get equipments
func (s *EquipmentAPIService) GetEquipmentEquipmentGet(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.Equipment
	rows, err := dbConnection.GetRows("equipment", &dest)
	if err != nil {
		fmt.Println("Error:", err)

	}

	var Assets []models.Equipment
	for _, row := range rows {
		equipment, ok := row.(models.Equipment)
		if !ok {
			fmt.Println("Error: Unexpected type in row")
			continue
		}
		Assets = append(Assets, equipment)
	}

	return utils.Response(200, Assets), nil
}

// GetEquipmentsEquipmentEquipmentIdGet - Get equipment
func (s *EquipmentAPIService) GetEquipmentsEquipmentEquipmentIdGet(ctx context.Context, equipmentId int32) (utils.ImplResponse, error) {
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.Equipment
	row, err := dbConnection.GetByID("equipment", "equipmentId", equipmentId, &dest)
	if err != nil {
		fmt.Println("Data Not Found:", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	equipment, ok := row.(models.Equipment)
	if !ok {
		fmt.Println("Error: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	return utils.Response(200, equipment), nil
}

// UpdateEquipmentEquipmentEquipmentIdPut - Update equipment
func (s *EquipmentAPIService) UpdateEquipmentEquipmentEquipmentIdPut(ctx context.Context, equipmentId int32, equipment models.Equipment) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("equipment", "equipmentid", equipmentId, equipment)
	if err != nil {
		fmt.Println(err)
		return utils.Response(500, nil), errors.New("an error has occured while updating the data")
	}
	return utils.Response(202, nil), nil
}
