package main

import (
	"github.com/KibetBrian/geth/api"
)


func main() {
	server := api.NewServer()
	server.Start(":8000")
}
