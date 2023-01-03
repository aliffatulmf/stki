package web

import (
	"aliffatulmf/stki/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// --- START CORPUS HANDLER ---

func (h *Handler) FormPage(ctx *gin.Context) {
	ctx.HTML(200, "form.html", gin.H{})
}

func (h *Handler) PostForm(ctx *gin.Context) {
	var body model.Corpus

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err})
		return
	}

	if err := h.Service.Create(&body); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(301, "/form-post")
}

// --- STOP CORPUS HANDLER ---

// --- START TFIDF ---

func (h *Handler) Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func (h *Handler) FindDocuments(ctx *gin.Context) {
	var param struct {
		Document string `json:"document" form:"document"`
	}

	if err := ctx.ShouldBind(&param); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	articles, err := h.Service.Finds(param.Document)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(200, "index.html", gin.H{"articles": articles})
}

// --- START STOP ---
