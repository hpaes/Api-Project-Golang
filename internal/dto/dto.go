package dto

type CreateProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
