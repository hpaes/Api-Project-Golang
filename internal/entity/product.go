package entity

import (
	"apis/pkg/entity"
	"errors"
	"time"
)

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrIDIsRequired       = errors.New("id is required")
	ErrNameIsRequired     = errors.New("name is required")
	ErrInvalidPrice       = errors.New("invalid price")
	ErrQuantityIsRequired = errors.New("quantity is required")
)

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewProduct(name, description string, price float64, quantity int) (*Product, error) {
	product := &Product{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   entity.GetTime(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if p.Quantity <= 0 {
		return ErrQuantityIsRequired
	}
	return nil
}
