package main

import (
	"github.com/gomyred/api"
	"github.com/gomyred/config"
	"github.com/gomyred/src/article/repository"
	"github.com/gomyred/utils"
)

var (
	co config.Config
	lg utils.Logger
)

func main() {

	co = config.GetConfig()

	_, err := repository.NewMySQL(co.MySQL)
	if err != nil {
		lg.Info = "MySQL"
		lg.Error = err
		lg.NewLogError()
	}

	_, err = repository.NewRedis(co)

	if err != nil {
		lg.Info = "Redis"
		lg.Error = err
		lg.NewLogError()
	}

	api.RunningAPI()

}
