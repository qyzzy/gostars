package main

import (
	"gostars/global"
	_ "gostars/initialize"
	_ "gostars/routers"
)

func main() {
	db, _ := global.GDb.DB()
	defer db.Close()

	defer global.GRabbitMQ.Close()

	defer global.GRedis.Close()
}
