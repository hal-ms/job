package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/controller"
)

func NellowApiRouter(nellow *gin.RouterGroup) {
	nellow.GET("/sleep/:id", controller.NellowController.Sleep)
	nellow.GET("/wakeup/:id", controller.NellowController.Wakeup)
}
