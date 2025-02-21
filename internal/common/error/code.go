package error

const (
	SUCCESS = 200000

	INVALID_PARAMS = 400000
	WRONG_PASSWORD = 400001
	USER_EXIST     = 400002
	WRONG_OTP      = 400003

	UNAUTHORIZED = 401000

	FORBIDDEN = 403000

	USER_NOT_FOUND     = 404001
	USER_NOT_ADMIN     = 404002
	CATEGORY_NOT_FOUND = 404003

	INTERNAL_SERVER        = 500000
	UPDATE_PASSWORD_FAILED = 500001
	CREATE_USER_FAILED     = 500002
	UPDATE_USER_FAILED     = 500003
	CREATE_CATEGORY_FAILED = 500004
	UPDATE_CATEGORY_FAILED = 500005
)
