package category

import "gorm.io/gorm"

type CategoryService interface {
	CreateCategory(c *Category) error
	GetAllParents() ([]Category, error)
	GetAllCategories() ([]Category, error)
	GetByID(id uint) (*Category, error)
	UpdateCategory(id uint, data *Category) error
	DeleteCategory(id uint) error
}

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}
func (s *Service) CreateCategory(c *Category) error {
	return s.db.Create(&c).Error
}
func (s *Service) GetAllCategories() ([]Category, error) {
	var categories []Category
	err := s.db.Find(&categories).Error
	return categories, err
}
func (s *Service) GetAllParents() ([]Category, error) {
	var categories []Category
	err := s.db.Where("parentId IS NULL").Find(&categories).Error
	return categories, err
}
func (s *Service) GetByID(id uint) (Category, error) {
	var category Category
	err := s.db.Where("id = ?", id).First(&category).Error
	return category, err
}
func (s *Service) UpdateCategory(id uint, data *Category) error {
	return s.db.Model(&Category{}).Where("id=?", id).Updates(data).Error
}
func (s *Service) DeleteCategory(id uint) error {
	return s.db.Delete(&Category{}, id).Error
}
