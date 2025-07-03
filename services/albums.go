package services

import (
	"errors"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// AlbumRequest represents the request body for creating an album (no ID field)
type AlbumRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Business logic errors
var (
	ErrAlbumNotFound  = errors.New("album not found")
	ErrInvalidPrice   = errors.New("price must be greater than 0")
	ErrEmptyTitle     = errors.New("title cannot be empty")
	ErrEmptyArtist    = errors.New("artist cannot be empty")
	ErrDuplicateAlbum = errors.New("album with this title and artist already exists")
)

// AlbumRepository interface defines the methods that our repository must implement
type AlbumRepository interface {
	GetAll() ([]Album, error)
	GetByID(id string) (*Album, error)
	Create(req AlbumRequest) (*Album, error)
	Delete(id string) error
	ExistsByTitleAndArtist(title, artist string) (bool, error)
}

// AlbumService handles album business logic
type AlbumService struct {
	repo AlbumRepository
}

// NewAlbumService creates a new album service
func NewAlbumService(repo AlbumRepository) *AlbumService {
	return &AlbumService{
		repo: repo,
	}
}

// GetAllAlbums returns all albums
func (s *AlbumService) GetAllAlbums() ([]Album, error) {
	return s.repo.GetAll()
}

// GetAlbumByID finds an album by ID
func (s *AlbumService) GetAlbumByID(id string) (*Album, error) {
	return s.repo.GetByID(id)
}

// CreateAlbum creates a new album with business validation
func (s *AlbumService) CreateAlbum(req AlbumRequest) (*Album, error) {
	// Business validation
	if err := s.validateAlbumRequest(req); err != nil {
		return nil, err
	}

	// Business rule: no duplicate albums
	exists, err := s.repo.ExistsByTitleAndArtist(req.Title, req.Artist)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateAlbum
	}

	// Create album through repository
	return s.repo.Create(req)
}

// DeleteAlbum removes an album by ID
func (s *AlbumService) DeleteAlbum(id string) error {
	return s.repo.Delete(id)
}

// validateAlbumRequest validates business rules for album creation
func (s *AlbumService) validateAlbumRequest(req AlbumRequest) error {
	if req.Title == "" {
		return ErrEmptyTitle
	}

	if req.Artist == "" {
		return ErrEmptyArtist
	}

	if req.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
