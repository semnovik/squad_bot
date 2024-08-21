package main

import (
	"github.com/semnovik/squad_bot/internal/adapter/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.LoadConfig()

	pref := tele.Settings{
		Token:  cfg.ApiToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
