package routes

import (
	"github.com/banisys/api-gateway/internal/handlers/user_handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {

	route.POST("/signup", user_handler.Signup)

}
