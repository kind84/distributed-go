package main

import (
	"net/http"

	"github.com/kind84/distributed-go/chapter5/goophr/concierge/api"
	"github.com/kind84/distributed-go/chapter5/goophr/concierge/common"
)

func main() {
	common.Log("Adding API handlers...")
	http.HandleFunc("/api/feeder", api.FeedHandler)

	common.Log("Starting feeder...")
	api.StartFeederSystem()

	common.Log("Starting Goophr Concierge server on port :8080...")
	http.ListenAndServe(":8080", nil)
}
