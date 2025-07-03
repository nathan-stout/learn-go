package handlers

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
)

// AlbumHandler contains all album-related handlers
type AlbumHandler struct {
	albumService *services.AlbumService
}

// NewAlbumHandler creates a new album handler
func NewAlbumHandler(albumService *services.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		albumService: albumService,
	}
}

// GetAlbums responds with the list of all albums as JSON.
func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := h.albumService.GetAllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID responds with the album whose ID value matches the id parameter.
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.albumService.GetAlbumByID(id)
	if err != nil {
		if err == services.ErrAlbumNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// AddAlbum adds an album from JSON received in the request body.
func (h *AlbumHandler) AddAlbum(c *gin.Context) {
	var req services.AlbumRequest

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request format"})
		return
	}

	album, err := h.albumService.CreateAlbum(req)
	if err != nil {
		// Map business errors to HTTP status codes
		switch err {
		case services.ErrInvalidPrice:
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case services.ErrEmptyTitle:
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case services.ErrEmptyArtist:
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case services.ErrDuplicateAlbum:
			c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		default:
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}

// RemoveAlbum removes an album by ID.
func (h *AlbumHandler) RemoveAlbum(c *gin.Context) {
	id := c.Param("id")

	err := h.albumService.DeleteAlbum(id)
	if err != nil {
		if err == services.ErrAlbumNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	// Return the updated list of albums
	albums, err := h.albumService.GetAllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}
