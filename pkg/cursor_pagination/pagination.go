package cursor_pagination

import (
	"cloud.google.com/go/spanner/spansql"
	"fmt"
)

func (pagination *CursorPagination) ToPagination(desc bool, pk, pkValue, field spansql.Expr, fieldValue map[spansql.Expr]spansql.Expr) (spansql.BoolExpr, error) {
	if _, exists := fieldValue[field]; !exists {
		return nil, fmt.Errorf("fieldValue for %s not found", field)
	}

	var opLhs spansql.ComparisonOperator
	var opRhs spansql.ComparisonOperator

	switch desc {
	case true:
		opLhs = spansql.Le
		opRhs = spansql.Lt
	case false:
		opLhs = spansql.Ge
		opRhs = spansql.Gt
	}

	return spansql.LogicalOp{
		Op: spansql.And,
		LHS: spansql.ComparisonOp{
			Op:  opLhs,
			LHS: field,
			RHS: fieldValue[field],
		},
		RHS: spansql.Paren{
			Expr: spansql.LogicalOp{
				Op:  spansql.Or,
				LHS: spansql.ComparisonOp{Op: opRhs, LHS: field, RHS: fieldValue[field]},
				RHS: spansql.ComparisonOp{Op: opRhs, LHS: pk, RHS: pkValue},
			},
		},
	}, nil
}
