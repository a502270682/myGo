package main

import (
	"context"
	"math/rand"
	"myGo/adapter/log"
	"myGo/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	if err := server.NewServer(ctx).Run(os.Args); err != nil {
		log.Errorf(ctx, "server run with error: %v", err)
	}
	signals := make(chan os.Signal, 1)
	defer close(signals)
	signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGBUS, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	<-signals
	log.Info(ctx, "service start to shut down")
}
