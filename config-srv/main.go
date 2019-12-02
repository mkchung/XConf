package main

import (
	"github.com/Allenxuxu/XConf/config-srv/conf"
	"github.com/Allenxuxu/XConf/config-srv/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

var config conf.Config

func main() {
	log.Name("XConf")

	service := micro.NewService(
		micro.Name("go.micro.config"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_driver",
				Usage:  "database driver",
				EnvVar: "DATABASE_DRIVER",
				Value:  "mysql",
			},
			cli.StringFlag{
				Name:   "database_url",
				Usage:  "database url",
				EnvVar: "DATABASE_URL",
				Value:  "root:123@(127.0.0.1:3306)/xconf?charset=utf8&parseTime=true&loc=Local",
			}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			config.DB.DriverName = c.String("database_driver")
			config.DB.URL = c.String("database_url")
			log.Infof("database_driver: %s , database_url: %s\n", config.DB.DriverName, config.DB.URL)
		}),
		micro.BeforeStart(func() (err error) {
			if err = dao.Init(&config); err != nil {
				return
			}
			if err = dao.GetDao().Ping(); err != nil {
				return
			}

			return
		}),
	)

	if err := service.Run(); err != nil {
		panic(err)
	}
}
