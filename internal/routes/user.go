package routes

import (
	"github.com/banisys/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	// route.POST("/user-service/*path", handlers.User)
	route.POST("/signup", handlers.User)

}
