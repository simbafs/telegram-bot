package pixiv

import (
	"fmt"
	"log"
	"regexp"

	"bot/lib/pixiv"
	"bot/lib/translate"

	tele "gopkg.in/telebot.v3"
)

func getTranslateNovel(id string) (string, error) {
	novel, err := pixiv.GetNovelByID(id)
	if err != nil {
		return "", fmt.Errorf("Error occur when fetching novel: %w", err)
	}

	translated, err := translate.Translate(novel.Body.Content)
	if err != nil {
		return novel.Body.Content, fmt.Errorf("Error occur when trnaslate novel: %w", err)
	}

	return translated, nil
}

func Command(b *tele.Bot) *tele.Group {
	g := b.Group()

	g.Handle("/pixiv", func(c tele.Context) error {
		if len(c.Args()) < 1 {
			return c.Send("miss id, /pixiv <id>")
		}

		id := c.Args()[0]

		err := c.Send(fmt.Sprintf("Getting novel by id %s ......", id))
		if err != nil {
			return err
		}

		if translated, err := getTranslateNovel(id); err != nil {
			log.Print(err)
			return c.Send(fmt.Sprintf("Some errors occur when translate novel, %s\n-------------\n%s", err.Error(), translated))
		} else {
			return c.Send(fmt.Sprintf("https://www.pixiv.net/novel/show.php?id=%s\n---------------------------------------------------------------------------\n%s", id, translated))
		}
	})

	REid := regexp.MustCompile(`https:\/\/www\.pixiv\.net\/novel\/show\.php\?id=([0-9]+)`)
	g.Handle(tele.OnText, func(c tele.Context) error {
		matches := REid.FindSubmatch([]byte(c.Text()))
		if len(matches) < 2 {
			return nil
		}

		id := string(matches[1])

		err := c.Send(fmt.Sprintf("Getting novel by id %s ......", id))
		if err != nil {
			return err
		}

		if translated, err := getTranslateNovel(id); err != nil {
			log.Print(err)
			return c.Send(fmt.Sprintf("Some errors occur when translate novel, %s\n-------------\n%s", err.Error(), translated))
		} else {
			return c.Send(fmt.Sprintf("https://www.pixiv.net/novel/show.php?id=%s\n---------------------------------------------------------------------------\n%s", id, translated))
		}
	})

	return g
}
