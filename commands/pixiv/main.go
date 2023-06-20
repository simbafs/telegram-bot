package pixiv

import (
	"fmt"

    "bot/lib/pixiv"
    "bot/lib/translate"
	tele "gopkg.in/telebot.v3"
)

func Command(b *tele.Bot) *tele.Group {
    g := b.Group()

    g.Handle("/pixiv", func(c tele.Context) error {
        err  := c.Send(fmt.Sprintf("getting novel by id %s", c.Args()[0]))
        if err != nil {
            return err
        }

        novel, err := pixiv.GetNovelByID(c.Args()[0])
        if err != nil {
            return c.Send("Error occur when fetching novel")
        }

        translated, err :=  translate.Translate(novel.Body.Content)
        if err != nil {
            return c.Send(fmt.Sprintf("Error occur when trnaslate novel, here is the content: %s", novel.Body.Content))
        }

        return c.Send(translated)
    })
    
    return g
}

