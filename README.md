# QR Generator

A powerful command-line tool for generating QR codes with customizable options. Built with Go and Cobra.

![QR Code Example](docs/images/example.png)

## Features

- Generate QR codes from text or file input
- Preview QR codes in terminal as ASCII art
- Customize QR code appearance:
  - Size and border width
  - Foreground and background colors
  - Error correction levels
- Support for multiple output formats
- Easy-to-use command-line interface

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Installing from source

```bash
# Clone the repository
git clone https://github.com/iyuangang/qr-generator.git
cd qr-generator

# Install dependencies
go mod download

# Build the binary
make build
```

The binary will be available in the `bin` directory.

## Usage

### Basic Commands

```bash
# Generate a QR code from text
qr generate "Hello World" -o hello.png

# Generate from file content
qr generate input.txt -o output.png

# Preview QR code in terminal
qr preview "Hello World"

# Show version information
qr version
```

### Advanced Options

```bash
# Customize size and border
qr generate "Hello" -o hello.png -s 15 -b 2

# Change colors
qr generate "Hello" -o hello.png --fg-color blue --bg-color white

# Set error correction level
qr generate "Hello" -o hello.png -e high
```

### Command Line Options

```bash
Usage:
  qr [command]

Available Commands:
  generate    Generate QR code from text or file
  preview     Preview QR code in terminal
  version     Print version information
  help        Help about any command

Flags for generate:
  -o, --output string       Output image file path (default "qr_code.png")
  -s, --size int           QR code size (1-40) (default 10)
  -b, --border int         Border size (default 4)
      --fg-color string    Foreground color (black/white/blue/red/green) (default "black")
      --bg-color string    Background color (black/white/blue/red/green) (default "white")
  -e, --error-level string Error correction level (low/medium/high/highest) (default "medium")
```

## Project Structure

```
.
├── cmd/                # Command line interface
│   ├── generate.go    # Generate command
│   ├── preview.go     # Preview command
│   ├── root.go        # Root command
│   └── version.go     # Version command
├── pkg/
│   └── qr/           # QR code generation package
│       ├── config.go   # Configuration
│       └── generator.go # QR code generator
├── main.go           # Entry point
└── Makefile         # Build automation
```

## Development

### Building from Source

```bash
# Build the binary
make build

# Run tests
make test

# Clean build artifacts
make clean
```

### Running Tests

```bash
go test -v ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Version History

- v0.1.0
  - Initial release
  - Basic QR code generation
  - Preview functionality
  - Color customization
  - Error correction levels

## Acknowledgments

- [go-qrcode](https://github.com/skip2/go-qrcode) - QR code generation library
- [cobra](https://github.com/spf13/cobra) - CLI framework

## Author

Tim Yuan - [@iyuangang](https://github.com/iyuangang)

## Support

If you have any questions or need help, please:
1. Check the [documentation](docs/)
2. Open an issue
3. Contact the maintainers
