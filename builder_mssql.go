// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

// MssqlBuilder is the builder for SQL Server databases.
type MssqlBuilder struct {
	*BaseBuilder
	qb *MssqlQueryBuilder
}

var _ Builder = &MssqlBuilder{}

// MssqlQueryBuilder is the query builder for SQL Server databases.
type MssqlQueryBuilder struct {
	*BaseQueryBuilder
}

// NewMssqlBuilder creates a new MssqlBuilder instance.
func NewMssqlBuilder(db *DB, executor Executor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// QueryBuilder returns the query builder supporting the current DB.
func (b *MssqlBuilder) QueryBuilder() QueryBuilder {
	_ = "STUB: not implemented"

	// Select returns a new SelectQuery object that can be used to build a SELECT statement.
	// The parameters to this method should be the list column names to be selected.
	// A column name may have an optional alias name. For example, Select("id", "my_name AS name").
	return *new(QueryBuilder)
}

func (b *MssqlBuilder) Select(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Model returns a new ModelQuery object that can be used to perform model-based DB operations.
// The model passed to this method should be a pointer to a model struct.
func (b *MssqlBuilder) Model(model interface{}) *ModelQuery { _ = "STUB: not implemented"; return nil }

// QuoteSimpleTableName quotes a simple table name.
// A simple table name does not contain any schema prefix.
func (b *MssqlBuilder) QuoteSimpleTableName(s string) string { _ = "STUB: not implemented"; return "" }

// QuoteSimpleColumnName quotes a simple column name.
// A simple column name does not contain any table prefix.
func (b *MssqlBuilder) QuoteSimpleColumnName(s string) string { _ = "STUB: not implemented"; return "" }

// RenameTable creates a Query that can be used to rename a table.
func (b *MssqlBuilder) RenameTable(oldName, newName string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// RenameColumn creates a Query that can be used to rename a column in a table.
func (b *MssqlBuilder) RenameColumn(table, oldName, newName string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// AlterColumn creates a Query that can be used to change the definition of a table column.
func (b *MssqlBuilder) AlterColumn(table, col, typ string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// BuildOrderByAndLimit generates the ORDER BY and LIMIT clauses.
func (q *MssqlQueryBuilder) BuildOrderByAndLimit(sql string, cols []string, limit int64, offset int64) string {
	_ = "STUB: not implemented"
	return ""
}

// only SQL SERVER 2012 or newer are supported by this method

// ORDER BY clause is required when FETCH and OFFSET are in the SQL

// http://technet.microsoft.com/en-us/library/gg699618.aspx
