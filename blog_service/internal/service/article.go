package service

import (
	"github.com/noChaos1012/tour/blog_service/internal/dao"
	"github.com/noChaos1012/tour/blog_service/internal/model"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gt=0"`
	State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gt=0"`
	State uint8  `form:"state,default=1" binding:"oneof= 0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gt=0"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gt=0"`
	TagID         uint32 `form:"tag_id" binding:"required,gt=0"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gt=0"`
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	Tag           *model.Tag `json:"tag"`
}

func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := s.dao.CreateArticle(&dao.Article{
		TagID:         param.TagID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		CreatedBy:     param.CreatedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}

	err = s.dao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := s.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,
		TagID:         param.TagID,
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		ModifiedBy:    param.ModifiedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	err = s.dao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetArticle(param *ArticleRequest) (*Article, error) {

	a, err := s.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	at, err := s.dao.GetArticleTagByArticleID(param.ID)
	if err != nil {
		return nil, err
	}

	t, err := s.dao.GetTag(at.TagId, model.STATE_OPEN)

	if err != nil {
		return nil, err
	}
	return &Article{
		ID:            a.ID,
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CoverImageUrl: a.CoverImageUrl,
		Tag:           t,
	}, nil
}

func (s *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	count, err := s.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := s.dao.GetArticleListByTagId(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var atcs []*Article
	for _, article := range articles {
		atc := &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag: &model.Tag{
				Model: &model.Model{ID: article.TagID},
				Name:  article.TagName,
			},
		}
		atcs = append(atcs, atc)
	}

	return atcs, count, nil
}

func (s *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := s.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}
	return s.dao.DeleteArticleTag(param.ID)
}
