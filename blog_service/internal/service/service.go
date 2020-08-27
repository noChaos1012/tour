package service

import (
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(otgorm.WithContext(ctx, global.DBEngine)),
	}
}
