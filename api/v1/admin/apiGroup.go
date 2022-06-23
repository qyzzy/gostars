package admin

import "gostars/service"

type ApiGroup struct {
	UserApi
	ArticleApi
	CommentApi
	CategoryApi
	JwtApi
	LoggerApi
	UploadApi
}

type UserApi struct {
}

type ArticleApi struct {
}

type CommentApi struct {
}

type CategoryApi struct {
}

type JwtApi struct {
}

type LoggerApi struct {
}

type UploadApi struct {
}

var (
	adminUserService     = new(service.UserService)
	adminArticleService  = new(service.ArticleService)
	adminCommentService  = new(service.CommentService)
	adminCategoryService = new(service.CategoryService)
	adminJwtService      = new(service.JwtService)
	adminImageService    = new(service.ImageService)
)
