package repository

type UserRoleRepository interface {
	// Get user role by user id
	GetUserRoleByUserId(userId string) error
	// Get user role by user id and role
	GetUserRoleByUserIdAndRole(userId, role string) error
	// Create new user role
	CreateUserRole(userId, role string) error
	// Update user role by user id
	UpdateUserRoleByUserId(userId string) error
	// Delete user role by user id
	DeleteUserRoleByUserId(userId string) error
}
