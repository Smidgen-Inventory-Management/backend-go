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
	"database/sql"
	utils "smidgen-backend/go/utils"
)

type User struct {
	UserId         int32          `json:"user_id"`
	BusinessUnitId int32          `json:"business_unit_id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	SaltedPassword sql.NullString `json:"salted_password"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"business_unit_id": obj.BusinessUnitId,
		"username":         obj.Username,
		"email":            obj.Email,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj User) error {
	return nil
}
