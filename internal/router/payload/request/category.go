package request

type CreateCategoryRequest struct {
	Name   string `json:"name" validate:"required"`
	UserId int    `json:"user_id"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
