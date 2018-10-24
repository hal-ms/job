package req

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/model"
	"github.com/hal-ms/job/repo"
)

func Job(c *gin.Context) *model.Job {
	var job model.Job
	err := c.BindJSON(&job)
	if err != nil {
		util.BadRequest(err.Error(), c)
		return nil
	}
	jC := repo.JobCategory.All().FindByName(job.Name)
	if jC == nil {
		util.BadRequest("jobCategory err", c)
		return nil
	}
	job.DisplayName = jC.DisplayName
	job.Done = false
	job.Created = time.Now()
	return &job
}
