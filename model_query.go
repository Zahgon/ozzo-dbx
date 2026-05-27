package dbx

import (
	"context"
	"errors"
)

type (
	// TableModel is the interface that should be implemented by models which have unconventional table names.
	TableModel interface {
		TableName() string
	}

	// ModelQuery represents a query associated with a struct model.
	ModelQuery struct {
		db        *DB
		ctx       context.Context
		builder   Builder
		model     *structValue
		exclude   []string
		lastError error
	}
)

var (
	MissingPKError   = errors.New("missing primary key declaration")
	CompositePKError = errors.New("composite primary key is not supported")
)

func NewModelQuery(model interface{}, fieldMapFunc FieldMapFunc, db *DB, builder Builder) *ModelQuery {
	_ = "STUB: not implemented"
	return nil
}

// Context returns the context associated with the query.
func (q *ModelQuery) Context() context.Context {
	_ = "STUB: not implemented"

	// WithContext associates a context with the query.
	return *new(context.Context)
}

func (q *ModelQuery) WithContext(ctx context.Context) *ModelQuery {
	_ = "STUB: not implemented"
	return nil

	// Exclude excludes the specified struct fields from being inserted/updated into the DB table.
}

func (q *ModelQuery) Exclude(attrs ...string) *ModelQuery { _ = "STUB: not implemented"; return nil }

// Insert inserts a row in the table using the struct model associated with this query.
//
// By default, it inserts *all* public fields into the table, including those nil or empty ones.
// You may pass a list of the fields to this method to indicate that only those fields should be inserted.
// You may also call Exclude to exclude some fields from being inserted.
//
// If a model has an empty primary key, it is considered auto-incremental and the corresponding struct
// field will be filled with the generated primary key value after a successful insertion.
func (q *ModelQuery) Insert(attrs ...string) error { _ = "STUB: not implemented"; return nil }

// handle auto-incremental PK

func insertAndReturnPK(db *DB, query *Query, pkName string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// specially handle postgres (lib/pq) as it doesn't support LastInsertId

func isAutoInc(value interface{}) bool { _ = "STUB: not implemented"; return false }

// Update updates a row in the table using the struct model associated with this query.
// The row being updated has the same primary key as specified by the model.
//
// By default, it updates *all* public fields in the table, including those nil or empty ones.
// You may pass a list of the fields to this method to indicate that only those fields should be updated.
// You may also call Exclude to exclude some fields from being updated.
func (q *ModelQuery) Update(attrs ...string) error { _ = "STUB: not implemented"; return nil }

// Delete deletes a row in the table using the primary key specified by the struct model associated with this query.
func (q *ModelQuery) Delete() error { _ = "STUB: not implemented"; return nil }
