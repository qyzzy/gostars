package admin

import "gostars/api/v1/admin"

type RouterGroup struct {
	AdminRouterGroup
	LoggerRouterGroup
	JwtRouterGroup
	UploadRouterGroup
	CasbinRouterGroup
}

var adminApiGroup admin.ApiGroup
