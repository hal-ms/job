package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/controller"
)

func UserApiRouter(user *gin.RouterGroup) {
	user.POST("/", controller.UserController.Create)
	user.GET("/:id", controller.UserController.Get)
	user.PATCH("/:id", controller.UserController.Update)
	user.PATCH("/:id/icon", controller.UserController.UpdateIcon)
}
