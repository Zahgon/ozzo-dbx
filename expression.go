// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbx

// Expression represents a DB expression that can be embedded in a SQL statement.
type Expression interface {
	// Build converts an expression into a SQL fragment.
	// If the expression contains binding parameters, they will be added to the given Params.
	Build(*DB, Params) string
}

// HashExp represents a hash expression.
//
// A hash expression is a map whose keys are DB column names which need to be filtered according
// to the corresponding values. For example, HashExp{"level": 2, "dept": 10} will generate
// the SQL: "level"=2 AND "dept"=10.
//
// HashExp also handles nil values and slice values. For example, HashExp{"level": []interface{}{1, 2}, "dept": nil}
// will generate: "level" IN (1, 2) AND "dept" IS NULL.
type HashExp map[string]interface{}

// NewExp generates an expression with the specified SQL fragment and the optional binding parameters.
func NewExp(e string, params ...Params) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// Not generates a NOT expression which prefixes "NOT" to the specified expression.
func Not(e Expression) Expression {
	_ = "STUB: not implemented"

	// And generates an AND expression which concatenates the given expressions with "AND".
	return *new(Expression)
}

func And(exps ...Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Or generates an OR expression which concatenates the given expressions with "OR".
func Or(exps ...Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// In generates an IN expression for the specified column and the list of allowed values.
// If values is empty, a SQL "0=1" will be generated which represents a false expression.
func In(col string, values ...interface{}) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// NotIn generates an NOT IN expression for the specified column and the list of disallowed values.
// If values is empty, an empty string will be returned indicating a true expression.
func NotIn(col string, values ...interface{}) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// DefaultLikeEscape specifies the default special character escaping for LIKE expressions
// The strings at 2i positions are the special characters to be escaped while those at 2i+1 positions
// are the corresponding escaped versions.
var DefaultLikeEscape = []string{"\\", "\\\\", "%", "\\%", "_", "\\_"}

// Like generates a LIKE expression for the specified column and the possible strings that the column should be like.
// If multiple values are present, the column should be like *all* of them. For example, Like("name", "key", "word")
// will generate a SQL expression: "name" LIKE "%key%" AND "name" LIKE "%word%".
//
// By default, each value will be surrounded by "%" to enable partial matching. If a value contains special characters
// such as "%", "\", "_", they will also be properly escaped.
//
// You may call Escape() and/or Match() to change the default behavior. For example, Like("name", "key").Match(false, true)
// generates "name" LIKE "key%".
func Like(col string, values ...string) *LikeExp { _ = "STUB: not implemented"; return nil }

// NotLike generates a NOT LIKE expression.
// For example, NotLike("name", "key", "word") will generate a SQL expression:
// "name" NOT LIKE "%key%" AND "name" NOT LIKE "%word%". Please see Like() for more details.
func NotLike(col string, values ...string) *LikeExp { _ = "STUB: not implemented"; return nil }

// OrLike generates an OR LIKE expression.
// This is similar to Like() except that the column should be like one of the possible values.
// For example, OrLike("name", "key", "word") will generate a SQL expression:
// "name" LIKE "%key%" OR "name" LIKE "%word%". Please see Like() for more details.
func OrLike(col string, values ...string) *LikeExp { _ = "STUB: not implemented"; return nil }

// OrNotLike generates an OR NOT LIKE expression.
// For example, OrNotLike("name", "key", "word") will generate a SQL expression:
// "name" NOT LIKE "%key%" OR "name" NOT LIKE "%word%". Please see Like() for more details.
func OrNotLike(col string, values ...string) *LikeExp { _ = "STUB: not implemented"; return nil }

// Exists generates an EXISTS expression by prefixing "EXISTS" to the given expression.
func Exists(exp Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// NotExists generates an EXISTS expression by prefixing "NOT EXISTS" to the given expression.
func NotExists(exp Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Between generates a BETWEEN expression.
// For example, Between("age", 10, 30) generates: "age" BETWEEN 10 AND 30
func Between(col string, from, to interface{}) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// NotBetween generates a NOT BETWEEN expression.
// For example, NotBetween("age", 10, 30) generates: "age" NOT BETWEEN 10 AND 30
func NotBetween(col string, from, to interface{}) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// Exp represents an expression with a SQL fragment and a list of optional binding parameters.
type Exp struct {
	e      string
	params Params
}

// Build converts an expression into a SQL fragment.
func (e *Exp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// Build converts an expression into a SQL fragment.
func (e HashExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// ensure the hash exp generates the same SQL for different runs

// NotExp represents an expression that should prefix "NOT" to a specified expression.
type NotExp struct {
	e Expression
}

// Build converts an expression into a SQL fragment.
func (e *NotExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// AndOrExp represents an expression that concatenates multiple expressions using either "AND" or "OR".
type AndOrExp struct {
	exps []Expression
	op   string
}

// Build converts an expression into a SQL fragment.
func (e *AndOrExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// InExp represents an "IN" or "NOT IN" expression.
type InExp struct {
	col    string
	values []interface{}
	not    bool
}

// Build converts an expression into a SQL fragment.
func (e *InExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// LikeExp represents a variant of LIKE expressions.
type LikeExp struct {
	or          bool
	left, right bool
	col         string
	values      []string
	escape      []string

	// Like stores the LIKE operator. It can be "LIKE", "NOT LIKE".
	// It may also be customized as something like "ILIKE".
	Like string
}

// Escape specifies how a LIKE expression should be escaped.
// Each string at position 2i represents a special character and the string at position 2i+1 is
// the corresponding escaped version.
func (e *LikeExp) Escape(chars ...string) *LikeExp { _ = "STUB: not implemented"; return nil }

// Match specifies whether to do wildcard matching on the left and/or right of given strings.
func (e *LikeExp) Match(left, right bool) *LikeExp { _ = "STUB: not implemented"; return nil }

// Build converts an expression into a SQL fragment.
func (e *LikeExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// ExistsExp represents an EXISTS or NOT EXISTS expression.
type ExistsExp struct {
	exp Expression
	not bool
}

// Build converts an expression into a SQL fragment.
func (e *ExistsExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }

// BetweenExp represents a BETWEEN or a NOT BETWEEN expression.
type BetweenExp struct {
	col      string
	from, to interface{}
	not      bool
}

// Build converts an expression into a SQL fragment.
func (e *BetweenExp) Build(db *DB, params Params) string { _ = "STUB: not implemented"; return "" }
