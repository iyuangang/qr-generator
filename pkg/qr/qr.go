package qr

import (
	"fmt"
	"strings"

	"github.com/skip2/go-qrcode"
)

// Config holds the QR code generation configuration
type Config struct {
	Text   string
	Output string
	Size   int
	Border int
}

// Generate creates a QR code with the given configuration
func Generate(config *Config) error {
	// Validate input
	if strings.TrimSpace(config.Text) == "" {
		return fmt.Errorf("text content cannot be empty")
	}

	// Calculate QR code size
	size := config.Size * config.Border

	// Generate QR code
	err := qrcode.WriteFile(config.Text, qrcode.Medium, size, config.Output)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %v", err)
	}

	return nil
} 
