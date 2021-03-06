package api

import (
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/service"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/convert"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
	"github.com/noChaos1012/tour/blog_service/pkg/upload"
)

//上传文件
func UploadFile(c *gin.Context) () {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustUInt32()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails([]string{err.Error()}))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c,"svc.UploadFile err:%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails([]string{err.Error()}))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
