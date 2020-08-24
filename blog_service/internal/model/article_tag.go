package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (at ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		return err
	}
	return nil
}

func (at ArticleTag) GetByArticleID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", at.ArticleId, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return ArticleTag{}, err
	}
	return articleTag, nil
}

func (at ArticleTag) ListByTagID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("tag_id = ? AND is_del = ?", at.TagId, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (at ArticleTag) ListByArticleIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (at ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	err := db.Model(&at).Where("article_id = ? AND is_del = ?", at.ArticleId, 0).Limit(1).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (at ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", at.Model.ID, 0).Delete(&at).Error; err != nil {
		return err
	}
	return nil
}

func (at ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del = ?", at.ArticleId, 0).Delete(&at).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
