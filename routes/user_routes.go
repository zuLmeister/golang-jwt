package routes

import (
	"practice-golang/controllers"
	"practice-golang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/api/v1/users")
	users.Use(middleware.AuthRequired())
	{
		users.GET("/", controllers.GetUsers)
	}
}
