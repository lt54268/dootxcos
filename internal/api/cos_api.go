package api

import (
	"fmt"
	"net/http"

	"dootxcos/internal/service"

	"github.com/gin-gonic/gin"
)

// UploadHandler godoc
// @Summary 上传文件到腾讯云 COS
// @Description 接收文件并上传到腾讯云 COS
// @Tags 文件操作
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /upload_cos [post]
func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件读取失败"})
		return
	}

	// 打开文件流
	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件打开失败"})
		return
	}
	defer fileData.Close()

	// 使用 CosUploader 上传文件流
	uploader := service.NewCosUploader()
	url, err := uploader.Upload(fileData, file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": fmt.Sprintf("上传失败: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "url": url})
}

// DownloadCosFileHandler 下载文件接口
func DownloadFileHandler(c *gin.Context) {
	objectName := c.Query("object_Name")
	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "object_name 参数缺失"})
		return
	}

	downloader := service.NewCosDownloader()
	data, err := downloader.Download(objectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件下载失败", "error": err.Error()})
		return
	}

	// 设置响应头并返回文件数据
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func DeleteFileHandler(c *gin.Context) {
	// 从请求中获取文件名
	objectName := c.Query("objectName")
	if objectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "objectName is required",
		})
		return
	}

	cosDeleter := service.NewCosDeleter()
	err := cosDeleter.Delete(objectName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "File deleted successfully",
	})
}
