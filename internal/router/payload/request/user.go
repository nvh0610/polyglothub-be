package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" validate:"required"`
}
