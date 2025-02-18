package user

import "learn/internal/entity"

func IsValidRole(role string) bool {
	return role == entity.USER_ROLE_ADMIN || role == entity.USER_ROLE_USER || role == entity.USER_ROLE_SUPERADMIN
}

func IsValidAdminRole(role string) bool {
	return role == entity.USER_ROLE_ADMIN || role == entity.USER_ROLE_SUPERADMIN
}

func IsValidSuperAdminRole(role string) bool {
	return role == entity.USER_ROLE_SUPERADMIN
}
