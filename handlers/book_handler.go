package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"desent-test/m/models"
	"desent-test/m/storage"
)

type BookHandler struct {
	store *storage.BookStore
	log   *zerolog.Logger
}

func NewBookHandler(store *storage.BookStore, log *zerolog.Logger) *BookHandler {
	return &BookHandler{
		store: store,
		log:   log,
	}
}

func (h *BookHandler) List(c echo.Context) error {
	books := h.store.List()

	h.log.Debug().Int("count", len(books)).Msg("Listing books")

	response := models.SuccessResponse(books)
	response.Meta.Total = len(books)

	return c.JSON(http.StatusOK, response)
}

func (h *BookHandler) Get(c echo.Context) error {
	id := c.Param("id")

	book, err := h.store.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse("book not found"))
	}

	return c.JSON(http.StatusOK, models.SuccessResponse(book))
}

func (h *BookHandler) Create(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse("invalid request body"))
	}

	if err := book.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
	}

	created, err := h.store.Create(&book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse("failed to create book"))
	}

	h.log.Info().Str("id", created.ID).Str("title", created.Title).Msg("Book created")
	return c.JSON(http.StatusCreated, models.SuccessResponseWithMessage(created, "book created successfully"))
}

func (h *BookHandler) Update(c echo.Context) error {
	id := c.Param("id")

	var updates models.Book
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse("invalid request body"))
	}

	// Validate updates
	if updates.Title != "" || updates.Author != "" {
		temp := &models.Book{Title: updates.Title, Author: updates.Author}
		if err := temp.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		}
	}

	book, err := h.store.Update(id, &updates)
	if err != nil {
		if err == models.ErrBookNotFound {
			return c.JSON(http.StatusNotFound, models.ErrorResponse("book not found"))
		}
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse("failed to update book"))
	}

	h.log.Info().Str("id", id).Msg("Book updated")
	return c.JSON(http.StatusOK, models.SuccessResponseWithMessage(book, "book updated successfully"))
}

func (h *BookHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	err := h.store.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrorResponse("book not found"))
	}

	h.log.Info().Str("id", id).Msg("Book deleted")
	return c.JSON(http.StatusOK, models.SuccessResponseWithMessage(nil, "book deleted successfully"))
}
