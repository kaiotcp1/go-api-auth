package dtos

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email" example:"junior.dev@email.com"`
	Password string `json:"password" binding:"required,min=6" example:"123456"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email" example:"junior.dev@email.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginUserResponse struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Success bool   `json:"success" example:"true"`
}
