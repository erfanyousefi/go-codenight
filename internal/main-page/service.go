package mainPage

import (
	"github.com/erfanyousefi/codenight/internal/category"
	"gorm.io/gorm"
)

type MainPageType struct {
	Categories []category.Category `json:"categories"`
}
type MainPage interface {
	getCategories() ([]category.Category, error)
	getMainPageData() (MainPageType, error)
}
type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}
func (s *Service) getCategories() ([]category.Category, error) {
	var categories []category.Category
	err := s.db.Find(&categories).Error
	return categories, err
}
func (s *Service) getMainPageData() (any, error) {
	categories, err := s.getCategories()
	data := MainPageType{
		Categories: categories,
	}
	return data, err
}
