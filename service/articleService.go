package service

import (
	"gostars/global"
	"gostars/models"
	"gostars/utils/code"
)

type ArticleService struct {

}

func (articleService *ArticleService) CreateArticle(data *models.Article) int {
	err := global.GDb.Table(models.ArticleTableName()).
		Create(&data).Error
	if err != nil {
		return code.ERROR
	}
	return code.SUCCESS
}

func (articleService *ArticleService) GetArticles(pageSize, pageNum int) ([]models.Article, int, int64) {
	var articleList []models.Article
	var err error
	var total int64

	err = global.GDb.Table(models.ArticleTableName()).
		Select("id, title, summary, created_at, updated_at, deleted_at, content, click_count, status, " +
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, category_name").
		Limit(pageNum).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&articleList).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	global.GDb.Model(&articleList).Count(&total)

	return articleList, code.SUCCESS, total
}

func (articleService *ArticleService) GetArticlesByTitle(title string, pageSize, pageNum int) ([]models.Article, int, int64) {
	var articleList []models.Article
	var err error
	var total int64

	err = global.GDb.Table(models.ArticleTableName()).
		Select("id, title, summary, created_at, updated_at, deleted_at, content, click_count, status, "+
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, "+
		"category_name").
		Limit(pageNum).Offset((pageNum-1)*pageSize).Order("created_at desc").
		Where("title like ?", title+"%").Find(&articleList).Model(&total).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	return articleList, code.SUCCESS, total
}

func (articleService *ArticleService) GetArticlesByCategory(id, pageSize, pageNum int) ([]models.Article, int, int64) {
	var articleList []models.Article
	var err error
	var total int64

	err = global.GDb.Table(models.ArticleTableName()).
		Select("id, title, summary, created_at, updated_at, deleted_at, content, click_count, status, "+
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, "+
		"category_name").
		Limit(pageNum).Offset((pageNum-1)*pageSize).Order("created_at desc").
		Where("category_id = ?", id).Find(&articleList).Model(&total).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	return articleList, code.SUCCESS, total
}

// todo
func (articleService *ArticleService) GetArticleClickCount(id int) {

}

func (articleService *ArticleService) EditArticle(id int, data *models.Article) int {
	var article models.Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := global.GDb.Table(models.ArticleTableName()).
		Model(&article).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return code.ERROR
	}

	return code.SUCCESS
}

func (articleService *ArticleService) DeleteArticle(id int) int {
	var article models.Article
	err := global.GDb.Table(models.ArticleTableName()).
		Where("id = ?", id).Delete(&article).Error

	if err != nil {
		return code.ERROR
	}

	return code.SUCCESS
}