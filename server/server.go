package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"myGo/adapter/log"
	"myGo/adapter/mysql"
	"myGo/adapter/redis"
	"myGo/config"
)

type Server struct {
	*cli.App
	config *config.Config
}

func initWithConfig(filePath string) (*config.Config, error) {
	conf, err := config.Load(filePath)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func initLog(conf *config.Config) error {
	return log.NewLoggerWithOptions(conf.Log)
}

func initMysql(conf *config.Config) error {
	db, err := mysql.InitializeMainDb(conf.Mysql.Master)
	if err != nil {
		return err
	}
	mysql.InitEntityDao(db)
	return nil
}

func NewServer(ctx context.Context) *Server {
	s := &Server{
		App: cli.NewApp(),
	}
	s.Flags = []cli.Flag{cli.StringFlag{Name: "c", Usage: "Configuration file"}}
	s.Action = func(c *cli.Context) error {
		if c.GlobalString("c") == "" {
			return errors.New("usage: my_go -c configfilepath")
		}

		fmt.Println("start read config: ", c.GlobalString("c"))
		conf, err := initWithConfig(c.GlobalString("c"))
		if err != nil {
			return errors.Wrap(err, "fail to init conf")
		}
		fmt.Println("init config success. conf:", conf)
		s.config = conf

		// log
		err = initLog(conf)
		if err != nil {
			return errors.Wrap(err, "fail to init log")
		}

		// mysql
		err = initMysql(conf)
		if err != nil {
			return errors.Wrap(err, "fail to init mysql")
		}

		// redis init
		err = redis.Initialize(conf.Redis.Default)
		if err != nil {
			return err
		}

		r := gin.Default()
		routes(r)
		if err = r.Run(s.config.HTTPPort); err != nil {
			return errors.Wrap(err, "fail to run")
		}

		return nil
	}
	return s
}
