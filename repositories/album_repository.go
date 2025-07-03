package repositories

import (
	"database/sql"
	"fmt"
	"server/database"
	"server/services"
)

// AlbumRepository handles album database operations
type AlbumRepository struct {
	db *database.DB
}

// NewAlbumRepository creates a new album repository
func NewAlbumRepository(db *database.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

// GetAll retrieves all albums from the database
func (r *AlbumRepository) GetAll() ([]services.Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM albums 
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query albums: %v", err)
	}
	defer rows.Close()

	var albums []services.Album
	for rows.Next() {
		var album services.Album

		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Artist,
			&album.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan album: %v", err)
		}

		albums = append(albums, album)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over albums: %v", err)
	}

	return albums, nil
}

// GetByID retrieves an album by its ID
func (r *AlbumRepository) GetByID(id string) (*services.Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM albums 
		WHERE id = $1
	`

	var album services.Album

	err := r.db.QueryRow(query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, services.ErrAlbumNotFound
		}
		return nil, fmt.Errorf("failed to get album: %v", err)
	}

	return &album, nil
}

// Create inserts a new album into the database
func (r *AlbumRepository) Create(req services.AlbumRequest) (*services.Album, error) {
	query := `
		INSERT INTO albums (title, artist, price)
		VALUES ($1, $2, $3)
		RETURNING id, title, artist, price
	`

	var album services.Album

	err := r.db.QueryRow(query, req.Title, req.Artist, req.Price).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err != nil {
		// Check for unique constraint violation (duplicate album)
		if isUniqueViolation(err) {
			return nil, services.ErrDuplicateAlbum
		}
		return nil, fmt.Errorf("failed to create album: %v", err)
	}

	return &album, nil
}

// Delete removes an album from the database
func (r *AlbumRepository) Delete(id string) error {
	query := `DELETE FROM albums WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete album: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return services.ErrAlbumNotFound
	}

	return nil
}

// ExistsByTitleAndArtist checks if an album with the given title and artist exists
func (r *AlbumRepository) ExistsByTitleAndArtist(title, artist string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM albums WHERE title = $1 AND artist = $2)`

	var exists bool
	err := r.db.QueryRow(query, title, artist).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check album existence: %v", err)
	}

	return exists, nil
}

// isUniqueViolation checks if the error is a unique constraint violation
func isUniqueViolation(err error) bool {
	// PostgreSQL unique violation error code is 23505
	return err != nil &&
		err.Error() == `pq: duplicate key value violates unique constraint "albums_title_artist_key"`
}
