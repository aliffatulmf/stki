package web

import (
	"github.com/gin-gonic/gin"
)

func WebSetup(repo *Repository) error {
	service := NewService(repo)
	handler := NewHandler(service)

	g := gin.Default()

	g.LoadHTMLGlob("web/html/*.html")

	g.GET("/", handler.Index)
	g.POST("/", handler.FindDocuments)

	g.GET("/form-post", handler.FormPage)
	g.POST("/form-post", handler.PostForm)

	if err := g.Run(":8080"); err != nil {
		return err
	}

	return nil
}
