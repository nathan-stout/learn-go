package routes

import (
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, db *database.DB) {
	// Initialize repositories
	albumRepo := repositories.NewAlbumRepository(db)

	// Initialize services with repository dependencies
	albumService := services.NewAlbumService(albumRepo)

	// Initialize handlers with service dependencies
	albumHandler := handlers.NewAlbumHandler(albumService)

	// API version group
	v1 := router.Group("/api/v1")
	{
		// Album routes
		albums := v1.Group("/albums")
		{
			albums.GET("", albumHandler.GetAlbums)
			albums.GET("/:id", albumHandler.GetAlbumByID)
			albums.POST("", albumHandler.AddAlbum)
			albums.DELETE("/:id", albumHandler.RemoveAlbum)
		}

		// You can easily add more resource groups here
		// users := v1.Group("/users")
		// {
		//     userRepo := repositories.NewUserRepository(db)
		//     userService := services.NewUserService(userRepo)
		//     userHandler := handlers.NewUserHandler(userService)
		//     users.GET("", userHandler.GetUsers)
		//     users.POST("", userHandler.CreateUser)
		// }
	}
}
