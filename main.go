package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"

	"desent-test/m/handlers"
	"desent-test/m/models"
	"desent-test/m/storage"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize zerolog logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("APP_ENV") == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Create Echo instance
	e := echo.New()

	// Echo middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Custom middleware for zerolog
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			zl.Info().
				Str("method", c.Request().Method).
				Str("uri", c.Request().URL.String()).
				Str("remote_ip", c.RealIP()).
				Msg("Request received")
			return next(c)
		}
	})

	// Health check endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Welcome to Echo API",
			"status":  "healthy",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	// Ping endpoint
	e.GET("/ping", func(c echo.Context) error {
		startTime := time.Now()
		return c.JSON(200, models.PingResponse{
			Status: "ok",
			Pong:   true,
			Time:   time.Now().Unix(),
			Uptime: time.Since(startTime).String(),
		})
	})

	// Echo endpoint
	e.POST("/echo", func(c echo.Context) error {
		var req models.EchoRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, models.ErrorResponse("Invalid request body"))
		}

		return c.JSON(200, models.EchoResponse{
			Received: req.Message,
			Echoed:   req.Message,
			At:       time.Now().Unix(),
		})
	})

	// Initialize storage
	bookStore := storage.NewBookStore()

	// Initialize handlers
	logger := zl.Logger
	bookHandler := handlers.NewBookHandler(bookStore, &logger)

	// Book routes
	books := e.Group("/books")
	{
		books.GET("", bookHandler.List)          // GET /books - List all books
		books.POST("", bookHandler.Create)       // POST /books - Create a new book
		books.GET("/:id", bookHandler.Get)       // GET /books/:id - Get a book by ID
		books.PUT("/:id", bookHandler.Update)    // PUT /books/:id - Update a book
		books.DELETE("/:id", bookHandler.Delete) // DELETE /books/:id - Delete a book
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	zl.Info().Msgf("Starting server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		zl.Fatal().Err(err).Msg("Failed to start server")
	}
}
