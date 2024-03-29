package models

import "reflect"

// Models ...
type Models struct {
	columns   []string
	fieldsptr []interface{}
}

// Columns is a method to get columns name
func (m *Models) Columns(models interface{}) []string {
	m.columns = columns(models)
	return m.columns
}

// Fields is a method to get fields pointer
func (m *Models) Fields(models interface{}) []interface{} {
	m.fieldsptr = fields(models)
	return m.fieldsptr
}

// columns is a helper function to get all columns from a tag struct
func columns(models interface{}) []string {

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

// fields is a helper to populate field pointer from struct
func fields(models interface{}) []interface{} {
	// get value from interface
	mdl := reflect.ValueOf(models)
	// check if models kind is pointer
	if mdl.Kind() == reflect.Ptr {
		mdl = mdl.Elem()
	}

	var fields []interface{}
	// iterate all field on struct
	for i := 0; i < mdl.NumField(); i++ {
		field := mdl.Field(i)

		if _, ok := mdl.Type().Field(i).Tag.Lookup("column"); !ok {
			continue
		}

		// push column to slice
		fields = append(fields, field.Addr().Interface())
	}
	return fields

}
