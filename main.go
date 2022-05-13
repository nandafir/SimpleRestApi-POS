package main

import (
	"anaconda/config"
	"anaconda/server"

	_ "github.com/lib/pq"
)

func main() {
	config.Init()
	server.Start()
}
