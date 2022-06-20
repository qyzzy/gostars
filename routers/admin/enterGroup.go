package admin

import "gostars/api/v1/admin"

type RouterGroup struct {
	AdminRouterGroup
	LoggerRouterGroup
}

var adminApiGroup admin.ApiGroup
