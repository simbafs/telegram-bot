package counter

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

var (
	selector = &tele.ReplyMarkup{}

	btnAdd   = selector.Data("+", "add")
	btnZero  = selector.Data("0", "zero")
	btnMinus = selector.Data("-", "minus")
	// btnSet    = selector.Data("set", "set")
	btnDelete = selector.Data("x", "delete")
)
var counter = make(map[int]int)

func Command(b *tele.Bot) *tele.Group {
	g := b.Group()
	selector.Inline(
		selector.Row(btnAdd, btnZero, btnMinus, btnDelete),
	)
	g.Handle("/counter", func(c tele.Context) error {
		counter[c.Message().ID] = 0
		return c.Send(fmt.Sprint(0), selector)
	})

	g.Handle(&btnAdd, func(c tele.Context) error {
		counter[c.Message().ID]++
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	g.Handle(&btnZero, func(c tele.Context) error {
		counter[c.Message().ID] = 0
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	g.Handle(&btnMinus, func(c tele.Context) error {
		counter[c.Message().ID]--
		return c.Edit(fmt.Sprint(counter[c.Message().ID]), selector)
	})

	g.Handle(&btnDelete, func(c tele.Context) error {
		return c.Delete()
	})

	return g
}
