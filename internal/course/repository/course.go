package repository

import (
	"github.com/erfanyousefi/codenight/internal/course/domains"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(c *domains.Course) error {
	return r.db.Create(c).Error
}
func (r *CourseRepository) GetAll() ([]domains.Course, error) {
	var courses []domains.Course
	err := r.db.Find(&courses).Error
	return courses, err
}
func (r *CourseRepository) GetByID(id uint) (domains.Course, error) {
	var course domains.Course
	err := r.db.Preload("Chapters.Episodes").Where("id = ?", id).First(&course).Error
	return course, err
}
func (r *CourseRepository) GetBySlug(slug string) (*domains.Course, error) {
	var course domains.Course
	err := r.db.Preload("Chapters.Episodes").Where("slug = ?", slug).First(&course).Error
	return &course, err
}
func (r *CourseRepository) DeleteByID(id uint) error {
	return r.db.Delete(&domains.Course{}, id).Error
}
func (r *CourseRepository) Update(id uint, c *domains.Course) error {
	return r.db.Where("id = ?", id).Updates(c).Error
}
