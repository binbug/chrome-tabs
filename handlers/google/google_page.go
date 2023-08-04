package google

import (
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"fmt"
)

type PageHandler struct {
}

func (h *PageHandler) Init(_ *handler.InitConfig) {

}

func (h *PageHandler) AlfredSearch(qs handler.Queries, browserItems []*alfred.Item, _ []*alfred.Item) []*alfred.Item {

	if qs.First() == "google" || qs.First() == "gg" {
		googleQ := qs.Remains()
		return []*alfred.Item{
			{
				Uid:      "google",
				Arg:      fmt.Sprintf("https://www.google.com/search?q=%s&create=true", googleQ),
				Title:    "search " + googleQ + " on google",
				Subtitle: fmt.Sprintf("https://www.google.com/search?q=%s", googleQ),
				Icon:     "google.png",
			},
		}
	}

	return nil
}

func (h *PageHandler) Order() int {
	return 0
}

func (h *PageHandler) IsRod() bool {
	return false
}

func init() {
	handler.RegisterHandler(&PageHandler{})
}
