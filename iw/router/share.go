package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/controller"
)

func ShareApiRouter(r *gin.RouterGroup) {
	r.POST("/share", controller.Share)
}
