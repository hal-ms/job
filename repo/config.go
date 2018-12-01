package repo

import "github.com/hal-ms/job/model"

var Config configService

type configService struct {
}

var config = model.Config{
	Url: model.Url{
		Main:   "http://192.168.0.9",
		LED:    "http://192.168.0.2",
		Moving: "http://192.168.0.5",
		Game:   "http://192.168.0.10",
	},
}

func (c configService) Get() model.Config {
	return config
}

func (c configService) Set(cnf model.Config) {
	config = cnf
}
