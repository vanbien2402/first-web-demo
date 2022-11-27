package controllers

import "github.com/vanbien2402/first-web-demo/internal/api/domain"

type categoryController struct {
	categoryService domain.CategoryService
}

//NewCategoryController init Category Controller
func NewCategoryController(categoryService domain.CategoryService) domain.CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}
