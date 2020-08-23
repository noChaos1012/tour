package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/service"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/convert"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//@Summary 获取多个标签
//@Produce json
//@Param name query string false "标签名称" maxlength(100)
//@Param state query int false "状态" Enums(0,1) default(1)
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.TagSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	//入参校验与绑定
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid1 errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//获取标签总数
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	//获取标签雷彪
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err:%v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	//序列化结果集
	response.ToResponseList(tags, totalRows)
	return
}

//@Summary 新增标签
//@Produce json
//@Param name body string true "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0,1) default(1)
//@Param created_by body string false "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) { //curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=Go' -F created_by=waschild 表单形式提交

	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)

	if !vaild {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//@Summary 修改标签
//@Produce json
//@Param id path int true "标签ID"
//@Param name body string false "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0,1) default(1)
//@Param modified_by body string true "修改者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) { //curl -X PUT http://127.0.0.1:8000/api/v1/tags/6 -F modified_by=waschil -F state=1
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32(),}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//@Summary 删除标签
//@Produce json
//@Param id path int true "标签ID"
//@Success 200 {array} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) { //curl -X DELETE http://127.0.0.1:8000/api/v1/tags/6
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err:%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
