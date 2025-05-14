package user

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
	UserRepository UserRepositoryImpl
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &Handler{DB: db, UserRepository: UserRepositoryImpl{}}

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", h.getUsers)
		userGroup.POST("/", h.createUser)
	}
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.UserRepository.GetAll(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) createUser(c *gin.Context) {
	var u User

	// body 전체를 출력
	body, err := c.GetRawData()
	if err == nil {
		println("[createUser] request body:", string(body))
	}
	// 다시 바인딩을 위해 body를 복원
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existUser User
	existUser, err = h.UserRepository.GetUser(h.DB, u.Name)
	if err == nil && existUser.ID != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	

	if err := h.UserRepository.Add(h.DB, &u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, u)
} 
