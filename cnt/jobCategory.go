package cnt

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/repo"
)

func GetJobCategorys(c *gin.Context) {
	util.Json(repo.JobCategory.All(), c)
}
