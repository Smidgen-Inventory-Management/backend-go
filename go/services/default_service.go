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
	"net/http"
	api "smidgen-backend/go/api"
	utils "smidgen-backend/go/utils"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() api.DefaultAPIServicer {
	return &DefaultAPIService{}
}

// CheckHealthcheckGet - Check
func (s *DefaultAPIService) CheckHealthcheckGet(ctx context.Context) (utils.ImplResponse, error) {
	// TODO - update CheckHealthcheckGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response utils.Response(200, interface{}{}) or use other options such as http.Ok ...
	// return utils.Response(200, interface{}{}), nil

	// TODO: Uncomment the next line to return response utils.Response(404, {}) or use other options such as http.Ok ...
	// return utils.Response(404, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(401, {}) or use other options such as http.Ok ...
	// return utils.Response(401, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(403, {}) or use other options such as http.Ok ...
	// return utils.Response(403, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(500, {}) or use other options such as http.Ok ...
	// return utils.Response(500, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(501, {}) or use other options such as http.Ok ...
	// return utils.Response(501, nil),nil

	return utils.Response(http.StatusNotImplemented, nil), errors.New("CheckHealthcheckGet method not implemented")
}

// RootGet - Root
func (s *DefaultAPIService) RootGet(ctx context.Context) (utils.ImplResponse, error) {
	// TODO - update RootGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response utils.Response(200, interface{}{}) or use other options such as http.Ok ...
	// return utils.Response(200, interface{}{}), nil

	// TODO: Uncomment the next line to return response utils.Response(404, {}) or use other options such as http.Ok ...
	// return utils.Response(404, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(401, {}) or use other options such as http.Ok ...
	// return utils.Response(401, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(403, {}) or use other options such as http.Ok ...
	// return utils.Response(403, nil),nil

	// TODO: Uncomment the next line to return response utils.Response(500, {}) or use other options such as http.Ok ...
	// return utils.Response(500, nil),nil

	return utils.Response(http.StatusNotImplemented, nil), errors.New("RootGet method not implemented")
}
