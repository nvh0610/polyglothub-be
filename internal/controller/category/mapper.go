package category

import (
	"learn/internal/entity"
	"learn/internal/router/payload/request"
)

func ToModelCategoryEntity(category *request.CreateCategoryRequest) *entity.Category {
	return &entity.Category{
		Name:   category.Name,
		UserID: category.UserId,
	}
}
