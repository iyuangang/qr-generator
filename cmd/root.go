package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qr",
	Short: "QR code generator CLI tool",
	Long: `A command line tool for generating QR codes from text or files.
Supports customization of size, border, and output format.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add global flags here if needed
	rootCmd.AddCommand(generateCmd)
} 
