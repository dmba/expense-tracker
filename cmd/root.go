package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A a simple expense tracker application to manage your finances",
	Args:  cobra.MinimumNArgs(1),
}

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
