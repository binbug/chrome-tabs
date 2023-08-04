package http

import (
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"strings"
)

type PageHandler struct {
}

func (h PageHandler) Init(_ *handler.InitConfig) {

}

func (h PageHandler) AlfredSearch(qs handler.Queries, browserItems []*alfred.Item, _ []*alfred.Item) []*alfred.Item {
	if qs.Len() == 1 && (strings.HasSuffix(qs.First(), "http://") || strings.HasPrefix(qs.First(), "https://")) {
		return []*alfred.Item{
			{
				Title:    qs.First(),
				Subtitle: "Open in browser",
				Arg:      qs.First() + "&create=true",
			},
		}
	}

	return nil
}

func (h PageHandler) Order() int {
	return 0
}

func (h PageHandler) IsRod() bool {
	return false
}

func init() {
	handler.RegisterHandler(&PageHandler{})
}
