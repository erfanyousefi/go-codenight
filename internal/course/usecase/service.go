package usecase

import "github.com/erfanyousefi/codenight/internal/course/domains"

type CourseService interface {
	GetAll() ([]domains.Course, error)
	GetByID(id uint) (domains.Course, error)
	GetBySlug(slug string) (domains.Course, error)
	Create(c *domains.Course) error
	Update(id uint, c *domains.Course) error
	Delete(id uint) error
}

type ChapterService interface {
	GetAll() ([]domains.Chapter, error)
	GetAllByCourseID(courseId uint) ([]domains.Chapter, error)
	GetByID(id uint) (domains.Chapter, error)
	Create(c *domains.Chapter) error
	Delete(id uint) error
	Update(id uint, c *domains.Chapter) error
}

type EpisodeService interface {
	GetAll() ([]domains.Episode, error)
	GetAllByChapterID(chapterId uint) ([]domains.Episode, error)
	GetByID(id uint) (domains.Episode, error)
	Create(c *domains.Episode) error
	Delete(id uint) error
	Update(id uint, c *domains.Episode) error
}
