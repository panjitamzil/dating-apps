package route

import (
	"dating-apps/config"
	"dating-apps/domain/users/handler"

	"github.com/gin-gonic/gin"
)

func Routes(h *config.Handler) *gin.Engine {
	r := gin.Default()
	r.POST("/register", h.Users.Register)
	r.POST("/login", h.Users.Login)

	r.Use(handler.Authentication())
	r.GET("/explore", h.Users.GetProfiles)
	r.POST("/swipe/:id/:action", h.Users.Swipe)
	r.POST("/purchase", h.Users.Purchase)

	return r
}
