package http

import (
	stdHttp "net/http"
	"net/url"
)

func NewClient(proxy *url.URL) *stdHttp.Client {
	client := stdHttp.Client{
		Transport: &stdHttp.Transport{
			Proxy: stdHttp.ProxyURL(proxy),
		},
	}

	return &client
}
