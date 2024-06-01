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
	utils "smidgen-backend/src/utils"
	"time"
)

type EquipmentAssignment struct {
	AssignmentId int32 `json:"assignment_id"`
	UserId int32 `json:"user_id"`
	EquipmentId int32 `json:"equipment_id"`
	DateOfAssignment time.Time `json:"date_of_assignment"`
}

func AssertEquipmentAssignmentRequired(obj EquipmentAssignment) error {
	elements := map[string]interface{}{
		"assignment_id":      obj.AssignmentId,
		"user_id":            obj.UserId,
		"equipment_id":       obj.EquipmentId,
		"date_of_assignment": obj.DateOfAssignment,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

func AssertEquipmentAssignmentConstraints(obj EquipmentAssignment) error {
	return nil
}
