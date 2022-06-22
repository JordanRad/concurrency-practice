package main

import (
	"fmt"

	"github.com/JordanRad/concurrency-practice/cmd/fast_fetch/internal/server"
)

type Operations interface {
	FetchUsers() []string
	FetchPayments() []string
	FetchDeliveryTracker() []string
}

func main() {
	serverHandler := server.NewServer()
	resp, err := serverHandler.FetchAllSync()
	fmt.Println(resp, err)
	fmt.Println("Hello from Fast Fetcher")
}
