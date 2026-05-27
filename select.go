// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

import (
	"context"
)

// SelectQuery represents a DB-agnostic SELECT query.
// It can be built into a DB-specific query by calling the Build() method.
type SelectQuery struct {
	// FieldMapper maps struct field names to DB column names.
	FieldMapper FieldMapFunc
	// TableMapper maps structs to DB table names.
	TableMapper TableMapFunc

	builder Builder
	ctx     context.Context

	selects      []string
	distinct     bool
	selectOption string
	from         []string
	where        Expression
	join         []JoinInfo
	orderBy      []string
	groupBy      []string
	having       Expression
	union        []UnionInfo
	limit        int64
	offset       int64
	params       Params
}

// JoinInfo contains the specification for a JOIN clause.
type JoinInfo struct {
	Join  string
	Table string
	On    Expression
}

// UnionInfo contains the specification for a UNION clause.
type UnionInfo struct {
	All   bool
	Query *Query
}

// NewSelectQuery creates a new SelectQuery instance.
func NewSelectQuery(builder Builder, db *DB) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Context returns the context associated with the query.
func (q *SelectQuery) Context() context.Context {
	_ = "STUB: not implemented"

	// WithContext associates a context with the query.
	return *new(context.Context)
}

func (q *SelectQuery) WithContext(ctx context.Context) *SelectQuery {
	_ = "STUB: not implemented"
	return nil

	// Select specifies the columns to be selected.
	// Column names will be automatically quoted.
}

func (s *SelectQuery) Select(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndSelect adds additional columns to be selected.
// Column names will be automatically quoted.
func (s *SelectQuery) AndSelect(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Distinct specifies whether to select columns distinctively.
// By default, distinct is false.
func (s *SelectQuery) Distinct(v bool) *SelectQuery { _ = "STUB: not implemented"; return nil }

// SelectOption specifies additional option that should be append to "SELECT".
func (s *SelectQuery) SelectOption(option string) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// From specifies which tables to select from.
// Table names will be automatically quoted.
func (s *SelectQuery) From(tables ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Where specifies the WHERE condition.
func (s *SelectQuery) Where(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndWhere concatenates a new WHERE condition with the existing one (if any) using "AND".
func (s *SelectQuery) AndWhere(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// OrWhere concatenates a new WHERE condition with the existing one (if any) using "OR".
func (s *SelectQuery) OrWhere(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Join specifies a JOIN clause.
// The "typ" parameter specifies the JOIN type (e.g. "INNER JOIN", "LEFT JOIN").
func (s *SelectQuery) Join(typ string, table string, on Expression) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// InnerJoin specifies an INNER JOIN clause.
// This is a shortcut method for Join.
func (s *SelectQuery) InnerJoin(table string, on Expression) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// LeftJoin specifies a LEFT JOIN clause.
// This is a shortcut method for Join.
func (s *SelectQuery) LeftJoin(table string, on Expression) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// RightJoin specifies a RIGHT JOIN clause.
// This is a shortcut method for Join.
func (s *SelectQuery) RightJoin(table string, on Expression) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// OrderBy specifies the ORDER BY clause.
// Column names will be properly quoted. A column name can contain "ASC" or "DESC" to indicate its ordering direction.
func (s *SelectQuery) OrderBy(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndOrderBy appends additional columns to the existing ORDER BY clause.
// Column names will be properly quoted. A column name can contain "ASC" or "DESC" to indicate its ordering direction.
func (s *SelectQuery) AndOrderBy(cols ...string) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// GroupBy specifies the GROUP BY clause.
// Column names will be properly quoted.
func (s *SelectQuery) GroupBy(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndGroupBy appends additional columns to the existing GROUP BY clause.
// Column names will be properly quoted.
func (s *SelectQuery) AndGroupBy(cols ...string) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// Having specifies the HAVING clause.
func (s *SelectQuery) Having(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndHaving concatenates a new HAVING condition with the existing one (if any) using "AND".
func (s *SelectQuery) AndHaving(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// OrHaving concatenates a new HAVING condition with the existing one (if any) using "OR".
func (s *SelectQuery) OrHaving(e Expression) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Union specifies a UNION clause.
func (s *SelectQuery) Union(q *Query) *SelectQuery { _ = "STUB: not implemented"; return nil }

// UnionAll specifies a UNION ALL clause.
func (s *SelectQuery) UnionAll(q *Query) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Limit specifies the LIMIT clause.
// A negative limit means no limit.
func (s *SelectQuery) Limit(limit int64) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Offset specifies the OFFSET clause.
// A negative offset means no offset.
func (s *SelectQuery) Offset(offset int64) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Bind specifies the parameter values to be bound to the query.
func (s *SelectQuery) Bind(params Params) *SelectQuery { _ = "STUB: not implemented"; return nil }

// AndBind appends additional parameters to be bound to the query.
func (s *SelectQuery) AndBind(params Params) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Build builds the SELECT query and returns an executable Query object.
func (s *SelectQuery) Build() *Query { _ = "STUB: not implemented"; return nil }

// One executes the SELECT query and populates the first row of the result into the specified variable.
//
// If the query does not specify a "from" clause, the method will try to infer the name of the table
// to be selected from by calling getTableName() which will return either the variable type name
// or the TableName() method if the variable implements the TableModel interface.
//
// Note that when the query has no rows in the result set, an sql.ErrNoRows will be returned.
func (s *SelectQuery) One(a interface{}) error { _ = "STUB: not implemented"; return nil }

// Model selects the row with the specified primary key and populates the model with the row data.
//
// The model variable should be a pointer to a struct. If the query does not specify a "from" clause,
// it will use the model struct to determine which table to select data from. It will also use the model
// to infer the name of the primary key column. Only simple primary key is supported. For composite primary keys,
// please use Where() to specify the filtering condition.
func (s *SelectQuery) Model(pk, model interface{}) error { _ = "STUB: not implemented"; return nil }

// All executes the SELECT query and populates all rows of the result into a slice.
//
// Note that the slice must be passed in as a pointer.
//
// If the query does not specify a "from" clause, the method will try to infer the name of the table
// to be selected from by calling getTableName() which will return either the type name of the slice elements
// or the TableName() method if the slice element implements the TableModel interface.
func (s *SelectQuery) All(slice interface{}) error { _ = "STUB: not implemented"; return nil }

// Rows builds and executes the SELECT query and returns a Rows object for data retrieval purpose.
// This is a shortcut to SelectQuery.Build().Rows()
func (s *SelectQuery) Rows() (*Rows, error) { _ = "STUB: not implemented"; return nil, nil }

// Row builds and executes the SELECT query and populates the first row of the result into the specified variables.
// This is a shortcut to SelectQuery.Build().Row()
func (s *SelectQuery) Row(a ...interface{}) error { _ = "STUB: not implemented"; return nil }

// Column builds and executes the SELECT statement and populates the first column of the result into a slice.
// Note that the parameter must be a pointer to a slice.
// This is a shortcut to SelectQuery.Build().Column()
func (s *SelectQuery) Column(a interface{}) error { _ = "STUB: not implemented"; return nil }
