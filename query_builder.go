// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

import (
	"regexp"
)

// QueryBuilder builds different clauses for a SELECT SQL statement.
type QueryBuilder interface {
	// BuildSelect generates a SELECT clause from the given selected column names.
	BuildSelect(cols []string, distinct bool, option string) string
	// BuildFrom generates a FROM clause from the given tables.
	BuildFrom(tables []string) string
	// BuildGroupBy generates a GROUP BY clause from the given group-by columns.
	BuildGroupBy(cols []string) string
	// BuildJoin generates a JOIN clause from the given join information.
	BuildJoin([]JoinInfo, Params) string
	// BuildWhere generates a WHERE clause from the given expression.
	BuildWhere(Expression, Params) string
	// BuildHaving generates a HAVING clause from the given expression.
	BuildHaving(Expression, Params) string
	// BuildOrderByAndLimit generates the ORDER BY and LIMIT clauses.
	BuildOrderByAndLimit(string, []string, int64, int64) string
	// BuildUnion generates a UNION clause from the given union information.
	BuildUnion([]UnionInfo, Params) string
}

// BaseQueryBuilder provides a basic implementation of QueryBuilder.
type BaseQueryBuilder struct {
	db *DB
}

var _ QueryBuilder = &BaseQueryBuilder{}

// NewBaseQueryBuilder creates a new BaseQueryBuilder instance.
func NewBaseQueryBuilder(db *DB) *BaseQueryBuilder { _ = "STUB: not implemented"; return nil }

// DB returns the DB instance associated with the query builder.
func (q *BaseQueryBuilder) DB() *DB {
	_ = "STUB: not implemented"

	// the regexp for columns and tables.
	return nil
}

var selectRegex = regexp.MustCompile(`(?i:\s+as\s+|\s+)([\w\-_\.]+)$`)

// BuildSelect generates a SELECT clause from the given selected column names.
func (q *BaseQueryBuilder) BuildSelect(cols []string, distinct bool, option string) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildFrom generates a FROM clause from the given tables.
func (q *BaseQueryBuilder) BuildFrom(tables []string) string { _ = "STUB: not implemented"; return "" }

// BuildJoin generates a JOIN clause from the given join information.
func (q *BaseQueryBuilder) BuildJoin(joins []JoinInfo, params Params) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildWhere generates a WHERE clause from the given expression.
func (q *BaseQueryBuilder) BuildWhere(e Expression, params Params) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildHaving generates a HAVING clause from the given expression.
func (q *BaseQueryBuilder) BuildHaving(e Expression, params Params) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildGroupBy generates a GROUP BY clause from the given group-by columns.
func (q *BaseQueryBuilder) BuildGroupBy(cols []string) string { _ = "STUB: not implemented"; return "" }

// BuildOrderByAndLimit generates the ORDER BY and LIMIT clauses.
func (q *BaseQueryBuilder) BuildOrderByAndLimit(sql string, cols []string, limit int64, offset int64) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildUnion generates a UNION clause from the given union information.
func (q *BaseQueryBuilder) BuildUnion(unions []UnionInfo, params Params) string {
	_ = "STUB: not implemented"
	return ""
}

var orderRegex = regexp.MustCompile(`\s+((?i)ASC|DESC)$`)

// BuildOrderBy generates the ORDER BY clause.
func (q *BaseQueryBuilder) BuildOrderBy(cols []string) string { _ = "STUB: not implemented"; return "" }

// BuildLimit generates the LIMIT clause.
func (q *BaseQueryBuilder) BuildLimit(limit int64, offset int64) string {
	_ = "STUB: not implemented"
	return ""
}

// most DBMS requires LIMIT when OFFSET is present
// 2^63 - 1

func (q *BaseQueryBuilder) quoteTableNameAndAlias(table string) string {
	_ = "STUB: not implemented"
	return ""
}
