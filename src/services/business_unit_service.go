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
	api "smidgen-backend/src/api"
	models "smidgen-backend/src/models"
	utils "smidgen-backend/src/utils"
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

// AddBusinessUnit - Create Business Unit
func (s *BusinessUnitAPIService) AddBusinessUnit(ctx context.Context, businessUnit models.BusinessUnit) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("businessunit", businessUnit)
	if err != nil {
		log.Error(err)
		return utils.Response(500, nil), errors.New("an error has occured while adding new data")
	}

	return utils.Response(202, nil), nil
}

// DeleteBusinessUnit - Delete Business Unit
func (s *BusinessUnitAPIService) DeleteBusinessUnit(ctx context.Context, unitId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("businessUnit", "UnitID", unitId)
	if err != nil {
		log.Errorf("Error: %v", err)
		return utils.Response(404, nil), err
	}

	return utils.Response(200, nil), nil
}

// GetBusinessUnits - Get Business Units
func (s *BusinessUnitAPIService) GetBusinessUnits(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.BusinessUnit
	rows, err := dbConnection.GetRows("businessUnit", &dest)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	if len(rows) == 0 {
		return utils.Response(404, nil), fmt.Errorf("no users were found in the database")
	}

	var businessUnits []models.BusinessUnit
	for _, row := range rows {
		businessUnit, ok := row.(models.BusinessUnit)
		if !ok {
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		businessUnits = append(businessUnits, businessUnit)
	}

	return utils.Response(200, businessUnits), nil
}

// GetBusinessUnitById - Get Business Unit
func (s *BusinessUnitAPIService) GetBusinessUnitById(ctx context.Context, unitId int32) (utils.ImplResponse, error) {
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.BusinessUnit
	row, err := dbConnection.GetByID("businessUnit", "unitid", unitId, &dest)
	if err != nil {
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	unit, ok := row.(models.BusinessUnit)
	if !ok {
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	return utils.Response(200, unit), nil
}

// UpdateB  usinessUnit - Update Business Unit
func (s *BusinessUnitAPIService) UpdateBusinessUnit(ctx context.Context, unitId int32, businessUnit models.BusinessUnit) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("businessUnit", "unitid", unitId, businessUnit)
	if err != nil {
		log.Error(err)
		return utils.Response(400, nil), err
	}
	return utils.Response(202, nil), nil
}
