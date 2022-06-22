package main

import (
	"fmt"

	"github.com/JordanRad/concurrency-practice/cmd/fast_fetch/internal/server"
)

func main() {
	serverHandler := server.NewServer()
	fmt.Println("Fetching asynchronously....")
	_, _ = serverHandler.FetchAll()
}
