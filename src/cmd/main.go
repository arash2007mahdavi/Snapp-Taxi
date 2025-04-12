package main

import (
	"snapp/config"
	"snapp/servers"
)

func main() {
	cfg := config.GetConfig()
	servers.NewServer(cfg)
}