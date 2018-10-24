package cnt

import (
	"strings"

	"github.com/hal-ms/job/cache"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/req"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/repo"
)

func OpenJob(c *gin.Context) {
	token := repo.Token.FindByID(c.Param("token"))
	if token == nil {
		util.NotFound(c)
		return
	}
	if token.Done == true {
		util.NotFound(c)
		return
	}
	token.IsOpen = true
	repo.Token.Update(*token)
	util.NoContent(c)
}

func CreateJob(c *gin.Context) {
	req := req.Job(c)
	if req == nil {
		return
	}

	token := repo.Token.FindByID(c.Param("token"))
	if token == nil {
		util.NotFound(c)
		return
	}
	if token.Done {
		util.NotFound(c)
		return
	}

	// tokenの無効化
	token.IsOpen = true
	token.Done = true
	repo.Token.Update(*token)
	// jobの登録
	util.Json(repo.Job.Store(*req), c)

}

func GetJobs(c *gin.Context) {
	util.Json(cache.Job.All().FindByIds(strings.Split(c.Param("jobs"), ",")), c)
}
