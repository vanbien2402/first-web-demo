package repositories

import (
	"github.com/vanbien2402/first-web-demo/internal/api/domain"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

//NewCategoryRepository init Category Repository
func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
