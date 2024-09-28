package case19

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const prefix = "article_"

type ArticleService struct {
	rdb redis.Cmdable
	DB  *gorm.DB
}

func NewArticleService(rdb redis.Cmdable, DB *gorm.DB) *ArticleService {
	return &ArticleService{rdb: rdb, DB: DB}
}

func (s *ArticleService) UpdateDB(article *Article) (*Article, error) {
	if err := s.DB.Updates(article).Error; err != nil {
		log.Printf("update db error: %v", err.Error())
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) UpdateCache(ctx context.Context, article *Article) (*Article, error) {
	key := fmt.Sprintf("%s%d", prefix, article.ID)
	if err := s.rdb.Set(ctx, key, article, 0).Err(); err != nil {
		log.Printf("update cache error: %v", err.Error())
		return nil, err
	}
	return nil, nil
}

type Article struct {
	ID      uint   `gorm:"primaryKey", json:"id"`
	Title   string `gorm:"unique;type:varchar(255)", json:"title"`
	Author  string `gorm:"type:varchar(255)", json: "author"`
	Content string `gorm:"type:TEXT", json: "content"`
}

func (Article) TableName() string {
	return "article"
}

func (a *Article) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Article) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
