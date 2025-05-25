package cmd

import (
	"fmt"
	"github.com/dmba/expense-tracker/internal/appcontext"

	"github.com/spf13/cobra"
)

var (
	ErrInvalidID = fmt.Errorf("invalid ID: must be greater than zero")
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			cmd.PrintErrf("Error getting \"id\" flag: %v\n", err)
			return
		}

		if err := deleteHandler(cmd, id); err != nil {
			cmd.PrintErrf("Error deleting expense: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("id", "i", 0, "Expense ID to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		panic(err)
	}
}

func deleteHandler(cmd *cobra.Command, id int) error {
	if id <= 0 {
		return ErrInvalidID
	}

	service := appcontext.ExpenseServiceFromContext(cmd.Context())

	if err := service.DeleteExpense(id); err != nil {
		return err
	}
	cmd.Println("Expense deleted successfully")

	return nil
}
