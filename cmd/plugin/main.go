package main

import (
	"context"
	"math/rand"
	"myGo/adapter/log"
	"myGo/cron"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	cron.StartCron()
	signals := make(chan os.Signal, 1)
	defer close(signals)
	signal.Notify(signals, os.Kill, os.Interrupt, syscall.SIGBUS, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	<-signals
	cron.StopCron()
	log.Info(ctx, "service start to shut down")
}
