package cursor_pagination

import "cloud.google.com/go/spanner/spansql"

type CursorPagination struct {
}

type ICursorPagination interface {
	ToOrder(desc bool, pk, field spansql.Expr) ([2]spansql.Order, error)
	ToPagination(desc bool, pk, pkValue, field spansql.Expr, fieldValue map[spansql.Expr]spansql.Expr) (spansql.BoolExpr, error)
}
