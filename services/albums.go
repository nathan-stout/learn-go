package services

import (
	"errors"
	"slices"
	"strconv"
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

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// AlbumService handles album business logic
type AlbumService struct{}

// NewAlbumService creates a new album service
func NewAlbumService() *AlbumService {
	return &AlbumService{}
}

// GetAllAlbums returns all albums
func (s *AlbumService) GetAllAlbums() []Album {
	return albums
}

// GetAlbumByID finds an album by ID
func (s *AlbumService) GetAlbumByID(id string) (*Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, ErrAlbumNotFound
}

// CreateAlbum creates a new album with business validation
func (s *AlbumService) CreateAlbum(req AlbumRequest) (*Album, error) {
	// Business validation
	if err := s.validateAlbumRequest(req); err != nil {
		return nil, err
	}

	// Business rule: no duplicate albums
	if s.isDuplicateAlbum(req.Title, req.Artist) {
		return nil, ErrDuplicateAlbum
	}

	// Generate ID (business logic for ID generation)
	newID := s.generateNextID()

	// Create album
	newAlbum := Album{
		ID:     newID,
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	// Add to collection
	albums = append(albums, newAlbum)

	return &newAlbum, nil
}

// DeleteAlbum removes an album by ID
func (s *AlbumService) DeleteAlbum(id string) error {
	for i, album := range albums {
		if album.ID == id {
			albums = slices.Delete(albums, i, i+1)
			return nil
		}
	}
	return ErrAlbumNotFound
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

// isDuplicateAlbum checks if an album with the same title and artist exists
func (s *AlbumService) isDuplicateAlbum(title, artist string) bool {
	for _, album := range albums {
		if album.Title == title && album.Artist == artist {
			return true
		}
	}
	return false
}

// generateNextID generates the next available ID
func (s *AlbumService) generateNextID() string {
	return strconv.Itoa(len(albums) + 1)
}
