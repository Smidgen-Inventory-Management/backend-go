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
	"fmt"
	"reflect"
	"strings"

	"github.com/lib/pq"
)

// GetRows operates by using an instance of DatabaseConnection.
// GetRows must be provided a non-prefixed tableName, and uses a struct for data return.
// GetRows will return a new interface{} object.
func (dao *DatabaseConnection) GetRows(tableName string, dest interface{}) ([]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM smidgen.%s", tableName)

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("\nfailed to query rows from table smidgen.%s: %v", tableName, err)
	}
	defer rows.Close()

	elementType := reflect.TypeOf(dest).Elem()

	destValues := make([]interface{}, 0)
	for i := 0; i < elementType.NumField(); i++ {
		destValues = append(destValues, reflect.New(elementType.Field(i).Type).Interface())
	}

	var results []interface{}

	for rows.Next() {
		if err := rows.Scan(destValues...); err != nil {
			return nil, fmt.Errorf("\nfailed to scan rows from table smidgen.%s: %v", tableName, err)
		}

		result := reflect.New(elementType).Elem()
		for i := 0; i < elementType.NumField(); i++ {
			result.Field(i).Set(reflect.Indirect(reflect.ValueOf(destValues[i])))
		}

		results = append(results, result.Interface())
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("\nerror while iterating over rows from table smidgen.%s: %v", tableName, err)
	}

	return results, nil
}

func (dao *DatabaseConnection) GetByID(tableName string, idName string, id int32, dest interface{}) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM smidgen.%s WHERE %s=%d", tableName, idName, id)

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("\nfailed to query rows from table smidgen.%s: %v", tableName, err)
	}
	defer rows.Close()

	objectType := reflect.TypeOf(dest).Elem()

	destValues := make([]interface{}, 0)
	for i := 0; i < objectType.NumField(); i++ {
		destValues = append(destValues, reflect.New(objectType.Field(i).Type).Interface())
	}

	if !rows.Next() {
		return nil, fmt.Errorf("no rows returned by the query")
	}

	if err := rows.Scan(destValues...); err != nil {
		return nil, fmt.Errorf("\nfailed to scan rows from table smidgen.%s: %v", tableName, err)
	}

	result := reflect.New(objectType).Elem()
	for i := 0; i < objectType.NumField(); i++ {
		result.Field(i).Set(reflect.Indirect(reflect.ValueOf(destValues[i])))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("\nerror while iterating over rows from table smidgen.%s: %v", tableName, err)
	}

	return result.Interface(), nil
}

func (dao *DatabaseConnection) InsertRow(tableName string, values interface{}) error {
	v := reflect.ValueOf(values)
	primaryID := v.Type().Field(0).Name

	tx, err := dao.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var lastInsertedID int
	var columns []string
	var placeholders []string

	if err := tx.QueryRow("SELECT COALESCE(MAX(" + primaryID + "), 0) FROM smidgen." + tableName).Scan(&lastInsertedID); err != nil && err != sql.ErrNoRows {
		return err
	}

	newID := lastInsertedID + 1

	columns = append(columns, primaryID)
	placeholders = append(placeholders, "$1")
	for i := 1; i < v.NumField(); i++ {
		columns = append(columns, v.Type().Field(i).Name)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf("INSERT INTO smidgen.%s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fieldValues := make([]interface{}, v.NumField())
	fieldValues[0] = newID
	for i := 1; i < v.NumField(); i++ {
		fieldValues[i] = v.Field(i).Interface()
	}

	_, err = stmt.Exec(fieldValues...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23503" {
				return fmt.Errorf("23503")
			}
		}
	}

	return tx.Commit()
}

func (dao *DatabaseConnection) DeleteRow(tableName string, idLabel string, id int32, args ...interface{}) error {
	tx, err := dao.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := fmt.Sprintf("DELETE FROM smidgen.%s WHERE %s=%v", tableName, idLabel, id)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d does not exist in table %s", id, tableName)
	}
	return tx.Commit()
}

func (dao *DatabaseConnection) UpdateRow(tableName string, idLabel string, id int32, values interface{}) error {
	v := reflect.ValueOf(values)
	objectType := v.Type()

	var setValues []string
	for i := 1; i < v.NumField(); i++ { // Start from index 1 to exclude the first column
		fieldName := objectType.Field(i).Name
		setValues = append(setValues, fmt.Sprintf("%s=$%d", fieldName, i))
	}

	setClause := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE smidgen.%s SET %s WHERE %s=%v", tableName, setClause, idLabel, id)

	tx, err := dao.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var fieldValues []interface{}
	for i := 1; i < v.NumField(); i++ { // Start from index 1 to exclude the first column
		fieldValues = append(fieldValues, v.Field(i).Interface())
	}

	result, err := stmt.Exec(fieldValues...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d does not exist in table %s", id, tableName)
	}

	return tx.Commit()
}
