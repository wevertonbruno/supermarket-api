package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/wevertonbruno/supermarket-api-go/adapters/web"
)

func main() {
	server := web.NewServer("8080")
	server.Run()
}
