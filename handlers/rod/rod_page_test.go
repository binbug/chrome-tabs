package rod

import (
	"github.com/go-rod/rod"
	"testing"
)

func TestRod(t *testing.T) {
	browser := rod.New().MustConnect().NoDefaultDevice()

	pages := browser.MustPages()

	for _, page := range pages {
		t.Log(page.MustInfo().Title)
	}
	select {}
}
