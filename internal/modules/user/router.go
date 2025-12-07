package user

import (
	db "github.com/JangVincent/Gin_PlayGround/database/generated"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, queries *db.Queries) {
	userHandler := NewUserHandler(queries)

	users := r.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUser)
	users.GET("", userHandler.ListUsers)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
}
