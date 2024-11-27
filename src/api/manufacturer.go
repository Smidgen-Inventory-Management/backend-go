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

type ManufacturerAPIController struct {
	service      ManufacturerAPIServicer
	errorHandler utils.ErrorHandler
}

type ManufacturerAPIOption func(*ManufacturerAPIController)

func WithManufacturerAPIErrorHandler(h utils.ErrorHandler) ManufacturerAPIOption {
	return func(c *ManufacturerAPIController) {
		c.errorHandler = h
	}
}

func NewManufacturerAPIController(s ManufacturerAPIServicer, opts ...ManufacturerAPIOption) utils.Router {
	controller := &ManufacturerAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

func (c *ManufacturerAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddManufacturer": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "manufacturer",
			HandlerFunc: c.AddManufacturer,
		},
		"DeleteManufacturer": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "manufacturer/{manufacturer_id}",
			HandlerFunc: c.DeleteManufacturer,
		},
		"GetManufacturer": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "manufacturer/",
			HandlerFunc: c.GetManufacturer,
		},
		"GetManufacturerById": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "manufacturer/{manufacturer_id}",
			HandlerFunc: c.GetManufacturerById,
		},
		"UpdateManufacturer": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "manufacturer/{manufacturer_id}",
			HandlerFunc: c.UpdateManufacturer,
		},
	}
}

func (c *ManufacturerAPIController) AddManufacturer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ManufacturerParam := models.Manufacturer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&ManufacturerParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertManufacturerRequired(ManufacturerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertManufacturerConstraints(ManufacturerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddManufacturer(r.Context(), ManufacturerParam)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *ManufacturerAPIController) DeleteManufacturer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ManufacturerIdParam, err := utils.ParseNumericParameter[int32](
		params["manufacturer_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteManufacturer(r.Context(), ManufacturerIdParam)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *ManufacturerAPIController) GetManufacturer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	result, err := c.service.GetManufacturers(r.Context())
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *ManufacturerAPIController) GetManufacturerById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ManufacturerIdParam, err := utils.ParseNumericParameter[int32](
		params["manufacturer_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetManufacturerById(r.Context(), ManufacturerIdParam)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

func (c *ManufacturerAPIController) UpdateManufacturer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ManufacturerIdParam, err := utils.ParseNumericParameter[int32](
		params["manufacturer_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	ManufacturerParam := models.Manufacturer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&ManufacturerParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertManufacturerRequired(ManufacturerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertManufacturerConstraints(ManufacturerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateManufacturer(r.Context(), ManufacturerIdParam, ManufacturerParam)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
