package handler

import (
	"chrome-tabs/internal/config"
	"chrome-tabs/pkg/alfred"
	"github.com/gin-gonic/gin"
	"github.com/xyproto/simplebolt"
	"log"
	"reflect"
	"strings"
)

type Queries []string

func (q Queries) First() string {
	if len(q) > 0 {
		return q[0]
	}

	return ""
}

func (q Queries) Remains() string {
	return strings.Join(q[1:], " ")
}

func (q Queries) Contains(s string) bool {
	for _, v := range q {
		if v == s {
			return true
		}
	}

	return false
}

func (q Queries) Len() int {
	return len(q)
}

type PageHandler interface {
	Init(config *InitConfig)
	AlfredSearch(qs Queries, allBrowserTabs []*alfred.Item, existsItems []*alfred.Item) []*alfred.Item
	Order() int
}

type InitConfig struct {
	DB     *simplebolt.Database
	Engine *gin.Engine
	Cfg    *config.Config
}

var handlers = make([]PageHandler, 0)

func RegisterHandler(h PageHandler) {
	log.Println(reflect.TypeOf(h))
	handlers = append(handlers, h)
}

func AllHandlers() []PageHandler {
	return handlers
}
