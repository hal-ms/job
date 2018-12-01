package repo

import "github.com/hal-ms/job/model"

var Config configService

type configService struct {
}

var configRep = model.Config{
	Url: model.Url{
		Main:   "http://192.168.0.9:8000",
		LED:    "http://192.168.0.2:8000",
		Moving: "http://192.168.0.5:5000",
		Game:   "http://192.168.0.10:8080",
	},
}

func (c configService) Get() model.Config {
	return configRep
}

func (c configService) Set(cnf model.Config) {
	configRep = cnf
}
