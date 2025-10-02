package repository

import (
	"github.com/erfanyousefi/codenight/internal/course/domains"
	"gorm.io/gorm"
)

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}
func (r *ChapterRepository) Create(c *domains.Chapter) error {
	return r.db.Create(c).Error
}
func (r *ChapterRepository) GetAll() ([]domains.Chapter, error) {
	var chapters []domains.Chapter
	err := r.db.Find(&chapters).Error
	return chapters, err
}
func (r *ChapterRepository) GetByID(id uint) (domains.Chapter, error) {
	var chapter domains.Chapter
	err := r.db.Preload("Episodes").Where("id = ?", id).First(&chapter).Error
	return chapter, err
}
func (r *ChapterRepository) GetAllByCourseID(courseId uint) ([]domains.Chapter, error) {
	var chapters []domains.Chapter
	err := r.db.Preload("Episodes").Where("courseId = ?", courseId).Find(&chapters).Error
	return chapters, err
}
func (r *ChapterRepository) DeleteByID(id uint) error {
	return r.db.Where("id = ?", id).Delete(&domains.Chapter{}).Error
}
func (r *ChapterRepository) Update(id uint, c *domains.Chapter) error {
	return r.db.Where("id = ?", id).Updates(c).Error
}
