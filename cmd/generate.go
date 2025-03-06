package cmd

import (
	"fmt"
	"image/color"
	"os"

	"qr_generator/pkg/qr"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

var (
	output         string
	size           int
	border         int
	fgColor        string
	bgColor        string
	errorLevel     string
)

var generateCmd = &cobra.Command{
	Use:   "generate [text/file]",
	Short: "Generate QR code from text or file",
	Long: `Generate a QR code from text input or file content.
Examples:
  qr generate "Hello World" -o hello.png
  qr generate input.txt -o output.png -s 15 -b 2
  qr generate "Hello" --fg-color black --bg-color white`,
	Args: cobra.ExactArgs(1),
	RunE: runGenerate,
}

func init() {
	generateCmd.Flags().StringVarP(&output, "output", "o", "qr_code.png",
		"output image file path")
	generateCmd.Flags().IntVarP(&size, "size", "s", 10,
		"QR code size (1-40)")
	generateCmd.Flags().IntVarP(&border, "border", "b", 4,
		"border size around QR code")
	generateCmd.Flags().StringVar(&fgColor, "fg-color", "black",
		"foreground color (black/white/blue/red/green)")
	generateCmd.Flags().StringVar(&bgColor, "bg-color", "white",
		"background color (black/white/blue/red/green)")
	generateCmd.Flags().StringVarP(&errorLevel, "error-level", "e", "medium",
		"error correction level (low/medium/high/highest)")

	generateCmd.MarkFlagRequired("output")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Create config with default values
	config := qr.DefaultConfig()
	config.Output = output
	config.Size = size
	config.Border = border

	// Set colors
	fg, err := parseColor(fgColor)
	if err != nil {
		return fmt.Errorf("invalid foreground color: %v", err)
	}
	config.ForegroundColor = fg

	bg, err := parseColor(bgColor)
	if err != nil {
		return fmt.Errorf("invalid background color: %v", err)
	}
	config.BackgroundColor = bg

	// Set error correction level
	level, err := parseErrorLevel(errorLevel)
	if err != nil {
		return fmt.Errorf("invalid error correction level: %v", err)
	}
	config.ErrorCorrection = level

	// Handle input
	if fileInfo, err := os.Stat(args[0]); err == nil && !fileInfo.IsDir() {
		content, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}
		config.Text = string(content)
	} else {
		config.Text = args[0]
	}

	// Create generator
	generator, err := qr.NewGenerator(config)
	if err != nil {
		return err
	}

	// Generate QR code
	if err := generator.Generate(); err != nil {
		return err
	}

	fmt.Printf("QR code generated successfully: %s\n", config.Output)
	return nil
}

func parseColor(name string) (color.Color, error) {
	colors := map[string]color.Color{
		"black": color.Black,
		"white": color.White,
		"blue":  color.RGBA{0, 0, 255, 255},
		"red":   color.RGBA{255, 0, 0, 255},
		"green": color.RGBA{0, 255, 0, 255},
	}

	c, ok := colors[name]
	if !ok {
		return nil, fmt.Errorf("unsupported color: %s", name)
	}
	return c, nil
}

func parseErrorLevel(level string) (qrcode.RecoveryLevel, error) {
	levels := map[string]qrcode.RecoveryLevel{
		"low":     qrcode.Low,
		"medium":  qrcode.Medium,
		"high":    qrcode.High,
		"highest": qrcode.Highest,
	}

	l, ok := levels[level]
	if !ok {
		return 0, fmt.Errorf("unsupported error level: %s", level)
	}
	return l, nil
} 
