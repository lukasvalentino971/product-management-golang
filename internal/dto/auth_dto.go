package dto

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    Name     string `json:"name" validate:"required,min=2,max=50"`
    Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
    Token string      `json:"token"`
    User  interface{} `json:"user"`
}
