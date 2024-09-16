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
	"net/http"
	utils "smidgen-backend/src/utils"
	"strings"
	"github.com/gorilla/mux"
)

// AuditLogAPIController binds http requests to an api service and writes the service results to the http response
type AuditLogAPIController struct {
	service      AuditLogAPIServicer
	errorHandler utils.ErrorHandler
}

// AuditLogAPIOption for how the controller is set up.
type AuditLogAPIOption func(*AuditLogAPIController)

// WithAuditLogAPIErrorHandler inject ErrorHandler into controller
func WithAuditLogAPIErrorHandler(h utils.ErrorHandler) AuditLogAPIOption {
	return func(c *AuditLogAPIController) {
		c.errorHandler = h
	}
}

// NewAuditLogAPIController creates a default api controller
func NewAuditLogAPIController(s AuditLogAPIServicer, opts ...AuditLogAPIOption) utils.Router {
	controller := &AuditLogAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AuditLogAPIController
func (c *AuditLogAPIController) Routes() utils.Routes {
	return utils.Routes{
		"GetAuditLog": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "audit_log/",
			HandlerFunc: c.GetAuditLog,
		},
		"GetAuditLogById": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "audit_log/{log_id}",
			HandlerFunc: c.GetAuditLogById,
		},
	}
}

// GetAuditLog - Get Business Units
func (c *AuditLogAPIController) GetAuditLog(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAuditLogs(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAuditLogById - Get Business Unit
func (c *AuditLogAPIController) GetAuditLogById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["log_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetAuditLogById(r.Context(), unitIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
