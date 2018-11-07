package cnt

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/model"
	"github.com/hal-ms/job/repo"
	"github.com/hal-ms/job/socket"
)

func GetToken(c *gin.Context) {
	util.Json(repo.Token.Store(model.Token{}), c)
}

func AliveToken(c *gin.Context) {
	token := repo.Token.FindByID(c.Param("token"))
	if c.Param("token") != "masui" {
		if token == nil {
			util.NotFound(c)
			return
		}
		if token.Done {
			util.NotFound(c)
			return
		}
		if !token.IsOpen {
			fmt.Println("send!")
			socket.SendAll("open", "空いたよ！")
		}
		token.IsOpen = true
		repo.Token.Update(*token)
	}
	util.NoContent(c)
}
