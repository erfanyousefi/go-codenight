package category

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/erfanyousefi/codenight/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) CreateCategory(ctx *gin.Context) {
	var category Category
	category.Title = ctx.PostForm("title")
	category.Slug = ctx.PostForm("slug")
	category.Description = ctx.PostForm("description")

	file, err := ctx.FormFile("image")
	if err == nil {
		var filename string
		ext := path.Ext(file.Filename)
		root, _ := os.Getwd()
		filename = fmt.Sprintf("%d%s", time.Now().UnixMilli(), ext)
		address := fmt.Sprintf("/uploads/categories/%s", filename)
		if err := ctx.SaveUploadedFile(file, path.Join(root, "uploads", "categories", filename)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		category.Image = address
	}
	if err := h.service.CreateCategory(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "category created successfully"})
}
func (h *Handler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
func (h *Handler) GetAllParents(ctx *gin.Context) {
	categories, err := h.service.GetAllParents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}
func (h *Handler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	category, err := h.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}
func (h *Handler) UpdateCategory(ctx *gin.Context) {
	var category Category
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	category.Title = ctx.PostForm("title")
	category.Slug = ctx.PostForm("slug")
	category.Description = ctx.PostForm("description")
	address, err := utils.SaveUploadFile(ctx, "image", "categories")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if address != "" {
		category.Image = address
	}
	if err := h.service.UpdateCategory(uint(id), &category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "category updated successfully"})
}
func (h *Handler) DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.service.DeleteCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
