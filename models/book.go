package models

import (
	"errors"
	"strings"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	ISBN   string `json:"isbn,omitempty"`
	Year   int    `json:"year,omitempty"`
}

var (
	ErrBookNotFound   = errors.New("book not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrBookAlreadyExists = errors.New("book already exists")
)

func (b *Book) Validate() error {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)

	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.Year < 0 || b.Year > 2100 {
		return errors.New("invalid year")
	}
	return nil
}
