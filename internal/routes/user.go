package routes

import (
	"github.com/banisys/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(route *gin.Engine) {

	route.POST("/signup", handlers.Signup)

}
