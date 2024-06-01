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
	"net/http"
	models "smidgen-backend/src/models"
	utils "smidgen-backend/src/utils"
)

// BusinessUnitAPIRouter defines the required methods for binding the api requests to a responses for the BusinessUnitAPI
// The BusinessUnitAPIRouter implementation should parse necessary information from the http request,
// pass the data to a BusinessUnitAPIServicer to perform the required actions, then write the service results to the http response.
type BusinessUnitAPIRouter interface {
	AddBusinessUnit(http.ResponseWriter, *http.Request)
	DeleteBusinessUnitById(http.ResponseWriter, *http.Request)
	GetBusinessUnits(http.ResponseWriter, *http.Request)
	GetBusinessUnitById(http.ResponseWriter, *http.Request)
	UpdateBusinessUnit(http.ResponseWriter, *http.Request)
	GetUserAssignmentsById(http.ResponseWriter, *http.Request)
}

// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface {
	HealthCheck(http.ResponseWriter, *http.Request)
	RootGet(http.ResponseWriter, *http.Request)
}

// EquipmentAPIRouter defines the required methods for binding the api requests to a responses for the EquipmentAPI
// The EquipmentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a EquipmentAPIServicer to perform the required actions, then write the service results to the http response.
type EquipmentAPIRouter interface {
	AddEquipment(http.ResponseWriter, *http.Request)
	DeleteEquipment(http.ResponseWriter, *http.Request)
	GetEquipments(http.ResponseWriter, *http.Request)
	GetEquipmentById(http.ResponseWriter, *http.Request)
	UpdateEquipment(http.ResponseWriter, *http.Request)
}

// EquipmentAssignmentAPIRouter defines the required methods for binding the api requests to a responses for the EquipmentAssignmentAPI
// The EquipmentAssignmentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a EquipmentAssignmentAPIServicer to perform the required actions, then write the service results to the http response.
type EquipmentAssignmentAPIRouter interface {
	AddEquipmentAssignment(http.ResponseWriter, *http.Request)
	DeleteEquipmentAssignment(http.ResponseWriter, *http.Request)
	GetEquipmentAssignments(http.ResponseWriter, *http.Request)
	GetEquipmentAssignmentById(http.ResponseWriter, *http.Request)
	UpdateEquipmentAssignment(http.ResponseWriter, *http.Request)
}

// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface {
	AddUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetUserById(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
}

// BusinessUnitAPIServicer defines the api actions for the BusinessUnitAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BusinessUnitAPIServicer interface {
	AddBusinessUnit(context.Context, models.BusinessUnit) (utils.ImplResponse, error)
	DeleteBusinessUnit(context.Context, int32) (utils.ImplResponse, error)
	GetBusinessUnits(context.Context) (utils.ImplResponse, error)
	GetBusinessUnitById(context.Context, int32) (utils.ImplResponse, error)
	UpdateBusinessUnit(context.Context, int32, models.BusinessUnit) (utils.ImplResponse, error)
}

// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface {
	HealthCheck(context.Context) (utils.ImplResponse, error)
	RootGet(context.Context) (utils.ImplResponse, error)
}

// EquipmentAPIServicer defines the api actions for the EquipmentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EquipmentAPIServicer interface {
	AddEquipment(context.Context, models.Equipment) (utils.ImplResponse, error)
	DeleteEquipment(context.Context, int32) (utils.ImplResponse, error)
	GetEquipments(context.Context) (utils.ImplResponse, error)
	GetEquipmentById(context.Context, int32) (utils.ImplResponse, error)
	UpdateEquipment(context.Context, int32, models.Equipment) (utils.ImplResponse, error)
}

// EquipmentAssignmentAPIServicer defines the api actions for the EquipmentAssignmentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EquipmentAssignmentAPIServicer interface {
	AddEquipmentAssignment(context.Context, models.EquipmentAssignment) (utils.ImplResponse, error)
	DeleteEquipmentAssignment(context.Context, int32) (utils.ImplResponse, error)
	GetEquipmentAssignments(context.Context) (utils.ImplResponse, error)
	GetEquipmentAssignmentById(context.Context, int32) (utils.ImplResponse, error)
	UpdateEquipmentAssignment(context.Context, int32, models.EquipmentAssignment) (utils.ImplResponse, error)
}

// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface {
	AddUser(context.Context, models.User) (utils.ImplResponse, error)
	DeleteUser(context.Context, int32) (utils.ImplResponse, error)
	GetUsers(context.Context) (utils.ImplResponse, error)
	GetUserById(context.Context, int32) (utils.ImplResponse, error)
	UpdateUser(context.Context, int32, models.User) (utils.ImplResponse, error)
}
