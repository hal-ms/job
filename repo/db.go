package repo

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/makki0205/config"
	"gopkg.in/mgo.v2"
)

var db = NewDBConn()

func NewDBConn() *mgo.Database {
	fmt.Println(config.GoEnv)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    strings.Split(config.Env("dbhost"), ","),
		Timeout:  60 * time.Second,
		Username: config.Env("dbUser"),
		Password: config.Env("dbPassword"),
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Source: "admin",
	})
	if err != nil {
		panic(err)
	}
	return session.DB(config.Env("database"))
}

func GetDBConn() *mgo.Database {
	return db
}
