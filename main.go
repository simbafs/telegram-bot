package main

import (
	"fmt"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var (
	selector = &tele.ReplyMarkup{}

	btnAdd   = selector.Data("+", "add")
	btnZero  = selector.Data("0", "zero")
	btnMinus = selector.Data("-", "minus")
	// btnSet    = selector.Data("set", "set")
	btnDelete = selector.Data("x", "delete")
)

func Logger(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		fmt.Printf("%s %s: %s\n", c.Message().Time(), c.Message().Sender.Username, c.Text())
		return next(c)
	}
}

var counter = make(map[int]int)

func start(token string) error {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	b.Use(Logger)
	b.Use(middleware.AutoRespond())

	selector.Inline(
		selector.Row(btnAdd, btnZero, btnMinus, btnDelete),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Use /counter to create a counter")
	})

	b.Handle("/counter", func(c tele.Context) error {
		counter[c.Message().ID] = 0
		return c.Send(fmt.Sprint(0), selector)
	})

	b.Handle(&btnAdd, func(c tele.Context) error {
		counter[c.Message().ID]++
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	b.Handle(&btnZero, func(c tele.Context) error {
		counter[c.Message().ID] = 0
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	b.Handle(&btnMinus, func(c tele.Context) error {
		counter[c.Message().ID]--
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	b.Handle(&btnDelete, func(c tele.Context) error {
		return c.Delete()
	})

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
