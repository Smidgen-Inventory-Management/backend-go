package smidgen

import (
	"fmt"
	"reflect"
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
