package qr

import (
	"fmt"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

// Generator handles QR code generation
type Generator struct {
	config *Config
	qr     *qrcode.QRCode
}

// NewGenerator creates a new QR code generator
func NewGenerator(config *Config) (*Generator, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	qr, err := qrcode.New(config.Text, config.ErrorCorrection)
	if err != nil {
		return nil, fmt.Errorf("failed to create QR code: %v", err)
	}

	return &Generator{
		config: config,
		qr:     qr,
	}, nil
}

// Generate creates and saves the QR code
func (g *Generator) Generate() error {
	g.qr.BackgroundColor = g.config.BackgroundColor
	g.qr.ForegroundColor = g.config.ForegroundColor
	g.qr.DisableBorder = g.config.Border == 0

	// Calculate final size
	size := g.config.Size * g.config.Border

	// Create the image
	img := g.qr.Image(size)

	// Create output file
	f, err := os.Create(g.config.Output)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer f.Close()

	// Encode as PNG
	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

// Preview returns the QR code as ASCII art
func (g *Generator) Preview() string {
	return g.qr.ToSmallString(false)
} 
