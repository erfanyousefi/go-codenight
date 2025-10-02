package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	FirstName string         `gorm:"size:100; not null;column:firstname" json:"firstname"`
	LastName  string         `gorm:"size:100; not null;column:lastname" json:"lastname"`
	Email     string         `gorm:"size:255; not null;column:email" json:"email"`
	Password  string         `gorm:"size:255; null;column:password" json:"password"`
	Mobile    string         `gorm:"size:50; null;column:mobile" json:"mobile"`
	Balance   float64        `gorm:"type:decimal(10,2); default:0;column:balance" json:"balance"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
