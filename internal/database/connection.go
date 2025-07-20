package database

import (
    "jwt-auth-crud/internal/models"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.User{},
        &models.Product{},
    )
}
