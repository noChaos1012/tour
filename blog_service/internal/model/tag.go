package model

import "github.com/noChaos1012/tour/blog_service/pkg/app"

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
