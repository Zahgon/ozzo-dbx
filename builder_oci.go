// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

// OciBuilder is the builder for Oracle databases.
type OciBuilder struct {
	*BaseBuilder
	qb *OciQueryBuilder
}

var _ Builder = &OciBuilder{}

// OciQueryBuilder is the query builder for Oracle databases.
type OciQueryBuilder struct {
	*BaseQueryBuilder
}

// NewOciBuilder creates a new OciBuilder instance.
func NewOciBuilder(db *DB, executor Executor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// Select returns a new SelectQuery object that can be used to build a SELECT statement.
// The parameters to this method should be the list column names to be selected.
// A column name may have an optional alias name. For example, Select("id", "my_name AS name").
func (b *OciBuilder) Select(cols ...string) *SelectQuery { _ = "STUB: not implemented"; return nil }

// Model returns a new ModelQuery object that can be used to perform model-based DB operations.
// The model passed to this method should be a pointer to a model struct.
func (b *OciBuilder) Model(model interface{}) *ModelQuery { _ = "STUB: not implemented"; return nil }

// GeneratePlaceholder generates an anonymous parameter placeholder with the given parameter ID.
func (b *OciBuilder) GeneratePlaceholder(i int) string { _ = "STUB: not implemented"; return "" }

// QueryBuilder returns the query builder supporting the current DB.
func (b *OciBuilder) QueryBuilder() QueryBuilder {
	_ = "STUB: not implemented"

	// DropIndex creates a Query that can be used to remove the named index from a table.
	return *new(QueryBuilder)
}

func (b *OciBuilder) DropIndex(table, name string) *Query { _ = "STUB: not implemented"; return nil }

// RenameTable creates a Query that can be used to rename a table.
func (b *OciBuilder) RenameTable(oldName, newName string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// AlterColumn creates a Query that can be used to change the definition of a table column.
func (b *OciBuilder) AlterColumn(table, col, typ string) *Query {
	_ = "STUB: not implemented"
	return nil
}

// BuildOrderByAndLimit generates the ORDER BY and LIMIT clauses.
func (q *OciQueryBuilder) BuildOrderByAndLimit(sql string, cols []string, limit int64, offset int64) string {
	_ = "STUB: not implemented"
	return ""
}
