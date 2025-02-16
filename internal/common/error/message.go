package error

import (
	"learn/pkg/resp"
)

var MsgFlags = map[int]string{
	USER_NOT_FOUND:         "User not found",
	WRONG_PASSWORD:         "Wrong password",
	UPDATE_PASSWORD_FAILED: "Update password failed",
	SUCCESS:                "Success",
	INVALID_PARAMS:         "Invalid params",
	UNAUTHORIZED:           "Unauthorized",
	INTERNAL_SERVER:        "Internal server",
	FORBIDDEN:              "Not allow to process action",
	EXISTING_USER:          "User existed",
	CREATE_USER_FAILED:     "Create user failed",
}

func InitErrMsg() {
	for key, value := range MsgFlags {
		resp.MsgFlags[key] = value
	}
}
