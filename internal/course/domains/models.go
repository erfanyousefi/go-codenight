package domains

import (
	"os/user"

	"github.com/erfanyousefi/codenight/internal/category"
)

type Course struct {
	ID          uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string            `gorm:"type:varchar(200); not null" json:"title"`
	Slug        string            `gorm:"type:varchar(255);unique; not null" json:"slug"`
	Image       string            `gorm:"type:varchar(255); not null" json:"image"`
	Description string            `gorm:"type:text; not null" json:"description"`
	Content     string            `gorm:"type:text" json:"content"`
	CategoryID  uint              `gorm:"column:categoryId" json:"categoryId"`
	Category    category.Category `gorm:"foreignKey:categoryId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	TeacherID   uint              `gorm:"column:teacherId" json:"teacherId"`
	Teacher     user.User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"teacher"`
	Chapters    []Chapter         `gorm:"foreignKey:courseId" json:"chapters,omitempty"`
}

type Chapter struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255); not null" json:"title"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Priority    uint      `gorm:"default:0" json:"priority"`
	CourseID    uint      `gorm:"column:courseId" json:"courseId"`
	Course      Course    `gorm:"foreignKey:courseId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"course"`
	Episodes    []Episode `gorm:"foreignKey:chapterId" json:"episodes,omitempty"`
}
type Episode struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string  `gorm:"type:varchar(255) not null" json:"title"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Priority    uint    `gorm:"default:0" json:"priority"`
	ChapterID   uint    `gorm:"column:chapterId" json:"chapterId"`
	Chapter     Chapter `gorm:"foreignKey:chapterId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"chapter"`
}
