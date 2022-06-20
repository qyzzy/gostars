package admin

import "gostars/service"

type ApiGroup struct {
	UserApi
	ArticleApi
	CommentApi
	CategoryApi
}

type UserApi struct {
}

type ArticleApi struct {
}

type CommentApi struct {
}

type CategoryApi struct {
}

var (
	adminUserService     = new(service.UserService)
	adminArticleService  = new(service.ArticleService)
	adminCommentService  = new(service.CommentService)
	adminCategoryService = new(service.CategoryService)
)
