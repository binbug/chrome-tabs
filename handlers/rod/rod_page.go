package rod

import (
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"strings"
)

type PageHandler struct {
	cfg *handler.InitConfig
}

func (h *PageHandler) Init(cfg *handler.InitConfig) {
	h.cfg = cfg
}

func (h *PageHandler) Order() int {
	return -1
}

func (h *PageHandler) AlfredSearch(qs handler.Queries, browserItems []*alfred.Item, _ []*alfred.Item) []*alfred.Item {
	items := make([]*alfred.Item, 0)

	for _, page := range browserItems {
		if match(page, qs) {
			items = append(items, page)
		}
	}

	return items
}

func match(page *alfred.Item, qs []string) bool {
	for _, q := range qs {
		if !matchOne(page, q) {
			return false
		}
	}
	return true
}

func matchOne(page *alfred.Item, q string) bool {
	return q == "" || strings.Contains(strings.ToLower(page.Title), strings.ToLower(q)) || strings.Contains(strings.ToLower(page.Subtitle), strings.ToLower(q))
}

func init() {
	handler.RegisterHandler(&PageHandler{})
}
