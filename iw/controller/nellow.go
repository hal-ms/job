package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/job/iw/service"
	"net/http"
	"time"
	"github.com/hal-ms/job/iw/model"
	"encoding/json"
)

var NellowController nellowController

type nellowController struct{}

func (n *nellowController) Sleep(c *gin.Context) {
	now := time.Now()
	id := c.Param("id")
	user := service.User.FindByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	err := service.Nellow.Create(user.ID, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "寝ることができませんでした。",
		})
		return
	}

	var res model.NellowSleepRes
	res.PID = user.PID
	res.PRate = user.PRate
	res.ID = user.ID
	res.Name = user.Name
	res.Icon = user.Icon
	res.IsNellow = user.IsNellow
	res.Amount = user.Amount
	res.BedTime = now.UTC().Format("2006-01-02T15:04:05Z07:00")

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	url := "https://socket.patra.store/emit/nellow_sleep?data=" + string(b)

	resp, _ := http.Get(url)
	fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{"msg": "おやすみなさい！"})
}

func (n *nellowController) Wakeup(c *gin.Context) {
	id := c.Param("id")
	user := service.User.FindByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	ns := service.Nellow.FindByUserID(user.ID)
	if ns == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "指定されたIDの人は寝てないです。"})
		return
	}

	d := time.Since(ns.SleepTime)
	sec := d.Seconds()

	err := service.Nellow.Delete(*ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "サービス終了。",
		})
		return
	}

	fmt.Println(sec)
	user.Amount += int(sec) * (user.PRate / 60)
	err = service.User.Update(*user)
	if err != nil {
		panic(err)
	}

	var res model.NellowWakeupRes
	res.User = *user
	res.WakeupTime = time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	url := "https://socket.patra.store/emit/nellow_wakeup?data=" + string(b)

	resp, _ := http.Get(url)
	fmt.Println(resp)

	c.JSON(http.StatusOK, user)
}
