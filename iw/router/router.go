package router

import "github.com/gin-gonic/gin"

func GetRouter(r *gin.RouterGroup) {
	UserApiRouter(r.Group("/users"))
	NellowApiRouter(r.Group("/nellow"))
	ShareApiRouter(r)
}
