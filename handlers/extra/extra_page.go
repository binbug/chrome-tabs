package extra

import (
	"chrome-tabs/internal/handler"
	"chrome-tabs/pkg/alfred"
	"github.com/binbug/go-utils/jsonutils"
	"github.com/gin-gonic/gin"
	"github.com/xyproto/simplebolt"
	"go.etcd.io/bbolt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type PageHandler struct {
	kv          *simplebolt.KeyValue
	db          *simplebolt.Database
	regexpCache map[string]*regexp.Regexp
}

type PageModel struct {
	URL            string `json:"url"`
	Title          string `json:"title"`
	MatchRegexp    string `json:"match_regexp"`
	OverwriteTitle bool   `json:"overwrite_title"` //复写title
}

func (page *PageModel) matchOne(text string) bool {
	text = strings.ToLower(text)
	return strings.Contains(strings.ToLower(page.Title), text) || strings.Contains(strings.ToLower(page.URL), text)
}

func (page *PageModel) match(qs []string) bool {
	for _, q := range qs {
		if !page.matchOne(q) {
			return false
		}
	}
	return true
}

func (page *PageModel) toAlfredItem() *alfred.Item {
	return &alfred.Item{
		Uid:      page.URL,
		Title:    page.Title,
		Subtitle: page.URL,
		Arg:      page.URL + "&create=true",
		Icon:     "icon.png",
	}
}

const BucketPageInfo = "extra-page-info"

func (h *PageHandler) Init(cfg *handler.InitConfig) {
	kv, err := simplebolt.NewKeyValue(cfg.DB, BucketPageInfo)
	if err != nil {
		log.Fatal(err)
	}
	h.kv = kv
	h.db = cfg.DB
	h.regexpCache = make(map[string]*regexp.Regexp)

	if cfg.Engine == nil {
		log.Println("r is nil")
		return
	}

	g := cfg.Engine.Group("/extra-page")

	g.POST("/add", h.Add)
	g.POST("/delete", h.Delete)
	g.GET("/list", h.List)
}

func (h *PageHandler) Order() int {
	return 0
}

func (h *PageHandler) IsRod() bool {
	return false
}

func (h *PageHandler) Search(qs []string) []PageModel {
	pages := make([]PageModel, 0)
	(*bbolt.DB)(h.db).View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketPageInfo))

		bucket.ForEach(func(k, v []byte) error {
			page, _ := jsonutils.FromBytes[PageModel](v)
			if page.match(qs) {
				pages = append(pages, page)
			}
			return nil
		})

		return nil // Return from View function
	})

	return pages
}

func (h *PageHandler) AlfredSearch(qs handler.Queries, browserItems []*alfred.Item, existsItems []*alfred.Item) []*alfred.Item {
	pages := h.Search(qs)
	items := make([]*alfred.Item, 0)
	for _, page := range pages {
		item := h.filter(page, existsItems)
		if item != nil {
			continue
		}

		item = h.filter(page, browserItems)
		if item != nil {
			items = append(items, item)
			continue
		}

		items = append(items, page.toAlfredItem())
	}
	return items
}

func (h *PageHandler) filter(page PageModel, items []*alfred.Item) *alfred.Item {
	for _, item := range items {

		exp, ok := h.regexpCache[page.MatchRegexp]
		if !ok {
			exp = regexp.MustCompile(page.MatchRegexp)
			h.regexpCache[page.MatchRegexp] = exp
		}

		b := exp.MatchString(item.Subtitle)
		if b {
			if page.OverwriteTitle {
				item.Title = page.Title
			}

			return item
		}
	}
	return nil
}

func (h *PageHandler) Add(c *gin.Context) {
	page := PageModel{}
	c.BindJSON(&page)
	//page := PageModel{
	//	URL:         c.PostForm("url"),
	//	Title:       c.PostForm("title"),
	//	MatchRegexp: c.PostForm("match_regexp"),
	//}

	h.kv.Set(page.URL, jsonutils.ToJSON(page))
	c.String(http.StatusOK, "ok")
}

func (h *PageHandler) Delete(c *gin.Context) {
	key := c.PostForm("key")
	if key == "" {
		c.String(http.StatusBadRequest, "key is empty")
		return
	}

	value, _ := h.kv.Get(key)

	err := h.kv.Del(key)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	deleted, _ := jsonutils.FromString[map[string]interface{}](value)
	deletedMap := map[string]interface{}{
		"deleted": deleted,
	}

	c.JSON(http.StatusOK, deletedMap)

}

func (h *PageHandler) List(c *gin.Context) {
	pages := make([]PageModel, 0)
	(*bbolt.DB)(h.db).View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketPageInfo))

		bucket.ForEach(func(k, v []byte) error {
			page, _ := jsonutils.FromBytes[PageModel](v)
			pages = append(pages, page)
			return nil
		})

		return nil // Return from View function
	})

	c.JSON(http.StatusOK, pages)
}

func init() {
	handler.RegisterHandler(&PageHandler{})
}
