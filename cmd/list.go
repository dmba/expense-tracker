package cmd

import (
	"fmt"
	"github.com/dmba/expense-tracker/internal/appcontext"
	"github.com/dmba/expense-tracker/pkg/models"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listHandler(cmd); err != nil {
			cmd.PrintErrf("Error listing expenses: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listHandler(cmd *cobra.Command) error {
	service := appcontext.ExpenseServiceFromContext(cmd.Context())

	expenses, err := service.ListExpenses()
	if err != nil {
		return err
	}
	if err := printExpenses(expenses); err != nil {
		return err
	}

	return nil
}

func printExpenses(expenses []models.Expense) error {
	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	if _, err := fmt.Fprintln(w, "ID\tDate\tDescription\tAmount"); err != nil {
		return err
	}

	for _, task := range expenses {
		_, err := fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\n",
			task.ID,
			task.Date.Format(time.DateOnly),
			task.Description,
			task.Amount.String(),
		)
		if err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
