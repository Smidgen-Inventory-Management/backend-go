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

var log = utils.Log()

// UserAPIService is a service that implements the logic for the UserAPIServicer
// This service should implement the business logic for every endpoint for the UserAPI API.
// Include any external packages or services that will be required by this service.
type UserAPIService struct {
}

// NewUserAPIService creates a default api service
func NewUserAPIService() api.UserAPIServicer {
	return &UserAPIService{}
}

// AddUser - Create user
func (s *UserAPIService) AddUser(ctx context.Context, user models.User) (utils.ImplResponse, error) {
	privilege := "write"
	
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("smidgenUsers", user)
	if err != nil {
		log.Error(err)
		return utils.Response(500, nil), errors.New("an error has occured while adding new data")
	}
	return utils.Response(202, nil), nil
}

// DeleteUser - Delete user
func (s *UserAPIService) DeleteUser(ctx context.Context, userId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("smidgenUsers", "userid", userId)
	if err != nil {
		log.Errorf("Error: %v", err)
		return utils.Response(404, nil), err
	}

	return utils.Response(200, nil), nil
}

// GetUsers - Get Users
func (s *UserAPIService) GetUsers(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.User
	rows, err := dbConnection.GetRows("smidgenUsers", &dest)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	if len(rows) == 0 {
		return utils.Response(404, nil), fmt.Errorf("no users were found in the database")
	}

	var users []models.User
	for _, row := range rows {
		user, ok := row.(models.User)
		if !ok {
			log.Warn("Warn: Unexpected type in row")
			continue
		}
		users = append(users, user)
	}

	return utils.Response(200, users), nil
}

// GetUserById - Get user
func (s *UserAPIService) GetUserById(ctx context.Context, userId int32) (utils.ImplResponse, error) {
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.User
	row, err := dbConnection.GetByID("smidgenUsers", "userId", userId, &dest)
	if err != nil {
		log.Errorf("Data Not Found: %v", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	user, ok := row.(models.User)
	if !ok {
		log.Warn("Warn: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	return utils.Response(200, user), nil
}

// UpdateUser - Update user
func (s *UserAPIService) UpdateUser(ctx context.Context, userId int32, user models.User) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, privilege)
	if err != nil {
		log.Errorf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("smidgenUsers", "userid", userId, user)
	if err != nil {
		log.Error(err)
		return utils.Response(400, nil), err
	}
	return utils.Response(202, nil), nil
}
