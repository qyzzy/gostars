package main

import (
	"gostars/global"
	_ "gostars/initialize"
	_ "gostars/routers"
	_ "gostars/service"
)

func main() {
	db, _ := global.GDb.DB()
	defer db.Close()

	defer global.GRabbitMQ.Close()
}
