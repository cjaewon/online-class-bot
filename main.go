package main

import (
	"os"
	"os/signal"
	"syscall"

	"onlineclassbot/internal/config"
	"onlineclassbot/internal/logger"
	"onlineclassbot/internal/schedule"

	"github.com/bwmarrin/discordgo"
)

func main() {
	defer logger.Log().Sync()

	cfg := config.ReadInConfig()
	scheduler := schedule.New(cfg)

	discord, err := discordgo.New("Bot " + cfg.Bot.Token)

	if err != nil {
		logger.Log().Fatalw("Failed to create bot client", "err", err)
	}

	discord.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		scheduler.Register(s)
		scheduler.Start()
	})

	if err := discord.Open(); err != nil {
		logger.Log().Fatal("Failed to Open connecting", "err", err)
	}

	logger.Log().Infof("Start running a %q bot", discord.State.User.Username)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	discord.Close()
}
