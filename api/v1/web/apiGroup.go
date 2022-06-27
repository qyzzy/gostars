package web

import "gostars/service"

type ApiGroup struct {
	UserApi
	ArticleApi
	CommentApi
	CategoryApi
	LikeArticleApi
}

type UserApi struct {
}

type ArticleApi struct {
}

type CommentApi struct {
}

type CategoryApi struct {
}

type LikeArticleApi struct {
}

var (
	webUserService     = new(service.UserService)
	webArticleService  = new(service.ArticleService)
	webCategoryService = new(service.CategoryService)
	webCommentService  = new(service.CommentService)
	webJwtService      = new(service.JwtService)
	webFavoriteService = new(service.FavoriteService)
)
