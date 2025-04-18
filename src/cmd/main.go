package main

import (
	"snapp/config"
	"snapp/database"
	"snapp/servers"
)

func main() {
	cfg := config.GetConfig()
	database.InitRedis(cfg)
	defer database.CloseRedis()
	err := database.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()
	servers.NewServer(cfg)
}