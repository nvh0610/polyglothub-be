package error

import (
	"learn/pkg/resp"
)

var MsgFlags = map[int]string{
	USER_NOT_FOUND:            "User not found",
	WRONG_PASSWORD:            "Wrong password",
	UPDATE_PASSWORD_FAILED:    "Update password failed",
	SUCCESS:                   "Success",
	INVALID_PARAMS:            "Invalid params",
	UNAUTHORIZED:              "Unauthorized",
	INTERNAL_SERVER:           "Internal server",
	FORBIDDEN:                 "Not allow to process action",
	USER_EXIST:                "User existed",
	CREATE_USER_FAILED:        "Create user failed",
	UPDATE_USER_FAILED:        "Update user failed",
	USER_NOT_ADMIN:            "User not admin",
	WRONG_OTP:                 "Wrong otp",
	CREATE_CATEGORY_FAILED:    "Create category failed",
	CATEGORY_NOT_FOUND:        "Category not found",
	UPDATE_CATEGORY_FAILED:    "Update category failed",
	CREATE_VOCABULARY_FAILED:  "Create vocabulary failed",
	CREATE_EXAMPLE_FAILED:     "Create example failed",
	VOCABULARY_NOT_FOUND:      "Vocabulary not found",
	VOCABULARY_ALREADY_EXISTS: "Vocabulary already exists",
	FLASHCARD_DAILY_NOT_FOUND: "Flashcard daily not found",
	WRONG_ANSWER:              "Wrong answer",
	FLASHCARD_LOG_EXIST:       "User has already answered this vocabulary today",
}

func InitErrMsg() {
	for key, value := range MsgFlags {
		resp.MsgFlags[key] = value
	}
}
