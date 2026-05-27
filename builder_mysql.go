// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

import (
	"regexp"
)

// MysqlBuilder is the builder for MySQL databases.
type MysqlBuilder struct {
	*BaseBuilder
	qb *BaseQueryBuilder
}

var _ Builder = &MysqlBuilder{}

// NewMysqlBuilder creates a new MysqlBuilder instance.
func NewMysqlBuilder(db *DB, executor Executor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// QueryBuilder returns the query builder supporting the current DB.
func (b *MysqlBuilder) QueryBuilder() QueryBuilder {
	_ = "STUB: not implemented"

	// Select returns a new SelectQuery object that can be used to build a SELECT statement.
	// The parameters to this method should be the list column names to be selected.
	// A column name may have an optional alias name. For example, Select("id", "my_name AS name").
	return *new(QueryBuilder)
}

func (b *MysqlBuilder) Select(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Model returns a new ModelQuery object that can be used to perform model-based DB operations.
// The model passed to this method should be a pointer to a model struct.
func (b *MysqlBuilder) Model(model interface{}) *ModelQuery { _ = "STUB: not implemented"; return nil }

// QuoteSimpleTableName quotes a simple table name.
// A simple table name does not contain any schema prefix.
func (b *MysqlBuilder) QuoteSimpleTableName(s string) string { _ = "STUB: not implemented"; return "" }

// QuoteSimpleColumnName quotes a simple column name.
// A simple column name does not contain any table prefix.
func (b *MysqlBuilder) QuoteSimpleColumnName(s string) string { _ = "STUB: not implemented"; return "" }

// Upsert creates a Query that represents an UPSERT SQL statement.
// Upsert inserts a row into the table if the primary key or unique index is not found.
// Otherwise it will update the row with the new values.
// The keys of cols are the column names, while the values of cols are the corresponding column
// values to be inserted.
func (b *MysqlBuilder) Upsert(table string, cols Params, constraints ...string) *Query {
	_ = "STUB: not implemented"
	return nil
}

var mysqlColumnRegexp = regexp.MustCompile("(?m)^\\s*[`\"](.*?)[`\"]\\s+(.*?),?$")

// RenameColumn creates a Query that can be used to rename a column in a table.
func (b *MysqlBuilder) RenameColumn(table, oldName, newName string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// DropPrimaryKey creates a Query that can be used to remove the named primary key constraint from a table.
func (b *MysqlBuilder) DropPrimaryKey(table, name string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// DropForeignKey creates a Query that can be used to remove the named foreign key constraint from a table.
func (b *MysqlBuilder) DropForeignKey(table, name string) *Query {
	_ = "STUB: not implemented"
	return nil
}
