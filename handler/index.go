package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type H struct {
	Gin *gin.Engine
	DB  *gorm.DB
}

func NewHandler(gin *gin.Engine, db *gorm.DB) *H {
	return &H{
		Gin: gin,
		DB:  db,
	}
}

func (h *H) FormPage(ctx *gin.Context) {

}
