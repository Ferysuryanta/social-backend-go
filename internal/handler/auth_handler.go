package handler

import (
	"social-backend/internal/service"
	"social-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	ctx := c.Request.Context()
	user, err := h.service.Register(ctx, req.Email, req.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, _ := jwt.Generate(user.ID)

	c.JSON(200, gin.H{"token": token})
}
