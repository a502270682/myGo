package server

import (
	"context"
	"myGo/adapter/log"
	"myGo/config"
)

type Server struct {
}

func initWithConfig(ctx context.Context) error {
	conf, err := config.Load("../config/template.toml")
	if err != nil {
		return err
	}

	return nil
}

func NewServer(ctx context.Context) *Server {
	return nil
}
