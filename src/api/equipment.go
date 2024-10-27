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

// EquipmentAPIController binds http requests to an api service and writes the service results to the http response
type EquipmentAPIController struct {
	service      EquipmentAPIServicer
	errorHandler utils.ErrorHandler
}

// EquipmentAPIOption for how the controller is set up.
type EquipmentAPIOption func(*EquipmentAPIController)

// WithEquipmentAPIErrorHandler inject ErrorHandler into controller
func WithEquipmentAPIErrorHandler(h utils.ErrorHandler) EquipmentAPIOption {
	return func(c *EquipmentAPIController) {
		c.errorHandler = h
	}
}

// NewEquipmentAPIController creates a default api controller
func NewEquipmentAPIController(s EquipmentAPIServicer, opts ...EquipmentAPIOption) utils.Router {
	controller := &EquipmentAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EquipmentAPIController
func (c *EquipmentAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddEquipment": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "equipment",
			HandlerFunc: c.AddEquipment,
		},
		"DeleteEquipment": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.DeleteEquipment,
		},
		"GetEquipment": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment/",
			HandlerFunc: c.GetEquipment,
		},
		"GetEquipmentById": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.GetEquipmentById,
		},
		"UpdateEquipment": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.UpdateEquipment,
		},
	}
}

// AddEquipment - Create equipment
func (c *EquipmentAPIController) AddEquipment(w http.ResponseWriter, r *http.Request) {
	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	equipmentParam := models.Equipment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentRequired(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentConstraints(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddEquipment(r.Context(), equipmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEquipment - Delete equipment
func (c *EquipmentAPIController) DeleteEquipment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteEquipment(r.Context(), equipmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipment - Get equipments
func (c *EquipmentAPIController) GetEquipment(w http.ResponseWriter, r *http.Request) {
	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	result, err := c.service.GetEquipments(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipmentById - Get equipment
func (c *EquipmentAPIController) GetEquipmentById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetEquipmentById(r.Context(), equipmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEquipment - Update equipment
func (c *EquipmentAPIController) UpdateEquipment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	equipmentParam := models.Equipment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentRequired(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentConstraints(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEquipment(r.Context(), equipmentIdParam, equipmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
