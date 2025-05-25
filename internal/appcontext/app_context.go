package appcontext

import (
	"context"
	"github.com/dmba/expense-tracker/internal/expense"
)

type AppContext struct {
	context.Context
}

func NewAppContext(ctx context.Context, service *expense.Service) *AppContext {
	ctx = context.WithValue(ctx, ctxKeyExpenseService, service)
	return &AppContext{
		Context: ctx,
	}
}
