// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

import (
	"context"
	"database/sql"
)

// Params represents a list of parameter values to be bound to a SQL statement.
// The map keys are the parameter names while the map values are the corresponding parameter values.
type Params map[string]interface{}

// Executor prepares, executes, or queries a SQL statement.
type Executor interface {
	// Exec executes a SQL statement
	Exec(query string, args ...interface{}) (sql.Result, error)
	// ExecContext executes a SQL statement with the given context
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	// Query queries a SQL statement
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryContext queries a SQL statement with the given context
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// Prepare creates a prepared statement
	Prepare(query string) (*sql.Stmt, error)
}

// Query represents a SQL statement to be executed.
type Query struct {
	executor Executor

	sql, rawSQL  string
	placeholders []string
	params       Params

	stmt *sql.Stmt
	ctx  context.Context

	// FieldMapper maps struct field names to DB column names.
	FieldMapper FieldMapFunc
	// LastError contains the last error (if any) of the query.
	// LastError is cleared by Execute(), Row(), Rows(), One(), and All().
	LastError error
	// LogFunc is used to log the SQL statement being executed.
	LogFunc LogFunc
	// PerfFunc is used to log the SQL execution time. It is ignored if nil.
	// Deprecated: Please use QueryLogFunc and ExecLogFunc instead.
	PerfFunc PerfFunc
	// QueryLogFunc is called each time when performing a SQL query that returns data.
	QueryLogFunc QueryLogFunc
	// ExecLogFunc is called each time when a SQL statement is executed.
	ExecLogFunc ExecLogFunc
}

// NewQuery creates a new Query with the given SQL statement.
func NewQuery(db *DB, executor Executor, sql string) *Query { _ = "STUB: not implemented"; return nil }

// SQL returns the original SQL used to create the query.
// The actual SQL (RawSQL) being executed is obtained by replacing the named
// parameter placeholders with anonymous ones.
func (q *Query) SQL() string {
	_ = "STUB: not implemented"

	// Context returns the context associated with the query.
	return ""
}

func (q *Query) Context() context.Context {
	_ = "STUB: not implemented"

	// WithContext associates a context with the query.
	return *new(context.Context)
}

func (q *Query) WithContext(ctx context.Context) *Query { _ = "STUB: not implemented"; return nil }

// logSQL returns the SQL statement with parameters being replaced with the actual values.
// The result is only for logging purpose and should not be used to execute.
func (q *Query) logSQL() string { _ = "STUB: not implemented"; return "" }

// Params returns the parameters to be bound to the SQL statement represented by this query.
func (q *Query) Params() Params {
	_ = "STUB: not implemented"

	// Prepare creates a prepared statement for later queries or executions.
	// Close() should be called after finishing all queries.
	return *new(Params)
}

func (q *Query) Prepare() *Query { _ = "STUB: not implemented"; return nil }

// Close closes the underlying prepared statement.
// Close does nothing if the query has not been prepared before.
func (q *Query) Close() error { _ = "STUB: not implemented"; return nil }

// Bind sets the parameters that should be bound to the SQL statement.
// The parameter placeholders in the SQL statement are in the format of "{:ParamName}".
func (q *Query) Bind(params Params) *Query { _ = "STUB: not implemented"; return nil }

// Execute executes the SQL statement without retrieving data.
func (q *Query) Execute() (result sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// One executes the SQL statement and populates the first row of the result into a struct or NullStringMap.
// Refer to Rows.ScanStruct() and Rows.ScanMap() for more details on how to specify
// the variable to be populated.
// Note that when the query has no rows in the result set, an sql.ErrNoRows will be returned.
func (q *Query) One(a interface{}) error { _ = "STUB: not implemented"; return nil }

// All executes the SQL statement and populates all the resulting rows into a slice of struct or NullStringMap.
// The slice must be given as a pointer. Each slice element must be either a struct or a NullStringMap.
// Refer to Rows.ScanStruct() and Rows.ScanMap() for more details on how each slice element can be.
// If the query returns no row, the slice will be an empty slice (not nil).
func (q *Query) All(slice interface{}) error { _ = "STUB: not implemented"; return nil }

// Row executes the SQL statement and populates the first row of the result into a list of variables.
// Note that the number of the variables should match to that of the columns in the query result.
// Note that when the query has no rows in the result set, an sql.ErrNoRows will be returned.
func (q *Query) Row(a ...interface{}) error { _ = "STUB: not implemented"; return nil }

// Column executes the SQL statement and populates the first column of the result into a slice.
// Note that the parameter must be a pointer to a slice.
func (q *Query) Column(a interface{}) error { _ = "STUB: not implemented"; return nil }

// Rows executes the SQL statement and returns a Rows object to allow retrieving data row by row.
func (q *Query) Rows() (rows *Rows, err error) { _ = "STUB: not implemented"; return nil, nil }

// replacePlaceholders converts a list of named parameters into a list of anonymous parameters.
func replacePlaceholders(placeholders []string, params Params) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
