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
)

// BusinessUnitAPIService is a service that implements the logic for the BusinessUnitAPIServicer
// This service should implement the business logic for every endpoint for the BusinessUnitAPI API.
// Include any external packages or services that will be required by this service.
type BusinessUnitAPIService struct {
}

// NewBusinessUnitAPIService creates a default api service
func NewBusinessUnitAPIService() api.BusinessUnitAPIServicer {
	return &BusinessUnitAPIService{}
}

// AddUnitBusinessUnitPost - Create Business Unit
func (s *BusinessUnitAPIService) AddUnitBusinessUnitPost(ctx context.Context, businessUnit models.BusinessUnit) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("businessunit", businessUnit)
	if err != nil {
		fmt.Println(err)
		return utils.Response(500, nil), errors.New("an error has occured while adding new data")
	}
	return utils.Response(202, nil), nil
}

// DeleteUnitBusinessUnitUnitIdDelete - Delete Business Unit
func (s *BusinessUnitAPIService) DeleteUnitBusinessUnitUnitIdDelete(ctx context.Context, unitId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("businessUnit", "UnitID", unitId)
	if err != nil {
		fmt.Println("Error:", err)
		return utils.Response(500, nil), err
	}

	return utils.Response(200, nil), nil
}

// GetUnitBusinessUnitGet - Get Business Units
func (s *BusinessUnitAPIService) GetUnitBusinessUnitGet(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.BusinessUnit
	rows, err := dbConnection.GetRows("businessUnit", &dest)
	if err != nil {
		fmt.Println("Error:", err)

	}

	var businessUnits []models.BusinessUnit
	for _, row := range rows {
		businessUnit, ok := row.(models.BusinessUnit)
		if !ok {
			fmt.Println("Error: Unexpected type in row")
			continue
		}
		businessUnits = append(businessUnits, businessUnit)
	}

	return utils.Response(200, businessUnits), nil
}

// GetUnitsBusinessUnitUnitIdGet - Get Business Unit
func (s *BusinessUnitAPIService) GetUnitsBusinessUnitUnitIdGet(ctx context.Context, unitId int32) (utils.ImplResponse, error) {
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.BusinessUnit
	row, err := dbConnection.GetByID("businessUnit", "unitid", unitId, &dest)
	if err != nil {
		fmt.Println("Data Not Found:", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	unit, ok := row.(models.BusinessUnit)
	if !ok {
		fmt.Println("Error: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	return utils.Response(200, unit), nil
}

// UpdateUnitBusinessUnitUnitIdPut - Update Business Unit
func (s *BusinessUnitAPIService) UpdateUnitBusinessUnitUnitIdPut(ctx context.Context, unitId int32, businessUnit models.BusinessUnit) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("businessUnit", "unitid", unitId, businessUnit)
	if err != nil {
		fmt.Println(err)
		return utils.Response(500, nil), errors.New("an error has occured while updating the data")
	}
	return utils.Response(202, nil), nil
}
