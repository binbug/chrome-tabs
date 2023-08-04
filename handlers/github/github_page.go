package github

import (
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"fmt"
)

type PageHandler struct {
}

func (h *PageHandler) Init(_ *handler.InitConfig) {

}

func (h *PageHandler) Order() int {
	return 0
}

func (h *PageHandler) IsRod() bool {
	return false
}

func (h *PageHandler) AlfredSearch(qs handler.Queries, browserItems []*alfred.Item, existsItems []*alfred.Item) []*alfred.Item {
	if qs.First() == "github" {
		githubQ := qs.Remains()
		return []*alfred.Item{
			{
				Uid:      "github",
				Arg:      fmt.Sprintf("https://github.com/search?q=%s&create=true", githubQ),
				Title:    "search " + githubQ + " on github",
				Subtitle: fmt.Sprintf("https://github.com/search?q=%s", githubQ),
				Icon:     "github.png",
			},
		}
	}
	return nil
}

func init() {
	handler.RegisterHandler(&PageHandler{})
}
