package dao

import (
	"github.com/noChaos1012/tour/blog_service/internal/model"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	atc := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model: &model.Model{
			CreatedBy: param.CreatedBy,
		},
	}
	return atc.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{
		ID: param.ID,
	}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{
		State: state,
		Model: &model.Model{
			ID: id,
		},
	}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagID(tagID uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(tagID, d.engine)
}

func (d *Dao) GetArticleListByTagId(tagID uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagId(d.engine, tagID, app.GetPageOffset(page, pageSize), pageSize)
}