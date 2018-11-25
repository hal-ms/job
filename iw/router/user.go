package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/controller"
)

func UserApiRouter(user *gin.RouterGroup) {
	user.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ようこそ Userさま")
	})
	user.GET("/", controller.UserController.Get)
	user.POST("/", controller.UserController.Create)
}
