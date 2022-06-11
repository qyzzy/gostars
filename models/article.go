package models

import (
	"gorm.io/gorm"
	"gostars/utils/code"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Summary      string `gorm:"type:varchar(200);not null" json:"summary"`
	Content      string `gorm:"type:longtext;not null" json:"content"`
	ClickCount   int    `gorm:"type:int;not null" json:"clickcount"`
	Status       int    `gorm:"type:tinyint;not null;default:0" json:"status"`
	AdminID      int    `gorm:"type:int;not null" json:"adminid"`
	IsOriginal   int    `gorm:"type:int;not null" json:"isoriginal"`
	Author       string `gorm:"type:varchar(100);not null" json:"author"`
	OpenComment  bool   `gorm:"type:tinyint;not null;default:0" json:"clickcount"`
	TagList      []Tag  `gorm:"type:text" json:"tagList"`
	Img          string `gorm:"type:varchar(100);not null" json:"img"`
	CommentCount int    `gorm:"type:int;not null" json:"commentcount"`
	CategoryID   int    `gorm:"type:int;not null" json:"categoryid"`
	CategoryName string `gorm:"type:varchar(50)" json:"categoryname"`
}

func articleTableName() string {
	return "articles"
}

func GetArticles(pageSize, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Table(articleTableName()).Select("id, title, summary, created_at, update_at, delete_at, content, click_count, status, " +
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, category_name").Limit(pageNum).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&articleList).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	db.Model(&articleList).Count(&total)

	return articleList, code.SUCCESS, total
}

func GetArticlesByTitle(title string, pageSize, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Table(articleTableName()).Select("id, title, summary, created_at, update_at, delete_at, content, click_count, status, "+
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, "+
		"category_name").Limit(pageNum).Offset((pageNum-1)*pageSize).Order("created_at desc").Where("title like ?", title+"%").Find(&articleList).Model(&total).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	return articleList, code.SUCCESS, total
}

func GetArticlesByCategory(id, pageSize, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64

	err = db.Table(articleTableName()).Select("id, title, summary, created_at, update_at, delete_at, content, click_count, status, "+
		"is_original, author, open_comment, tag_list, img, comment_count, category_id, "+
		"category_name").Limit(pageNum).Offset((pageNum-1)*pageSize).Order("created_at desc").Where("category_id = ?", id).Find(&articleList).Model(&total).Error

	if err != nil {
		return articleList, code.ERROR, total
	}

	return articleList, code.SUCCESS, total
}
