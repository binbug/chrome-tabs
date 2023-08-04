package cmd

import (
	"chrome-tabs/internal/config"
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/go-rod/rod/lib/proto"
	"github.com/xyproto/simplebolt"
	"log"
	"sort"
	"strings"
)

type application struct {
	cfg          *handler.InitConfig
	browser      *rod.Browser
	pageHandlers []handler.PageHandler
}

func newApp(cfg *config.Config) *application {
	db, err := simplebolt.New("bolt.db")
	if err != nil {
		log.Panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	handlers := handler.AllHandlers()

	sort.Slice(handlers, func(i, j int) bool {
		return handlers[i].Order() < handlers[j].Order()
	})

	handlerInitConfig := &handler.InitConfig{
		DB:     db,
		Engine: r,
		Cfg:    cfg,
	}

	app := &application{
		pageHandlers: handlers,
		cfg:          handlerInitConfig,
	}

	app.initBrowser()

	for _, h := range handlers {
		h.Init(handlerInitConfig)
	}

	r.POST("/list", app.listPage)
	r.POST("/activate", app.activePage)

	return app
}

func (app *application) initBrowser() {
	log.Println("initBrowser")
	l := launcher.NewUserMode()
	if app.cfg.Cfg.RodBin != "" {
		l.Bin(app.cfg.Cfg.RodBin)
	}

	bins, _ := l.GetFlags(flags.Bin)
	log.Println("rod bin is :", bins)
	wsURL := l.MustLaunch()
	log.Println("wsURL:", wsURL)
	app.browser = rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice()
}

func (app *application) listPage(c *gin.Context) {

	q := c.PostForm("q")
	log.Println("q:", q)

	items := &alfred.Items{}
	items.Items = make([]*alfred.Item, 0)

	browserPages := app.pages()

	for _, h := range app.pageHandlers {
		qs := strings.Split(q, " ")
		items.Items = append(items.Items, h.AlfredSearch(qs, browserPages, items.Items)...)
	}

	c.String(200, items.ToXML())

}

func (app *application) activePage(c *gin.Context) {
	create := c.PostForm("create")
	u := c.PostForm("u")

	if create == "true" {
		log.Println("create1:", u)
		app.browser.MustPage(u).MustActivate()
	} else {
		app.browser.MustPageFromTargetID(proto.TargetTargetID(u)).MustActivate()

	}

	c.String(200, "ok")
}

func (app *application) Start() {
	log.Println("chrome tab started, listening on port:", app.cfg.Cfg.Port)
	err := app.cfg.Engine.Run(fmt.Sprintf(":%d", app.cfg.Cfg.Port))
	log.Panic(err)
}
