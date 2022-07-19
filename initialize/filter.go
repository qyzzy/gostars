package initialize

import (
	"github.com/importcjj/sensitive"
	"gostars/global"
	"gostars/utils"
	"log"
)

func init() {
	global.GFilter = sensitive.New()
	err := global.GFilter.LoadWordDict(utils.WordDictPath)
	if err != nil {
		log.Println(err)
	}
}
