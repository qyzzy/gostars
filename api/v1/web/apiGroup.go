package web

import "gostars/service"

type ApiGroup struct {
	UserApi
	ArticleApi
	CommentApi
	CategoryApi
	LikeArticleApi
	NavMenuApi
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

type NavMenuApi struct {
}

var (
	webUserService          = new(service.UserService)
	webArticleService       = new(service.ArticleService)
	webCategoryService      = new(service.CategoryService)
	webCommentService       = new(service.CommentService)
	webJwtService           = new(service.JwtService)
	webFavoriteService      = new(service.FavoriteService)
	webLevel0NavMenuService = new(service.Level0NavMenuService)
	webLevel1NavMenuService = new(service.Level1NavMenuService)
)
