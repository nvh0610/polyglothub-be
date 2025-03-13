package response

import "learn/internal/entity"

type DetailCategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToDetailCategoryResponse(category *entity.Category) *DetailCategoryResponse {
	return &DetailCategoryResponse{
		ID:        category.Id,
		Name:      category.Name,
		UserId:    category.UserID,
		CreatedAt: category.CreatedAt.String(),
		UpdatedAt: category.UpdatedAt.String(),
	}
}

func ToListCategoryResponse(categories []*entity.Category) []*DetailCategoryResponse {
	var res []*DetailCategoryResponse
	for _, category := range categories {
		res = append(res, ToDetailCategoryResponse(category))
	}
	return res
}

type ListCategoryResponse struct {
	PaginationResponse PaginationResponse        `json:"pagination"`
	Categories         []*DetailCategoryResponse `json:"categories"`
}
