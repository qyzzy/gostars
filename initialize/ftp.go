package initialize

import (
	"github.com/dutchcoders/goftp"
	"gostars/global"
	"gostars/utils"
	"log"
	"time"
)

func initFtp() {
	var err error
	global.GFtp, err = goftp.Connect(utils.FtpAddress)
	if err != nil {
		log.Println(err)
	}

	err = global.GFtp.Login(utils.FtpUser, utils.FtpPassword)
	if err != nil {
		log.Println(err)
	}

	go keepAlive()
}

func keepAlive() {
	time.Sleep(time.Duration(utils.HeartbeatTime) * time.Second)
	_ = global.GFtp.Noop()
}
