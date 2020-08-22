package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Create(c *gin.Context) {}
func (a Article) Get(c *gin.Context) {
	fmt.Println("日志测试")
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Delete(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
