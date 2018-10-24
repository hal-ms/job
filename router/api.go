package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/token", cnt.GetToken)
	api.GET("/alive_token/:token", cnt.AliveToken)

	api.POST("/job/:token", cnt.CreateJob)
	api.GET("/jobs/:jobs", cnt.GetJobs)

	api.GET("/job_categorys", cnt.GetJobCategorys)
}
