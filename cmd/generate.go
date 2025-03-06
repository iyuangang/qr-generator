package cmd

import (
	"fmt"
	"os"

	"qr_generator/pkg/qr"

	"github.com/spf13/cobra"
)

var (
	output string
	size   int
	border int
)

var generateCmd = &cobra.Command{
	Use:   "generate [text/file]",
	Short: "Generate QR code from text or file",
	Long: `Generate a QR code from text input or file content.
Examples:
  qr generate "Hello World" -o hello.png
  qr generate input.txt -o output.png -s 15 -b 2`,
	Args: cobra.ExactArgs(1),
	RunE: runGenerate,
}

func init() {
	// Add local flags
	generateCmd.Flags().StringVarP(&output, "output", "o", "qr_code.png",
		"output image file path")
	generateCmd.Flags().IntVarP(&size, "size", "s", 10,
		"QR code size (1-40)")
	generateCmd.Flags().IntVarP(&border, "border", "b", 4,
		"border size around QR code")

	// Mark required flags
	generateCmd.MarkFlagRequired("output")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Validate parameters
	if err := validateParams(); err != nil {
		return err
	}

	// Create config
	config := &qr.Config{
		Text:   args[0],
		Output: output,
		Size:   size,
		Border: border,
	}

	// Check if input is a file
	if fileInfo, err := os.Stat(config.Text); err == nil && !fileInfo.IsDir() {
		content, err := os.ReadFile(config.Text)
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}
		config.Text = string(content)
	}

	// Generate QR code
	if err := qr.Generate(config); err != nil {
		return err
	}

	fmt.Printf("QR code generated successfully: %s\n", config.Output)
	return nil
}

func validateParams() error {
	if size < 1 || size > 40 {
		return fmt.Errorf("size must be between 1 and 40")
	}
	if border < 0 {
		return fmt.Errorf("border must be non-negative")
	}
	return nil
} 
