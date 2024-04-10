package handlers

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

const (
	googleLensSearch = "https://lens.google.com/uploadbyurl?url="
)

func (h *Handler) HandleReverseUserSlash(data discord.SlashCommandInteractionData, event *handler.CommandEvent) error {
	return handleReverseUser(data.User("user"), event)
}

func (h *Handler) HandleReverseUserContext(data discord.UserCommandInteractionData, event *handler.CommandEvent) error {
	return handleReverseUser(data.TargetUser(), event)
}

func handleReverseUser(user discord.User, event *handler.CommandEvent) error {
	return handleReverse(user.EffectiveAvatarURL(discord.WithSize(4096)), event)
}

func (h *Handler) HandleReverseLink(data discord.SlashCommandInteractionData, event *handler.CommandEvent) error {
	return handleReverse(data.String("link"), event)
}

func handleReverse(link string, event *handler.CommandEvent) error {
	return event.CreateMessage(discord.NewMessageCreateBuilder().
		AddActionRow(discord.NewLinkButton("Open reverse image search", googleLensSearch+link)).
		SetEphemeral(true).
		Build())
}
