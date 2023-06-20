package main

import (
	"fmt"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"

	"bot/commands/counter"
	"bot/commands/pixiv"
)

func Logger(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		fmt.Printf("%s %s: %s\n", c.Message().Time(), c.Message().Sender.Username, c.Text())
		return next(c)
	}
}

func start(token string) error {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return fmt.Errorf("Error when create new bot, %w", err)
	}

	b.Use(Logger)
	b.Use(middleware.AutoRespond())

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi, this is SimbaFs' Telegram bot")
	})

	b.Handle("/ping", func (c tele.Context) error {
		return c.Reply("pong!")
	})

	// commands
	counter.Command(b)
	pixiv.Command(b)

	fmt.Println("Bot is running!")
	b.Start()
	return nil
}

func main() {
	token := os.Getenv("TELEBOT_TOKEN")
	if token == "" {
		fmt.Printf("Set env var TELEBOT_TOKEN to provide token\n")
		os.Exit(1)
	}
	if err := start(token); err != nil {
		fmt.Printf("Oops, there's an error: %v\n", err)
		os.Exit(1)
	}
}
