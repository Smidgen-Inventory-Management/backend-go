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
	"encoding/json"
	"net/http"
	models "smidgen-backend/src/models"
	utils "smidgen-backend/src/utils"
	"strings"

	"github.com/gorilla/mux"
)

// BusinessUnitAPIController binds http requests to an api service and writes the service results to the http response
type BusinessUnitAPIController struct {
	service      BusinessUnitAPIServicer
	errorHandler utils.ErrorHandler
}

// BusinessUnitAPIOption for how the controller is set up.
type BusinessUnitAPIOption func(*BusinessUnitAPIController)

// WithBusinessUnitAPIErrorHandler inject ErrorHandler into controller
func WithBusinessUnitAPIErrorHandler(h utils.ErrorHandler) BusinessUnitAPIOption {
	return func(c *BusinessUnitAPIController) {
		c.errorHandler = h
	}
}

// NewBusinessUnitAPIController creates a default api controller
func NewBusinessUnitAPIController(s BusinessUnitAPIServicer, opts ...BusinessUnitAPIOption) utils.Router {
	controller := &BusinessUnitAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BusinessUnitAPIController
func (c *BusinessUnitAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddBusinessUnit": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "business_unit",
			HandlerFunc: c.AddBusinessUnit,
		},
		"DeletBusinessUnit": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.DeleteBusinessUnit,
		},
		"GetBusinessUnit": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "business_unit/",
			HandlerFunc: c.GetBusinessUnit,
		},
		"GetBusinessUnitById": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.GetBusinessUnitById,
		},
		"UpdateBusinessUnit": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.UpdateBusinessUnit,
		},
	}
}

// AddBusinessUnit - Create Business Unit
func (c *BusinessUnitAPIController) AddBusinessUnit(w http.ResponseWriter, r *http.Request) {

	businessUnitParam := models.BusinessUnit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&businessUnitParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBusinessUnitRequired(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBusinessUnitConstraints(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddBusinessUnit(r.Context(), businessUnitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteBusinessUnit - Delete Business Unit
func (c *BusinessUnitAPIController) DeleteBusinessUnit(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteBusinessUnit(r.Context(), unitIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetBusinessUnit - Get Business Units
func (c *BusinessUnitAPIController) GetBusinessUnit(w http.ResponseWriter, r *http.Request) {

	result, err := c.service.GetBusinessUnits(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetBusinessUnitById - Get Business Unit
func (c *BusinessUnitAPIController) GetBusinessUnitById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetBusinessUnitById(r.Context(), unitIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateBusinessUnit - Update Business Unit
func (c *BusinessUnitAPIController) UpdateBusinessUnit(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	businessUnitParam := models.BusinessUnit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&businessUnitParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBusinessUnitRequired(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBusinessUnitConstraints(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateBusinessUnit(r.Context(), unitIdParam, businessUnitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
