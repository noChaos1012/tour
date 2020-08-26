/**
限流控制
*/
package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//限流器必须方法
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucket(rules ...LimiterBucketRule) LimiterIface
}

//令牌桶的规则属性
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}

//存储令牌桶与键值对名称映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}
