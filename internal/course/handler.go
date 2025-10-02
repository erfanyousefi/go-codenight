package course

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/erfanyousefi/codenight/internal/course/domains"
	"github.com/erfanyousefi/codenight/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	courSer *courseService
	chapSer *chapterService
	epoSer  *episodeService
}

func NewHandler(cos *courseService, chs *chapterService, eps *episodeService) *Handler {
	return &Handler{courSer: cos, chapSer: chs, epoSer: eps}
}

type CourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	CategoryID  uint   `json:"categoryId"`
	Slug        string `json:"slug"`
}

func (h *Handler) CreateCourse(ctx *gin.Context) {
	var course domains.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if course.Slug != "" {
		course, err := h.courSer.GetBySlug(course.Slug)
		if err != nil && err != gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}
		if course != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": "already exist course with this slug"})
			return
		}
	}
	address, err := utils.SaveUploadFile(ctx, "image", "courses")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if address != "" {
		course.Image = address
	}
	if err := h.courSer.Create(&course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "course created successfully"})
}
func (h *Handler) CreateChapter(ctx *gin.Context) {
	var chapter domains.Chapter
	if err := ctx.ShouldBindJSON(&chapter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.chapSer.Create(&chapter); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created chapter successfully"})
}
func (h *Handler) CreateEpisode(ctx *gin.Context) {
	var episode domains.Episode
	if err := ctx.ShouldBindJSON(&episode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.epoSer.Create(&episode); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created chapter successfully"})
}

func (h *Handler) GetAllCourse(ctx *gin.Context) {
	courses, err := h.courSer.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}
func (h *Handler) GetAllChapters(ctx *gin.Context) {
	courses, err := h.chapSer.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}
func (h *Handler) GetAllEpisodes(ctx *gin.Context) {
	courses, err := h.epoSer.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}
func (h *Handler) GetCourseByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	course, err := h.courSer.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)

}
func (h *Handler) GetCourseBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	course, err := h.courSer.GetBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}
func (h *Handler) GetChapterByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	course, err := h.chapSer.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}
func (h *Handler) GetChapterByCourseID(ctx *gin.Context) {
	idStr := ctx.Param("courseId")
	courseId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	chapters, err := h.chapSer.GetAllByCourseID(uint(courseId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chapters)
}
func (h *Handler) GetEpisodesByChapterID(ctx *gin.Context) {
	idStr := ctx.Param("chapterId")
	chapterId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	episodes, err := h.epoSer.GetAllByChapterID(uint(chapterId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, episodes)
}

func (h *Handler) DeleteCourse(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if err := h.courSer.Delete(uint(id)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (h *Handler) DeleteChapter(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if err := h.chapSer.Delete(uint(id)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func (h *Handler) DeleteEpisode(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		if err := h.epoSer.Delete(uint(id)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) UpdateCourse(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := h.courSer.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req domains.Course
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Title != "" {
		course.Title = req.Title
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	address, err := utils.SaveUploadFile(ctx, "image", "courses")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if address != "" {
		course.Image = address
	}
	if err := h.courSer.Update(course.ID, &course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "updated successfully"})
}
func (h *Handler) UpdateChapter(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapter, err := h.chapSer.GetByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req domains.Chapter
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Title != "" {
		chapter.Title = req.Title
	}
	if req.Description != "" {
		chapter.Description = req.Description
	}
	if req.Course.Slug != "" {
		chapter.Course.Slug = req.Course.Slug
	}
	if err := h.chapSer.Update(chapter.ID, &chapter); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "updated successfully"})
}
