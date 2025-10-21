package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "infrar",
	Short: "Infrar CLI - Multi-cloud infrastructure intelligence",
	Long: `Infrar is a command-line tool for managing multi-cloud infrastructure.

Write cloud-agnostic code once, deploy to any cloud provider.
Transform your code at deployment time for zero runtime overhead.`,
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
}
