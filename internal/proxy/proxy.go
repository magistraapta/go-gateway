package proxy

import (
	"net/http"
)

type Proxy struct {
	client *http.Client
}

func NewProxy(urls []string) *Proxy {
	return &Proxy{
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
			},
		},
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
