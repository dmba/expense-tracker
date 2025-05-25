package cmd

import (
	"fmt"
	"github.com/dmba/expense-tracker/internal/appcontext"
	"github.com/dmba/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
	"time"
)

var (
	ErrInvalidMonth = fmt.Errorf("invalid month: must be between 1 and 12")
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Total expenses and income summary for a given period",
	Run: func(cmd *cobra.Command, args []string) {
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			cmd.PrintErrf("Error getting \"month\" flag: %v\n", err)
			return
		}
		if err := summaryHandler(cmd, month); err != nil {
			cmd.PrintErrf("Error generating summary: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntP("month", "m", 0, "Month for summary (1-12)")
}

func summaryHandler(cmd *cobra.Command, month int) error {
	service := appcontext.ExpenseServiceFromContext(cmd.Context())

	switch {
	case month == 0:
		return totalSummary(cmd, service)
	case month >= 1 && month <= 12:
		return totalSummaryForMonth(cmd, service, time.Month(month))
	default:
		return ErrInvalidMonth
	}
}

func totalSummary(cmd *cobra.Command, service *expense.Service) error {
	summary, err := service.Summary()
	if err != nil {
		return err
	}
	cmd.Printf("Total expenses: %s\n", summary.String())
	return nil
}

func totalSummaryForMonth(cmd *cobra.Command, service *expense.Service, month time.Month) error {
	summary, err := service.SummaryByMonth(month)
	if err != nil {
		return err
	}

	cmd.Printf("Total expenses for %v: %s\n", month, summary.String())
	return nil
}
