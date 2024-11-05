package handler

import (
	"GO-JWT/internal/domain/repository"
	"GO-JWT/internal/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authRepo repository.AuthRepository
}

type LoginRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}

func NewAuthHandler(authRepo repository.AuthRepository) *AuthHandler {
	return &AuthHandler{
		authRepo: authRepo,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authRepo.GenerateToken(req.UserID, req.Role)
	if err != nil {
		logger.Info("token generation failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	logger.Info("token successfully made")

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
