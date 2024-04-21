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
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
)

// BusinessUnitAPIRouter defines the required methods for binding the api requests to a responses for the BusinessUnitAPI
// The BusinessUnitAPIRouter implementation should parse necessary information from the http request,
// pass the data to a BusinessUnitAPIServicer to perform the required actions, then write the service results to the http response.
type BusinessUnitAPIRouter interface {
	AddUnitBusinessUnitPost(http.ResponseWriter, *http.Request)
	DeleteUnitBusinessUnitUnitIdDelete(http.ResponseWriter, *http.Request)
	GetUnitBusinessUnitGet(http.ResponseWriter, *http.Request)
	GetUnitsBusinessUnitUnitIdGet(http.ResponseWriter, *http.Request)
	UpdateUnitBusinessUnitUnitIdPut(http.ResponseWriter, *http.Request)
	GetUserAssignmentIdGetAssignments(http.ResponseWriter, *http.Request)
}

// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface {
	CheckHealthcheckGet(http.ResponseWriter, *http.Request)
	RootGet(http.ResponseWriter, *http.Request)
}

// EquipmentAPIRouter defines the required methods for binding the api requests to a responses for the EquipmentAPI
// The EquipmentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a EquipmentAPIServicer to perform the required actions, then write the service results to the http response.
type EquipmentAPIRouter interface {
	AddEquipmentEquipmentPost(http.ResponseWriter, *http.Request)
	DeleteEquipmentEquipmentEquipmentIdDelete(http.ResponseWriter, *http.Request)
	GetEquipmentEquipmentGet(http.ResponseWriter, *http.Request)
	GetEquipmentsEquipmentEquipmentIdGet(http.ResponseWriter, *http.Request)
	UpdateEquipmentEquipmentEquipmentIdPut(http.ResponseWriter, *http.Request)
}

// EquipmentAssignmentAPIRouter defines the required methods for binding the api requests to a responses for the EquipmentAssignmentAPI
// The EquipmentAssignmentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a EquipmentAssignmentAPIServicer to perform the required actions, then write the service results to the http response.
type EquipmentAssignmentAPIRouter interface {
	AddAssignmentEquipmentAssignmentPost(http.ResponseWriter, *http.Request)
	DeleteAssignmentEquipmentAssignmentAssignmentIdDelete(http.ResponseWriter, *http.Request)
	GetAssignmentEquipmentAssignmentGet(http.ResponseWriter, *http.Request)
	GetAssignmentsEquipmentAssignmentAssignmentIdGet(http.ResponseWriter, *http.Request)
	UpdateAssignmentEquipmentAssignmentAssignmentIdPut(http.ResponseWriter, *http.Request)
}

// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface {
	AddUserUserPost(http.ResponseWriter, *http.Request)
	DeleteUserUserUserIdDelete(http.ResponseWriter, *http.Request)
	GetUserUserGet(http.ResponseWriter, *http.Request)
	GetUserUserUserIdGet(http.ResponseWriter, *http.Request)
	UpdateUserUserUserIdPut(http.ResponseWriter, *http.Request)
}

// BusinessUnitAPIServicer defines the api actions for the BusinessUnitAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BusinessUnitAPIServicer interface {
	AddUnitBusinessUnitPost(context.Context, models.BusinessUnit) (utils.ImplResponse, error)
	DeleteUnitBusinessUnitUnitIdDelete(context.Context, int32) (utils.ImplResponse, error)
	GetUnitBusinessUnitGet(context.Context) (utils.ImplResponse, error)
	GetUnitsBusinessUnitUnitIdGet(context.Context, int32) (utils.ImplResponse, error)
	UpdateUnitBusinessUnitUnitIdPut(context.Context, int32, models.BusinessUnit) (utils.ImplResponse, error)
}

// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface {
	CheckHealthcheckGet(context.Context) (utils.ImplResponse, error)
	RootGet(context.Context) (utils.ImplResponse, error)
}

// EquipmentAPIServicer defines the api actions for the EquipmentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EquipmentAPIServicer interface {
	AddEquipmentEquipmentPost(context.Context, models.Equipment) (utils.ImplResponse, error)
	DeleteEquipmentEquipmentEquipmentIdDelete(context.Context, int32) (utils.ImplResponse, error)
	GetEquipmentEquipmentGet(context.Context) (utils.ImplResponse, error)
	GetEquipmentsEquipmentEquipmentIdGet(context.Context, int32) (utils.ImplResponse, error)
	UpdateEquipmentEquipmentEquipmentIdPut(context.Context, int32, models.Equipment) (utils.ImplResponse, error)
}

// EquipmentAssignmentAPIServicer defines the api actions for the EquipmentAssignmentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EquipmentAssignmentAPIServicer interface {
	AddAssignmentEquipmentAssignmentPost(context.Context, models.EquipmentAssignment) (utils.ImplResponse, error)
	DeleteAssignmentEquipmentAssignmentAssignmentIdDelete(context.Context, int32) (utils.ImplResponse, error)
	GetAssignmentEquipmentAssignmentGet(context.Context) (utils.ImplResponse, error)
	GetAssignmentsEquipmentAssignmentAssignmentIdGet(context.Context, int32) (utils.ImplResponse, error)
	UpdateAssignmentEquipmentAssignmentAssignmentIdPut(context.Context, int32, models.EquipmentAssignment) (utils.ImplResponse, error)
}

// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserAPIServicer interface {
	AddUserUserPost(context.Context, models.User) (utils.ImplResponse, error)
	DeleteUserUserUserIdDelete(context.Context, int32) (utils.ImplResponse, error)
	GetUserUserGet(context.Context) (utils.ImplResponse, error)
	GetUserUserUserIdGet(context.Context, int32) (utils.ImplResponse, error)
	UpdateUserUserUserIdPut(context.Context, int32, models.User) (utils.ImplResponse, error)
}
