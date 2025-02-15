package main

import (
	"os"

	"log/slog"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/app"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/config"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/pkg/motd"
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
