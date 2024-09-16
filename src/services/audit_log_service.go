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
 *   along with this program. If not, see <https://www.gnu.org/licenses/>.
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

// AuditLogAPIService is a service that implements the logic for the AuditLogAPIServicer
// This service should implement the business logic for every endpoint for the AuditLogAPI API.
// Include any external packages or services that will be required by this service.
type AuditLogAPIService struct {
}

// NewAuditLogAPIService creates a default api service
func NewAuditLogAPIService() api.AuditLogAPIServicer {
	return &AuditLogAPIService{}
}

// GetAuditLogs - Get Audit Log
func (s *AuditLogAPIService) GetAuditLogs(ctx context.Context) (utils.ImplResponse, error) {
	privilege := "read"
	
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.AuditLog
	rows, err := dbConnection.GetRows("AuditLog", &dest)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	if len(rows) == 0 {
		return utils.Response(404, nil), fmt.Errorf("no logs were found in the database")
	}

	var AuditLogs []models.AuditLog
	for _, row := range rows {
		AuditLog, ok := row.(models.AuditLog)
		if !ok {
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		AuditLogs = append(AuditLogs, AuditLog)
	}

	return utils.Response(200, AuditLogs), nil
}

// GetAuditLogById - Get Business Unit
func (s *AuditLogAPIService) GetAuditLogById(ctx context.Context, unitId int32) (utils.ImplResponse, error) {
	privilege := "read"

	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.AuditLog
	row, err := dbConnection.GetByID("AuditLog", "logid", unitId, &dest)
	if err != nil {
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	unit, ok := row.(models.AuditLog)
	if !ok {
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}

	return utils.Response(200, unit), nil
}
