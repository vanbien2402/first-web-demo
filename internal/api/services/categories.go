package services

import "github.com/vanbien2402/first-web-demo/internal/api/domain"

type categoryService struct {
	categoryRepo domain.CategoryRepository
}

//NewCategoryService init Category Service
func NewCategoryService(categoryRepo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
