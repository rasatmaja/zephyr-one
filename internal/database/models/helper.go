package models

import (
	"reflect"
)

// Columns is a helper function to get all columns from a tag struct
func Columns(models interface{}) []string {

	// get value from interface
	mdl := reflect.ValueOf(models)
	// check if models kind is pointer
	if mdl.Kind() == reflect.Ptr {
		mdl = mdl.Elem()
	}

	var columns []string
	// iterate all field on struct
	for i := 0; i < mdl.NumField(); i++ {
		types := mdl.Type().Field(i)
		// serach for "column" tag
		column, ok := types.Tag.Lookup("column")
		if !ok {
			continue
		}

		// push column to slice
		columns = append(columns, column)
	}
	return columns
}

// Fields is a helper to populate field pointer from struct
func Fields(models interface{}) []interface{} {
	// get value from interface
	mdl := reflect.ValueOf(models)
	// check if models kind is pointer
	if mdl.Kind() == reflect.Ptr {
		mdl = mdl.Elem()
	}

	var fields []interface{}
	// iterate all field on struct
	for i := 0; i < mdl.NumField(); i++ {
		value := mdl.Field(i)
		//field := mdl.Type().Field(i)

		//fmt.Println(value.Addr().Pointer())
		//fmt.Println(value.Kind())
		//fmt.Println(field)

		// push column to slice
		fields = append(fields, value.Addr().Interface())
	}
	return fields

}
