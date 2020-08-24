package dao

import "github.com/noChaos1012/tour/blog_service/internal/model"

func (d *Dao) GetArticleTagByArticleID(articleID uint32) (model.ArticleTag, error) {
	at := model.ArticleTag{ArticleId: articleID}
	return at.GetByArticleID(d.engine)
}

func (d *Dao) ListArticleTagByTagID(tagID uint32) ([]*model.ArticleTag, error) {
	at := model.ArticleTag{TagId: tagID}
	return at.ListByTagID(d.engine)
}

func (d *Dao) ListArticleTagByArticleID(articleIDs []uint32) ([]*model.ArticleTag, error) {
	at := model.ArticleTag{}
	return at.ListByArticleIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBy string) error {
	at := model.ArticleTag{ArticleId: articleID, TagId: tagID, Model: &model.Model{CreatedBy: createBy}}
	return at.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	at := model.ArticleTag{ArticleId: articleID}
	values := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return at.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	at := model.ArticleTag{ArticleId: articleID}
	return at.DeleteOne(d.engine)
}
