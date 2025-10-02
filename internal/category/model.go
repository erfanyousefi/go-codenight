package category

import "time"

type Category struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"type:varchar(100) not null;" json:"title" binding:"required"`
	Slug        string     `gorm:"type:varchar(150) unique not null" json:"slug" binding:"required"`
	Description string     `gorm:"type:text" json:"description"`
	Image       string     `gorm:"type:varchar(255)" json:"image"`
	ParentID    *uint      `gorm:"column:parentId" json:"parentId"`
	Children    []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	CreatedAt   time.Time  `gorm:"type:timestamp; default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp; default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"`
}
