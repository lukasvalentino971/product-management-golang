package repositories

import (
    "jwt-auth-crud/internal/models"
    "gorm.io/gorm"
)

type ProductRepository interface {
    Create(product *models.Product) error
    GetAll(userID uint) ([]models.Product, error)
    GetByID(id, userID uint) (*models.Product, error)
    Update(id, userID uint, product *models.Product) error
    Delete(id, userID uint) error
    //GetAllProducts
    GetAllProducts() ([]models.Product, error)
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
    return r.db.Create(product).Error
}

func (r *productRepository) GetAll(userID uint) ([]models.Product, error) {
    var products []models.Product
    err := r.db.Where("user_id = ?", userID).Find(&products).Error
    return products, err
}

func (r *productRepository) GetByID(id, userID uint) (*models.Product, error) {
    var product models.Product
    err := r.db.
        Preload("User"). // <-- Load data User terkait
        Where("id = ? AND user_id = ?", id, userID).
        First(&product).
        Error

    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *productRepository) Update(id, userID uint, product *models.Product) error {
    return r.db.Model(&models.Product{}).
        Where("id = ? AND user_id = ?", id, userID).
        Updates(product).Error
}

func (r *productRepository) Delete(id, userID uint) error {
    return r.db.Where("id = ? AND user_id = ?", id, userID).
        Delete(&models.Product{}).Error
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
    var products []models.Product
    err := r.db.Preload("User").Find(&products).Error // <-- Tambahkan Preload
    return products, err
}
