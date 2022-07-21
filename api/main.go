package main

import (
	"tm/internal/conf"
	"tm/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
