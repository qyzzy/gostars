package routers

import (
	"gostars/routers/admin"
	"gostars/routers/user"
)

type RouterGroup struct {
	User  user.RouterGroup
	Admin admin.RouterGroup
}
