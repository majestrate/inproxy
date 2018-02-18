package main

import (
	"inproxy/config"
	"inproxy/filter"
	"inproxy/log"
	"inproxy/proxy"
	"net/http"
	"time"
)

func main() {
	conf := new(config.Config)
	err := conf.Load("inproxy.ini")
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}
	f := filter.New()
	proxy := proxy.New(f)
	go func(h http.Handler) {
		for {
			bind := conf.HTTP.BindAddr
			err := http.ListenAndServe(bind, h)
			if err != nil {
				log.Errorf("ListenAndServe: %s", err.Error())
				time.Sleep(time.Second)
			}
		}
	}(proxy)
}
