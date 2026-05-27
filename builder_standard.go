// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

// StandardBuilder is the builder that is used by DB for an unknown driver.
type StandardBuilder struct {
	*BaseBuilder
	qb *BaseQueryBuilder
}

var _ Builder = &StandardBuilder{}

// NewStandardBuilder creates a new StandardBuilder instance.
func NewStandardBuilder(db *DB, executor Executor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// QueryBuilder returns the query builder supporting the current DB.
func (b *StandardBuilder) QueryBuilder() QueryBuilder {
	_ = "STUB: not implemented"

	// Select returns a new SelectQuery object that can be used to build a SELECT statement.
	// The parameters to this method should be the list column names to be selected.
	// A column name may have an optional alias name. For example, Select("id", "my_name AS name").
	return *new(QueryBuilder)
}

func (b *StandardBuilder) Select(cols ...string) *SelectQuery {
	_ = "STUB: not implemented"
	return nil
}

// Model returns a new ModelQuery object that can be used to perform model-based DB operations.
// The model passed to this method should be a pointer to a model struct.
func (b *StandardBuilder) Model(model interface{}) *ModelQuery {
	_ = "STUB: not implemented"
	return nil
}
