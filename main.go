package main

import (
	"context"
	"flag"
	"github.com/major28/monitoring/application"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var configPath string
	var env string
	flag.StringVar(&configPath, "config", "./", "A path to config file")
	flag.StringVar(&env, "env", "dev", "Environment")
	flag.Parse()

	app := application.NewApp()
	app.Initialize(configPath, env)
	defer app.Resources.GetDbManager().Close()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(quit)
		<-quit
		cancel()
	}()

	app.Run(ctx)
}
