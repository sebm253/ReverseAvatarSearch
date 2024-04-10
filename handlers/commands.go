package handlers

import "github.com/disgoorg/disgo/handler"

func NewHandler() *Handler {
	h := &Handler{
		Router: handler.New(),
	}
	h.Group(func(r handler.Router) {
		r.Route("/reverse", func(r handler.Router) {
			r.SlashCommand("/user", h.HandleReverseUserSlash)
			r.SlashCommand("/link", h.HandleReverseLink)
		})
		r.UserCommand("/Reverse avatar", h.HandleReverseUserContext)
	})
	return h
}

type Handler struct {
	handler.Router
}
