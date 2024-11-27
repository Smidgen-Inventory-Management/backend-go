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
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrTypeAssertionError = errors.New("unable to assert type")
)

type ParsingError struct {
	Err error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	return e.Err.Error()
}

type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse) {
	if _, ok := err.(*ParsingError); ok {
		EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
	} else if _, ok := err.(*RequiredError); ok {
		EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
	} else {
		EncodeJSONResponse(err.Error(), &result.Code, w)
	}
}
