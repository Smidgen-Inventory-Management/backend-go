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
	utils "smidgen-backend/go/utils"
)

type ValidationError struct {
	Loc []LocationInner `json:"loc"`
	Msg string `json:"msg"`
	Type string `json:"type"`
}

// AssertValidationErrorRequired checks if the required fields are not zero-ed
func AssertValidationErrorRequired(obj ValidationError) error {
	elements := map[string]interface{}{
		"loc":  obj.Loc,
		"msg":  obj.Msg,
		"type": obj.Type,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	for _, el := range obj.Loc {
		if err := AssertLocationInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertValidationErrorConstraints checks if the values respects the defined constraints
func AssertValidationErrorConstraints(obj ValidationError) error {
	return nil
}
