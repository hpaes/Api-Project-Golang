package database

import (
	"errors"

	"github.com/hpaes/api-project-golang/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductRepository) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "desc" {
		sort = "asc"
	} else {
		sort = "desc"
	}

	if page != 0 && limit != 0 {
		err = p.db.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.db.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *ProductRepository) Update(product *entity.Product) error {
	productFound, err := p.FindById(product.ID.String())
	if err != nil {
		return errors.New("product not found")
	}

	p.db.Model(productFound).Updates(product)
	return nil
}

func (p *ProductRepository) Delete(id string) error {
	productToDelete, err := p.FindById(id)
	if err != nil {
		return errors.New("product not found")
	}

	p.db.Delete(productToDelete)
	return nil
}
