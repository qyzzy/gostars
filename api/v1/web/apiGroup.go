package web

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
	webUserService     = new(service.UserService)
	webArticleService  = new(service.ArticleService)
	webCategoryService = new(service.CategoryService)
	webCommentService  = new(service.CommentService)
)
