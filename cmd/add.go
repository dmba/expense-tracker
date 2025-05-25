package cmd

import (
	"errors"
	"github.com/dmba/expense-tracker/internal/appcontext"
	"github.com/spf13/cobra"
)

var (
	ErrEmptyDescription = errors.New("--description cannot be empty")
	ErrInvalidAmount    = errors.New("--amount must be greater than zero")
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			cmd.PrintErrf("Error getting \"description\" flag: %v\n", err)
			return
		}
		amount, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			cmd.PrintErrf("Error getting \"amount\" flag: %v\n", err)
			return
		}

		if err := addHandler(cmd, description, amount); err != nil {
			cmd.PrintErrf("Error adding expense: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "Expense description")
	addCmd.Flags().Float64P("amount", "a", 0, "Expense amount in $")

	addCmd.MarkFlagsRequiredTogether("description", "amount")
}

func addHandler(cmd *cobra.Command, description string, amount float64) error {
	if description == "" {
		return ErrEmptyDescription
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	service := appcontext.ExpenseServiceFromContext(cmd.Context())

	expense, err := service.AddExpense(description, amount)
	if err != nil {
		return err
	}
	cmd.Printf("Expense added successfully (ID: %d)\n", expense.ID)

	return nil
}
