package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LIYINGZHEN/go-consul-example/internal/service"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	addrsStr := flag.String("addrs", "localhost:6379", "Redis addrs")
	ttl := flag.Duration("ttl", time.Second*15, "Service TTL check duration")
	flag.Parse()

	addrs := strings.Split(*addrsStr, ";")

	s, err := service.New(addrs, *ttl)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", s)

	l := fmt.Sprintf(":%d", *port)
	log.Print("Listening on ", l)
	log.Fatal(http.ListenAndServe(l, nil))
}
