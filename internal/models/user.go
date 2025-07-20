package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
    Email     string         `json:"email" gorm:"not null;unique;type:varchar(255)"`
    Password  string         `json:"-" gorm:"not null;type:varchar(255)"`
    Name      string         `json:"name" gorm:"not null;type:varchar(50)"`
    Role      string         `json:"role" gorm:"type:enum('admin','user');default:'user'"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
    
    Products []Product `json:"products,omitempty" gorm:"foreignKey:UserID"`
}
