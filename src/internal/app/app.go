package app

import (
	"os"
	"syscall"

	"log/slog"
	"os/signal"

	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/adapters/discord"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/adapters/mongo"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/config"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/repositories"
	"github.com/subhadip0x539/bum-bot-main-srv/src/internal/core/services"
	handlers "github.com/subhadip0x539/bum-bot-main-srv/src/internal/handlers/discord"
)

func Run(cfg config.Config) {
	discord, err := discord.NewDiscordClient(cfg.Discord.Token)
	if err != nil {
		slog.Error(err.Error())
	}

	mongo, err := mongo.NewMongoClient(cfg.Mongo.URI)
	if err != nil {
		slog.Error(err.Error())
	}

	if err := mongo.Connect(); err != nil {
		slog.Error(err.Error())
	}
	defer mongo.Disconnect()

	discordRepo := repositories.NewDiscordRepo(discord.Session)
	mongoRepo := repositories.NewMongoRepo(mongo.Client, cfg.Mongo.Database)

	welcomeService := services.NewWelcomeService(discordRepo, mongoRepo)
	welcomeHandler := handlers.NewWelcomeHandler(welcomeService)

	setupService := services.NewSetupService(mongoRepo)
	setupHandler := handlers.NewSetupHandler(setupService)

	discord.RegisterHandler(welcomeHandler.WelcomeHandlerFunc)
	discord.RegisterHandler(setupHandler.SetupHandlerFunc)

	discord.Start()
	defer discord.Stop()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
