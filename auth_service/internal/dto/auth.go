package dto

type UserCreate struct {
	Username string
	Password string
}

type UserResponse struct {
	ID       uint
	Username string
	// Role     string
}
