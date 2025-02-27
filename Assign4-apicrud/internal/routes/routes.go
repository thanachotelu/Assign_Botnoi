package routes

import (
	"crud/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userHandlers *handler.UserHandlers) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandlers.AddUser)
			users.GET("", userHandlers.GetAllUsers)
			users.GET("/:id", userHandlers.GetUserById)
			users.PUT("/:id", userHandlers.UpdateUser)
			users.DELETE("/:id", userHandlers.DeleteUser)
		}

	}
}
