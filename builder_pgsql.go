// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

// PgsqlBuilder is the builder for PostgreSQL databases.
type PgsqlBuilder struct {
	*BaseBuilder
	qb *BaseQueryBuilder
}

var _ Builder = &PgsqlBuilder{}

// NewPgsqlBuilder creates a new PgsqlBuilder instance.
func NewPgsqlBuilder(db *DB, executor Executor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// Select returns a new SelectQuery object that can be used to build a SELECT statement.
// The parameters to this method should be the list column names to be selected.
// A column name may have an optional alias name. For example, Select("id", "my_name AS name").
func (b *PgsqlBuilder) Select(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Model returns a new ModelQuery object that can be used to perform model-based DB operations.
// The model passed to this method should be a pointer to a model struct.
func (b *PgsqlBuilder) Model(model interface{}) *ModelQuery { _ = "STUB: not implemented"; return nil }

// GeneratePlaceholder generates an anonymous parameter placeholder with the given parameter ID.
func (b *PgsqlBuilder) GeneratePlaceholder(i int) string { _ = "STUB: not implemented"; return "" }

// QueryBuilder returns the query builder supporting the current DB.
func (b *PgsqlBuilder) QueryBuilder() QueryBuilder {
	_ = "STUB: not implemented"

	// Upsert creates a Query that represents an UPSERT SQL statement.
	// Upsert inserts a row into the table if the primary key or unique index is not found.
	// Otherwise it will update the row with the new values.
	// The keys of cols are the column names, while the values of cols are the corresponding column
	// values to be inserted.
	return *new(QueryBuilder)
}

func (b *PgsqlBuilder) Upsert(table string, cols Params, constraints ...string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// DropIndex creates a Query that can be used to remove the named index from a table.
func (b *PgsqlBuilder) DropIndex(table, name string) *Query { _ = "STUB: not implemented"; return nil }

// RenameTable creates a Query that can be used to rename a table.
func (b *PgsqlBuilder) RenameTable(oldName, newName string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// AlterColumn creates a Query that can be used to change the definition of a table column.
func (b *PgsqlBuilder) AlterColumn(table, col, typ string) *Query {
	_ = "STUB: not implemented"
	return nil
}
