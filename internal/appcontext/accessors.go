package appcontext

import (
	"context"
	"github.com/dmba/expense-tracker/internal/expense"
	"log"
)

func ExpenseServiceFromContext(ctx context.Context) *expense.Service {
	if l, ok := ctx.Value(ctxKeyExpenseService).(*expense.Service); ok {
		return l
	}
	log.Fatal("ExpenseService not found in context")
	return nil
}
