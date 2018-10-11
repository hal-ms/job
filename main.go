package main

import (
	"github.com/hal-ms/job/router"
	"github.com/makki0205/config"
)

func main()  {
	router.GetRouter().Run(config.Env("port"))
}