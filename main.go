package main

import (
	"context"
	"github.com/dmba/expense-tracker/cmd"
	"github.com/dmba/expense-tracker/internal/appcontext"
	"github.com/dmba/expense-tracker/internal/expense"
)

const (
	filepath = ".expense-tracker.csv"
)

func main() {
	app := appcontext.NewAppContext(
		context.Background(),
		expense.NewService(filepath),
	)

	cmd.Execute(app)
}
