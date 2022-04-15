package main

import (
	"ippool_api/db/redis"
	"ippool_api/route"
	"ippool_api/utils/log"
)

func init() {
	log.NewLogger(3)
	if err := redis.Connect(); err != nil {
		log.GlobalLog.Errorf("Redis >>> %s", err)
		return
	}
}

func main() {
	route.Run()
}
