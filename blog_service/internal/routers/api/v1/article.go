package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/service"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/convert"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
	return
}

func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}

	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)
	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)

	vaild, errs := app.BindAndValid(c, &param)
	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c)}
	articles, count, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	response.ToResponseList(articles, count)
	return
}

func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)
	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
