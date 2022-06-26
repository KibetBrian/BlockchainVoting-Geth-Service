package main

import (
	"github.com/KibetBrian/geth/api"
)


func main() {
	server := api.NewServer()
	server.Start("127.0.0.1:8000")
}
