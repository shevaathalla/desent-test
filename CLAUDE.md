# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web API project using the Echo framework (v4.15.1) with Go 1.24.7. The project includes:
- **Echo Framework**: High-performance web framework
- **Zerolog**: Zero-allocation logging library
- **Godotenv**: Environment variable management via .env files

Module name: `desent-test/m`

## Common Commands

```bash
# Run the application
go run main.go

# Build the project
go build -o app

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a specific test
go test -v -run TestFunctionName ./path/to/package

# Format code
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Download dependencies
go mod download

# Tidy up dependencies
go mod tidy

# Add a new dependency
go get <package>
```

## Architecture

### Entry Point (main.go)
The application bootstraps in this order:
1. Load environment variables from `.env` file using `godotenv`
2. Configure `zerolog` logger based on `APP_ENV` (development/production)
3. Create Echo instance with middleware stack:
   - `Logger` - HTTP request logging
   - `Recover` - Panic recovery
   - `CORS` - Cross-origin resource sharing
   - `Gzip` - Response compression
   - Custom zerolog middleware for structured logging
4. Register routes and start server on `PORT` (default: 8080)

### Environment Variables
| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Application environment (development/production) | - |
| `PORT` | Server port | `8080` |

### Project Structure

```
.
├── main.go           # Application entry point
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
├── .env              # Environment variables (git-ignored)
├── .env.example      # Environment variables template
├── .gitignore        # Git ignore patterns
└── CLAUDE.md         # This file
```
