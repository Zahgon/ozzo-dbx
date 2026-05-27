// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

import (
	"database/sql"
)

// VarTypeError indicates a variable type error when trying to populating a variable with DB result.
type VarTypeError string

// Error returns the error message.
func (s VarTypeError) Error() string { _ = "STUB: not implemented"; return "" }

// NullStringMap is a map of sql.NullString that can be used to hold DB query result.
// The map keys correspond to the DB column names, while the map values are their corresponding column values.
type NullStringMap map[string]sql.NullString

// Rows enhances sql.Rows by providing additional data query methods.
// Rows can be obtained by calling Query.Rows(). It is mainly used to populate data row by row.
type Rows struct {
	*sql.Rows
	fieldMapFunc FieldMapFunc
}

// ScanMap populates the current row of data into a NullStringMap.
// Note that the NullStringMap must not be nil, or it will panic.
// The NullStringMap will be populated using column names as keys and their values as
// the corresponding element values.
func (r *Rows) ScanMap(a NullStringMap) error { _ = "STUB: not implemented"; return nil }

// ScanStruct populates the current row of data into a struct.
// The struct must be given as a pointer.
//
// ScanStruct associates struct fields with DB table columns through a field mapping function.
// It populates a struct field with the data of its associated column.
// Note that only exported struct fields will be populated.
//
// By default, DefaultFieldMapFunc() is used to map struct fields to table columns.
// This function separates each word in a field name with a underscore and turns every letter into lower case.
// For example, "LastName" is mapped to "last_name", "MyID" is mapped to "my_id", and so on.
// To change the default behavior, set DB.FieldMapper with your custom mapping function.
// You may also set Query.FieldMapper to change the behavior for particular queries.
func (r *Rows) ScanStruct(a interface{}) error { _ = "STUB: not implemented"; return nil }

// all populates all rows of query result into a slice of struct or NullStringMap.
// Note that the slice must be given as a pointer.
func (r *Rows) all(slice interface{}) error { _ = "STUB: not implemented"; return nil }

// create an empty slice

// column populates the given slice with the first column of the query result.
// Note that the slice must be given as a pointer.
func (r *Rows) column(slice interface{}) error { _ = "STUB: not implemented"; return nil }

// one populates a single row of query result into a struct or a NullStringMap.
// Note that if a struct is given, it should be a pointer.
func (r *Rows) one(a interface{}) error { _ = "STUB: not implemented"; return nil }

// pointer to map

// row populates a single row of query result into a list of variables.
func (r *Rows) row(a ...interface{}) error { _ = "STUB: not implemented"; return nil }
