package course

import (
	"github.com/erfanyousefi/codenight/internal/course/domains"
	"github.com/erfanyousefi/codenight/internal/course/repository"
)

// Course
type courseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *courseService {
	return &courseService{repo: repo}
}
func (s *courseService) Create(c *domains.Course) error {
	return s.repo.Create(c)
}
func (s *courseService) GetAll() ([]domains.Course, error) {
	return s.repo.GetAll()
}
func (s *courseService) GetByID(id uint) (domains.Course, error) {
	return s.repo.GetByID(id)
}
func (s *courseService) GetBySlug(slug string) (*domains.Course, error) {
	return s.repo.GetBySlug(slug)
}
func (s *courseService) Delete(id uint) error {
	return s.repo.DeleteByID(id)
}
func (s *courseService) Update(id uint, c *domains.Course) error {
	return s.repo.Update(id, c)
}

// Chapter
type chapterService struct {
	repo *repository.ChapterRepository
}

func NewChapterService(repo *repository.ChapterRepository) *chapterService {
	return &chapterService{repo: repo}
}
func (s *chapterService) Create(c *domains.Chapter) error {
	return s.repo.Create(c)
}

func (s *chapterService) GetAll() ([]domains.Chapter, error) {
	return s.repo.GetAll()
}
func (s *chapterService) GetAllByCourseID(courseId uint) ([]domains.Chapter, error) {
	return s.repo.GetAllByCourseID(courseId)
}
func (s *chapterService) GetByID(id uint) (domains.Chapter, error) {
	return s.repo.GetByID(id)
}
func (s *chapterService) Delete(id uint) error {
	return s.repo.DeleteByID(id)
}
func (s *chapterService) Update(id uint, c *domains.Chapter) error {
	return s.repo.Update(id, c)
}

// Episode
type episodeService struct {
	repo *repository.EpisodeRepository
}

func NewEpisodeService(repo *repository.EpisodeRepository) *episodeService {
	return &episodeService{repo: repo}
}
func (s *episodeService) Create(c *domains.Episode) error {
	return s.repo.Create(c)
}
func (s *episodeService) GetAll() ([]domains.Episode, error) {
	return s.repo.GetAll()
}
func (s *episodeService) GetAllByChapterID(chapterId uint) ([]domains.Episode, error) {
	return s.repo.GetAllByChapterID(chapterId)
}
func (s *episodeService) Delete(id uint) error {
	return s.repo.DeleteByID(id)
}
func (s *episodeService) Update(id uint, c *domains.Episode) error {
	return s.repo.Update(id, c)
}
