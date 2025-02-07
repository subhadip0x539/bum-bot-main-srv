package main

import (
	"log/slog"
	"os"

	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/app"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/config"
	"github.com/subhadip0x539/bum-bot-main-srv/src/pkg/motd"
)

func init() {
	motd.Info()
}

func main() {
	config := config.NewConfig()
	if err := config.Register(".env", "env", os.Getenv("GIN_MODE")); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	cfg := config.Get()

	app.Run(cfg)
}
