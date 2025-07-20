package models

import (
    "time"
    "gorm.io/gorm"
)

type Product struct {
    ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
    Name        string         `json:"name" gorm:"not null;type:varchar(100)"`
    Price       float64        `json:"price" gorm:"not null;type:decimal(10,2)"`
    Description string         `json:"description" gorm:"type:text"`
    Image       string         `json:"image" gorm:"type:varchar(255)"`
    UserID      uint           `json:"user_id" gorm:"not null"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

    User User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}
