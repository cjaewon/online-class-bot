package main

import (
	"onlineclassbot/internal/config"
	"onlineclassbot/internal/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg := config.ReadInConfig()
	discord, err := discordgo.New("Bot " + cfg.Bot.Token)

	if err != nil {
		logger.Log().Fatalw("Failed to create bot client", "err", err)
	}
	if err := discord.Open(); err != nil {
		logger.Log().Fatal("Failed to Open connecting", "err", err)
	}

	logger.Log().Infof("Start running a %q bot", discord.State.User.Username)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	discord.Close()
}
