package proxy

import (
	"inproxy/filter"
	"net/http"
)

type Proxy interface {
	http.Handler
}

func New(f filter.Filter) Proxy {
	return nil
}
