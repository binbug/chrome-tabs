package cmd

import (
	"chrome-tabs/pkg/alfred"
	retry2 "chrome-tabs/pkg/retry"
	"github.com/go-rod/rod"
	"log"
	"reflect"
	"time"
)

func (app *application) pages() []*alfred.Item {

	begin := time.Now()
	pages := make([]*alfred.Item, 0)

	rodPages, err := retry2.Do[rod.Pages](func() (rod.Pages, error) {
		return app.browser.Pages()
	}, retry2.OnRetry(func(n uint, err error) {
		log.Printf("retrying after error: %v", err)
		app.initBrowser()
	}), retry2.Attempts(1))

	if err != nil {
		log.Println(reflect.TypeOf(err))
		log.Panicln(err)
	}

	for _, page := range rodPages {
		chromePage := ChromePage{
			TargetID: string(page.TargetID),
			URL:      page.MustInfo().URL,
			Title:    page.MustInfo().Title,
		}
		pages = append(pages, chromePage.toAlfredItem())
	}

	elapsed := time.Since(begin)
	log.Printf("pages took %s", elapsed)
	return pages
}

type ChromePage struct {
	TargetID string `json:"target_id"`
	URL      string `json:"url"`
	Title    string `json:"title"`
}

func (page *ChromePage) toAlfredItem() *alfred.Item {
	return &alfred.Item{
		Uid:      page.TargetID,
		Title:    page.Title,
		Subtitle: page.URL + "[exist]",
		Arg:      string(page.TargetID) + "&create=false",
		Icon:     "icon.png",
	}
}
