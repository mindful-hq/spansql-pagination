package cursor_pagination

import (
	"cloud.google.com/go/spanner/spansql"
)

func (pagination *CursorPagination) ToOrder(desc bool, pk, field spansql.Expr) ([2]spansql.Order, error) {
	return [2]spansql.Order{
		{Expr: field, Desc: desc},
		{Expr: pk, Desc: desc},
	}, nil
}
