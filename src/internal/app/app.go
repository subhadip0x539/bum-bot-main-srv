package app

import (
	"os"
	"syscall"

	"log/slog"
	"os/signal"

	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/adapters"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/config"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/repositories"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/core/services"
	"github.com/subhadip0x539/bum-bot-event-hdl/src/internal/handlers"
)

func Run(cfg config.Config) {
	discord, err := adapters.NewDiscordClient(cfg.Discord.Token)
	if err != nil {
		slog.Error(err.Error())
	}

	mongo, err := adapters.NewMongoClient(cfg.Mongo.URI)
	if err != nil {
		slog.Error(err.Error())
	}

	if err := mongo.Connect(); err != nil {
		slog.Error(err.Error())
	}
	defer mongo.Disconnect()

	discordRepo := repositories.NewDiscordRepo(discord.Session)
	mongoRepo := repositories.NewMongoRepo(mongo.Client, cfg.Mongo.Database)

	greetingsService := services.NewWelcomeService(discordRepo, mongoRepo)
	greetingsHandler := handlers.NewGreetingsHandler(greetingsService)

	setupService := services.NewSetupService(mongoRepo)
	setupHandler := handlers.NewSetupHandler(setupService)

	discord.RegisterHandler(greetingsHandler.MemberAddHandlerFunc)
	discord.RegisterHandler(greetingsHandler.MemberRemoveHandlerFunc)
	discord.RegisterHandler(setupHandler.SetupHandlerFunc)

	discord.Start()
	defer discord.Stop()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
