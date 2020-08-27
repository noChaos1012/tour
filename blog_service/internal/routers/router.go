package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/noChaos1012/tour/blog_service/docs"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/middleware"
	"github.com/noChaos1012/tour/blog_service/internal/routers/api"
	v1 "github.com/noChaos1012/tour/blog_service/internal/routers/api/v1"
	"github.com/noChaos1012/tour/blog_service/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBucket(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {

	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		gin.SetMode(gin.ReleaseMode)
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.RequestTimeOut))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.Article{}
	tag := v1.Tag{}

	r.GET("/auth", api.GetAuth)
	r.POST("/upload/file", api.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //配置文件存储位置地址

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
