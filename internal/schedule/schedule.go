package schedule

import (
	"context"
	"crypto"
	"encoding/hex"
	"fmt"
	"time"

	"onlineclassbot/internal/config"
	"onlineclassbot/internal/logger"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

// A Schedule represents calendar of online class
type Schedule struct {
	config *config.Config
	cron   *cron.Cron
}

// New creates a new Schedule
func New(cfg *config.Config) *Schedule {
	s := Schedule{
		config: cfg,
		cron:   cron.New(),
	}

	return &s
}

// Register registers all the cron schedules
func (s *Schedule) Register(session *discordgo.Session) {
	for k, v := range s.config.Schedules {
		s.cron.AddFunc(k, func(schedule config.ScheduleConfig) func() {
			return func() {
				hashed := hex.EncodeToString(crypto.MD5.New().Sum([]byte(schedule.URL)))

				embed := discordgo.MessageEmbed{
					Title:       fmt.Sprintf("%s님의 수업", schedule.Teacher),
					Description: fmt.Sprintf("[%s](%s)", schedule.URL, schedule.URL),

					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=64&d=retro", hashed),
					},

					Timestamp: time.Now().Format(time.RFC3339),
					Color:     0xc5e1a5,
				}

				session.UpdateListeningStatus(schedule.Teacher)

				if _, err := session.ChannelMessageSendEmbed(s.config.Bot.NotifyID, &embed); err != nil {
					logger.Log().Errorw("Failed to send a notify schedules message", "err", err)
				}
			}
		}(v))
	}
}

// Start starts all the cron schedules
func (s *Schedule) Start() {
	s.cron.Start()
}

// Stop stops all the cron schedules
func (s *Schedule) Stop() context.Context {
	return s.cron.Stop()
}
