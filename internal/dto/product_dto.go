package dto

type CreateProductRequest struct {
    Name        string  `json:"name" validate:"required,min=2,max=100"`
    Price       float64 `json:"price" validate:"required,min=0"`
    Description string  `json:"description"`
    Image       string  `json:"image"`
}

type UpdateProductRequest struct {
    Name        string  `json:"name" validate:"omitempty,min=2,max=100"`
    Price       float64 `json:"price" validate:"omitempty,min=0"`
    Description string  `json:"description"`
    Image       string  `json:"image"`
}

type ProductResponse struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
    Image       string  `json:"image"`
    UserID      uint    `json:"user_id"`
    CreatedAt   string  `json:"created_at"`
    UpdatedAt   string  `json:"updated_at"`
}
