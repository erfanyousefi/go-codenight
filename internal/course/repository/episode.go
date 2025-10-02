package repository

import (
	"github.com/erfanyousefi/codenight/internal/course/domains"
	"gorm.io/gorm"
)

type EpisodeRepository struct {
	db *gorm.DB
}

func NewEpisodeRepository(db *gorm.DB) *EpisodeRepository {
	return &EpisodeRepository{db: db}
}
func (r *EpisodeRepository) Create(c *domains.Episode) error {
	return r.db.Create(c).Error
}
func (r *EpisodeRepository) GetAll() ([]domains.Episode, error) {
	var episodes []domains.Episode
	err := r.db.Find(&episodes).Error
	return episodes, err
}
func (r *EpisodeRepository) GetAllByChapterID(chapterId uint) ([]domains.Episode, error) {
	var episodes []domains.Episode
	err := r.db.Where("chapterId = ?", chapterId).Find(&episodes).Error
	return episodes, err
}
func (r *EpisodeRepository) DeleteByID(id uint) error {
	return r.db.Where("id = ?", id).Delete(&domains.Episode{}).Error
}
func (r *EpisodeRepository) Update(id uint, c *domains.Episode) error {
	return r.db.Where("id = ?", id).Updates(c).Error
}
