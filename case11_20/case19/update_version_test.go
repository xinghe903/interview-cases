package case19

import (
	"context"
	"fmt"
	"interview-cases/test"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Case19TestSuite struct {
	suite.Suite
	db  *gorm.DB
	rdb redis.Cmdable
}

func (s *Case19TestSuite) SetupTest() {
	fmt.Printf("start test case19")
	s.db = test.InitDB()
	s.rdb = test.InitRedis()
	if err := s.db.AutoMigrate(&Article{}); err != nil {
		s.T().Fatal(err)
	}
	ctx := context.Background()
	if err := s.rdb.FlushAll(ctx).Err(); err != nil {
		s.T().Fatal(err)
	}
	// if err := s.rdb.Set(ctx, "article_11", &Article{
	// 	ID:      11,
	// 	Title:   "Test Article111",
	// 	Content: "This is a test article.",
	// 	Author:  "xinghe",
	// }, 0).Err(); err != nil {
	// 	s.T().Fatal(err)
	// }
	// s.db.Create(&Article{
	// 	ID:      11,
	// 	Title:   "Test Article111",
	// 	Content: "This is a test article.",
	// 	Author:  "xinghe",
	// })
}

func (s *Case19TestSuite) TestUpdate() {
	assert.Equal(s.T(), 5, s.db.Update)
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(Case19TestSuite))
}
