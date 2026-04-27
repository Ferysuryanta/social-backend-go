package handler

import (
	"context"
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
func (h *PostHandler) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err := h.service.Create(ctx, userID, req.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post created successfully"})
}

func (h *PostHandler) Feed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	posts, _ := h.service.Feed(ctx)
	c.JSON(200, posts)
}
