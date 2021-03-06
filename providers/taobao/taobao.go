package taobao

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("taobao", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Taobao.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://s.taobao.com/search?q=%s", url.QueryEscape(q))
}
