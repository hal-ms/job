package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Share(c *gin.Context) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	b, err := json.Marshal(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	proxyRequest, err := http.NewRequest(
		"POST",
		"",
		bytes.NewBuffer(b),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	proxyRequest.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(proxyRequest)
	if err != nil {
		fmt.Println("request")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil ReadAll")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var proxyResponse struct {
		URL string `json:"url"`
	}

	err = json.Unmarshal(body, &proxyResponse)
	if err != nil {
		fmt.Println("json unmarshal error")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, proxyResponse)
}
