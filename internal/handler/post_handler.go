package handler

import (
	"context"
	"net/http"
	"social-backend/internal/dto"
	"social-backend/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{service: s}
}

// Create POST//
func (h *PostHandler) Create(c *gin.Context) {
	var req dto.CreatePostRequest

	// validasi request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil user dari middleware jwt
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	// Call service
	post, err := h.service.Create(ctx, userID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "data": post})
}

func (h *PostHandler) Feed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	posts, err := h.service.Feed(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *PostHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	post, err := h.service.GetById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (h *PostHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err := h.service.Delete(ctx, id, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not allowed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
