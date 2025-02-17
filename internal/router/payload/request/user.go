package request

type CreateUserRequest struct {
	Username string `json:"username" required:"true"`
	Role     string `json:"role" required:"true"`
	Password string `json:"password" required:"true"`
}
