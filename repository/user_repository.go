package repository

// user repository with planet scale
type UserRepository interface {
	// Get all users
	GetAllUsers() error
	// Get user by email
	GetUserByEmail(email string) error
	// Get user by email and password
	GetUserByEmailAndPassword(email, password string) error
	// Create new user
	CreateUser(email, password string) error
	// Update user by email
	UpdateUserByEmail(email string) error
	// Delete user by email
	DeleteUserByEmail(email string) error
}
