package cmd

import (
	"fmt"

	"qr_generator/pkg/qr"

	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview [text]",
	Short: "Preview QR code in terminal",
	Long: `Preview QR code as ASCII art in terminal before generating image.
Example: qr preview "Hello World"`,
	Args: cobra.ExactArgs(1),
	RunE: runPreview,
}

func init() {
	rootCmd.AddCommand(previewCmd)
}

func runPreview(cmd *cobra.Command, args []string) error {
	config := qr.DefaultConfig()
	config.Text = args[0]

	generator, err := qr.NewGenerator(config)
	if err != nil {
		return err
	}

	fmt.Println(generator.Preview())
	return nil
} 
