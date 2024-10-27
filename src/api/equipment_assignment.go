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

// EquipmentAssignmentAPIController binds http requests to an api service and writes the service results to the http response
type EquipmentAssignmentAPIController struct {
	service      EquipmentAssignmentAPIServicer
	errorHandler utils.ErrorHandler
}

// EquipmentAssignmentAPIOption for how the controller is set up.
type EquipmentAssignmentAPIOption func(*EquipmentAssignmentAPIController)

// WithEquipmentAssignmentAPIErrorHandler inject ErrorHandler into controller
func WithEquipmentAssignmentAPIErrorHandler(h utils.ErrorHandler) EquipmentAssignmentAPIOption {
	return func(c *EquipmentAssignmentAPIController) {
		c.errorHandler = h
	}
}

// NewEquipmentAssignmentAPIController creates a default api controller
func NewEquipmentAssignmentAPIController(s EquipmentAssignmentAPIServicer, opts ...EquipmentAssignmentAPIOption) utils.Router {
	controller := &EquipmentAssignmentAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EquipmentAssignmentAPIController
func (c *EquipmentAssignmentAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddEquipmentAssignment": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "equipment_assignment/",
			HandlerFunc: c.AddEquipmentAssignment,
		},
		"DeleteEquipmentAssignment": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "equipment_assignment/{assignment_id}",
			HandlerFunc: c.DeleteEquipmentAssignment,
		},
		"GetEquipmentAssignment": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment_assignment/",
			HandlerFunc: c.GetEquipmentAssignment,
		},
		"GetEquipmentAssignmentById": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment_assignment/{assignment_id}",
			HandlerFunc: c.GetEquipmentAssignmentById,
		},
		"UpdateEquipmentAssignment": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "equipment_assignment/{assignment_id}",
			HandlerFunc: c.UpdateEquipmentAssignment,
		},
	}
}

// AddEquipmentAssignment - Create assignment
func (c *EquipmentAssignmentAPIController) AddEquipmentAssignment(w http.ResponseWriter, r *http.Request) {

	equipmentAssignmentParam := models.EquipmentAssignment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentAssignmentRequired(equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentAssignmentConstraints(equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddEquipmentAssignment(r.Context(), equipmentAssignmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEquipmentAssignment - Delete assignment
func (c *EquipmentAssignmentAPIController) DeleteEquipmentAssignment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	assignmentIdParam, err := utils.ParseNumericParameter[int32](
		params["assignment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteEquipmentAssignment(r.Context(), assignmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipmentAssignment - Get assignments
func (c *EquipmentAssignmentAPIController) GetEquipmentAssignment(w http.ResponseWriter, r *http.Request) {

	result, err := c.service.GetEquipmentAssignments(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipmentAssignmentById - Get assignment
func (c *EquipmentAssignmentAPIController) GetEquipmentAssignmentById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	assignmentIdParam, err := utils.ParseNumericParameter[int32](
		params["assignment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetEquipmentAssignmentById(r.Context(), assignmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEquipmentAssignment - Update assignment
func (c *EquipmentAssignmentAPIController) UpdateEquipmentAssignment(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	assignmentIdParam, err := utils.ParseNumericParameter[int32](
		params["assignment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	equipmentAssignmentParam := models.EquipmentAssignment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentAssignmentRequired(equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentAssignmentConstraints(equipmentAssignmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEquipmentAssignment(r.Context(), assignmentIdParam, equipmentAssignmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
