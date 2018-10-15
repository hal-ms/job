package cache

import (
	"log"
	"time"

	"github.com/hal-ms/job/repo"
)

func init() {
	go loder()
}

func loder() {
	defer func() {
		if err := recover(); err != nil {
			repo.GetDBConn().Session.Refresh()
			log.Printf("[cache panic]\n%+v\n", err)
			// TODO slackに送る
		}
	}()
	for {
		ReloadAll()
		time.Sleep(3 * time.Second)
	}
}
func ReloadAll() {
	Job.Load()
}
