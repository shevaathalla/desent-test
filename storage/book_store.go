package storage

import (
	"sync"

	"github.com/google/uuid"

	"desent-test/m/models"
)

type BookStore struct {
	mu    sync.RWMutex
	books map[string]*models.Book
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[string]*models.Book),
	}
}

func (s *BookStore) List() []*models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]*models.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

func (s *BookStore) Get(id string) (*models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, ok := s.books[id]
	if !ok {
		return nil, models.ErrBookNotFound
	}
	return book, nil
}

func (s *BookStore) Create(book *models.Book) (*models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate UUID
	book.ID = uuid.New().String()

	s.books[book.ID] = book
	return book, nil
}

func (s *BookStore) Update(id string, updates *models.Book) (*models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.books[id]
	if !ok {
		return nil, models.ErrBookNotFound
	}

	// Update fields
	if updates.Title != "" {
		existing.Title = updates.Title
	}
	if updates.Author != "" {
		existing.Author = updates.Author
	}
	if updates.ISBN != "" {
		existing.ISBN = updates.ISBN
	}
	if updates.Year != 0 {
		existing.Year = updates.Year
	}

	return existing, nil
}

func (s *BookStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		return models.ErrBookNotFound
	}

	delete(s.books, id)
	return nil
}
