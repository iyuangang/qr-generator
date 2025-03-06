package qr

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/skip2/go-qrcode"
)

// Config holds the QR code generation configuration
type Config struct {
	Text           string
	Output         string
	Size           int
	Border         int
	ForegroundColor color.Color
	BackgroundColor color.Color
	ErrorCorrection qrcode.RecoveryLevel
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		Size:            10,
		Border:          4,
		ForegroundColor: color.Black,
		BackgroundColor: color.White,
		ErrorCorrection: qrcode.Medium,
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if strings.TrimSpace(c.Text) == "" {
		return fmt.Errorf("text content cannot be empty")
	}
	if c.Size < 1 || c.Size > 40 {
		return fmt.Errorf("size must be between 1 and 40")
	}
	if c.Border < 0 {
		return fmt.Errorf("border must be non-negative")
	}
	return nil
} 
