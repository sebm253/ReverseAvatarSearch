package main

import (
	"context"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"os"
	"os/signal"
	"syscall"
)

const (
	googleLensSearch = "https://lens.google.com/uploadbyurl?url="
)

func main() {
	log.SetLevel(log.LevelInfo)
	log.Info("starting the bot...")
	log.Info("disgo version: ", disgo.Version)

	client, err := disgo.New(os.Getenv("REVERSE_AVATAR_SEARCH_TOKEN"),
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentsNone)),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagsNone)),
		bot.WithEventListeners(&events.ListenerAdapter{
			OnApplicationCommandInteraction: onCommand,
		}))
	if err != nil {
		log.Fatal("error while building disgo instance: ", err)
	}

	defer client.Close(context.TODO())

	if err := client.OpenGateway(context.TODO()); err != nil {
		log.Fatal("error while connecting to the gateway: ", err)
	}

	log.Info("reverse avatar search is now running.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-s
}

func onCommand(event *events.ApplicationCommandInteractionCreate) {
	data := event.Data
	var link string
	switch data := data.(type) {
	case discord.SlashCommandInteractionData:
		subcommand := *data.SubCommandName
		if subcommand == "user" {
			user := data.User("user")
			link = user.EffectiveAvatarURL(discord.WithSize(4096))
		} else if subcommand == "link" {
			link = data.String("link")
		}
	case discord.UserCommandInteractionData:
		user := data.TargetUser()
		link = user.EffectiveAvatarURL(discord.WithSize(4096))
	}
	_ = event.CreateMessage(discord.NewMessageCreateBuilder().
		AddActionRow(discord.NewLinkButton("Open reverse image search", googleLensSearch+link)).
		SetFlags(discord.MessageFlagEphemeral).
		Build())
}
