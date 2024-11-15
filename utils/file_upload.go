package utils

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// 从请求中解析文件
func ParseFile(c *gin.Context, formKey string) (multipart.File, string, error) {
	file, header, err := c.Request.FormFile(formKey)
	if err != nil {
		return nil, "", err
	}
	return file, header.Filename, nil
}
