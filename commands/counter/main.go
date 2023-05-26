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

func calc(str string, n int) string {
	var i int
	fmt.Sscanf(str, "%d", &i)
	return fmt.Sprint(i + n)
}

func Command(b *tele.Bot) *tele.Group {
	g := b.Group()
	selector.Inline(
		selector.Row(btnAdd, btnZero, btnMinus, btnDelete),
	)
	g.Handle("/counter", func(c tele.Context) error {
		return c.Send(fmt.Sprint(0), selector)
	})

	g.Handle(&btnAdd, func(c tele.Context) error {
		return c.Edit(calc(c.Text(), 1), selector)
	})

	g.Handle(&btnZero, func(c tele.Context) error {
		return c.Edit("0", selector)
	})

	g.Handle(&btnMinus, func(c tele.Context) error {
		return c.Edit(calc(c.Text(), -1), selector)
	})

	g.Handle(&btnDelete, func(c tele.Context) error {
		return c.Delete()
	})

	return g
}
