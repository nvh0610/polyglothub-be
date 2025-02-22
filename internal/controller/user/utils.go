package user

import "learn/internal/entity"

func IsValidUserRole(role string) bool {
	return role == entity.USER_ROLE_USER
}

func IsValidRole(role string) bool {
	return role == entity.USER_ROLE_USER || role == entity.USER_ROLE_ADMIN || role == entity.USER_ROLE_SUPERADMIN
}

func IsValidAdminRole(role string) bool {
	return role == entity.USER_ROLE_ADMIN || role == entity.USER_ROLE_SUPERADMIN
}

func IsValidSuperAdminRole(role string) bool {
	return role == entity.USER_ROLE_SUPERADMIN
}

func IsValidDeleteUser(role string, userRole string) bool {
	switch {
	case role == entity.USER_ROLE_ADMIN && userRole == entity.USER_ROLE_USER:
		return true
	case role == entity.USER_ROLE_ADMIN && userRole == entity.USER_ROLE_SUPERADMIN:
		return false
	case role == entity.USER_ROLE_ADMIN && userRole == entity.USER_ROLE_ADMIN:
		return false
	case role == entity.USER_ROLE_USER:
		return false
	case role == entity.USER_ROLE_SUPERADMIN:
		return true
	default:
		return false
	}
}
