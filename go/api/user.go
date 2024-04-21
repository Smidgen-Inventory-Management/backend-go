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
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
	"strings"

	"github.com/gorilla/mux"
)

// UserAPIController binds http requests to an api service and writes the service results to the http response
type UserAPIController struct {
	service      UserAPIServicer
	errorHandler utils.ErrorHandler
}

// UserAPIOption for how the controller is set up.
type UserAPIOption func(*UserAPIController)

// WithUserAPIErrorHandler inject ErrorHandler into controller
func WithUserAPIErrorHandler(h utils.ErrorHandler) UserAPIOption {
	return func(c *UserAPIController) {
		c.errorHandler = h
	}
}

// NewUserAPIController creates a default api controller
func NewUserAPIController(s UserAPIServicer, opts ...UserAPIOption) utils.Router {
	controller := &UserAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// utils.Routes returns all the api utils.routes for the UserAPIController
func (c *UserAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddUserUserPost": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "user",
			HandlerFunc: c.AddUserUserPost,
		},
		"DeleteUserUserUserIdDelete": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "user/{user_id}",
			HandlerFunc: c.DeleteUserUserUserIdDelete,
		},
		"GetUserUserGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "user/",
			HandlerFunc: c.GetUserUserGet,
		},
		"GetUserAssignmentIdGetAssignments": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "user/{user_id}/assignments",
			HandlerFunc: c.GetUserAssignmentIdGetAssignments,
		},
		"GetUserUserUserIdGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "user/{user_id}",
			HandlerFunc: c.GetUserUserUserIdGet,
		},
		"UpdateUserUserUserIdPut": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "user/{user_id}",
			HandlerFunc: c.UpdateUserUserUserIdPut,
		},
	}
}

// AddUserUserPost - Create user
func (c *UserAPIController) AddUserUserPost(w http.ResponseWriter, r *http.Request) {
	userParam := models.User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertUserConstraints(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddUserUserPost(r.Context(), userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteUserUserUserIdDelete - Delete user
func (c *UserAPIController) DeleteUserUserUserIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdParam, err := utils.ParseNumericParameter[int32](
		params["user_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteUserUserUserIdDelete(r.Context(), userIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserUserGet - Get Users
func (c *UserAPIController) GetUserUserGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetUserUserGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserUserUserIdGet - Get user
func (c *UserAPIController) GetUserUserUserIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdParam, err := utils.ParseNumericParameter[int32](
		params["user_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetUserUserUserIdGet(r.Context(), userIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserAssignmentIdGetAssignments - Get user assignments
func (c *UserAPIController) GetUserAssignmentIdGetAssignments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdParam, err := utils.ParseNumericParameter[int32](
		params["user_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetUserUserUserIdGet(r.Context(), userIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateUserUserUserIdPut - Update user
func (c *UserAPIController) UpdateUserUserUserIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdParam, err := utils.ParseNumericParameter[int32](
		params["user_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	userParam := models.User{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertUserRequired(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertUserConstraints(userParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserUserUserIdPut(r.Context(), userIdParam, userParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
