package router

import (
	"dootxcos/internal/api"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 Gin 路由
func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", api.UploadHandler)
	r.GET("/download", api.DownloadFileHandler)
	r.DELETE("/delete", api.DeleteFileHandler)
	// r.GET("/list", api.ListFilesHandler)
}
