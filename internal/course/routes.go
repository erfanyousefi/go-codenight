package course

import (
	"github.com/erfanyousefi/codenight/config"
	"github.com/erfanyousefi/codenight/internal/course/repository"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	courseRepo := repository.NewCourseRepository(config.GORMDB)
	chapterRepo := repository.NewChapterRepository(config.GORMDB)
	episodeRepo := repository.NewEpisodeRepository(config.GORMDB)

	courseService := NewCourseService(courseRepo)
	chapterService := NewChapterService(chapterRepo)
	episodeService := NewEpisodeService(episodeRepo)
	handler := NewHandler(courseService, chapterService, episodeService)
	courseGroup := r.Group("/course")
	{
		courseGroup.POST("/", handler.CreateCourse)
		courseGroup.POST("/chapter", handler.CreateChapter)
		courseGroup.POST("/episode", handler.CreateEpisode)
		courseGroup.GET("/", handler.GetAllCourse)
		courseGroup.GET("/chapter", handler.GetAllChapters)
		courseGroup.GET("/episode", handler.GetAllEpisodes)
		courseGroup.GET("/chapter/:courseId", handler.GetChapterByCourseID)
		courseGroup.GET("/episode/:chapterId", handler.GetEpisodesByChapterID)
		courseGroup.GET("/:id", handler.GetCourseByID)
	}

}
