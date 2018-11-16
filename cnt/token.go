package cnt

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/cnt/util"
	"github.com/hal-ms/job/model"
	"github.com/hal-ms/job/repo"
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
			http.Get("https://socket.patra.store/emit/qr/open")
		}
		token.IsOpen = true
		repo.Token.Update(*token)
	}
	util.NoContent(c)
}
