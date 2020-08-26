package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noChaos1012/tour/blog_service/pkg/app"
	"github.com/noChaos1012/tour/blog_service/pkg/errcode"
	"github.com/noChaos1012/tour/blog_service/pkg/limter"
)

func RateLimiter(l limter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}

		}
	}
}
