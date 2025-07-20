package services

import (
	"jwt-auth-crud/internal/dto"
	"jwt-auth-crud/internal/models"
	"jwt-auth-crud/internal/repositories"
	"jwt-auth-crud/internal/utils"
)

type ProductService interface {
	CreateProduct(userID uint, req *dto.CreateProductRequest) (*models.Product, error)
	GetProducts(userID uint) ([]models.Product, error)
	GetProduct(id, userID uint) (*models.Product, error)
	UpdateProduct(id, userID uint, req *dto.UpdateProductRequest) error
	DeleteProduct(id, userID uint) error
	GetAllProducts() ([]models.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(userID uint, req *dto.CreateProductRequest) (*models.Product, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Image:       req.Image,
		UserID:      userID,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProducts(userID uint) ([]models.Product, error) {
	return s.productRepo.GetAll(userID)
}

func (s *productService) GetProduct(id, userID uint) (*models.Product, error) {
	return s.productRepo.GetByID(id, userID)
}

func (s *productService) UpdateProduct(id, userID uint, req *dto.UpdateProductRequest) error {
	if err := utils.ValidateStruct(req); err != nil {
		return err
	}

	product := &models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Image:       req.Image,
	}

	return s.productRepo.Update(id, userID, product)
}

func (s *productService) DeleteProduct(id, userID uint) error {
	return s.productRepo.Delete(id, userID)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
    // This method retrieves all products regardless of user ownership
    return s.productRepo.GetAllProducts()
}
