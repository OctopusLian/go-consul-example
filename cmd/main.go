package main

import (
	"log"
	"strings"
	"time"

	"github.com/LIYINGZHEN/go-consul-example/configs"
	"github.com/LIYINGZHEN/go-consul-example/internal/app/http"
	"github.com/LIYINGZHEN/go-consul-example/internal/app/service"
)

func main() {
	c := configs.C

	addrs := strings.Split(c.Redis.Addrs, ";")
	ttl := time.Second * c.Server.TTL

	s, err := service.New(addrs, ttl)
	if err != nil {
		log.Fatal(err)
	}

	server := http.AppServer{
		Service: s,
	}
	server.Run(c.Server.Port)
}
