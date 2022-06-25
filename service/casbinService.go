package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gostars/global"
	"gostars/utils"
	"strings"
	"sync"
)

type CasbinService struct {
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GDb)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(utils.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

func (casbinService *CasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return casbinService.ParamsMatch(name1, name2), nil
}

func (casbinService *CasbinService) ParamsMatch(fullNameKey1, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// find key1
	return util.KeyMatch2(key1, key2)
}
