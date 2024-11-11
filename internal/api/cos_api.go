package api

import (
	"fmt"
	"net/http"
	"os"

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
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "文件解析失败"})
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
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "object_Name 参数缺失"})
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
			"msg":  "object_Name 参数缺失",
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
		"msg":  "文件删除成功",
	})
}

// ListFilesHandler 处理获取文件列表请求
func ListFilesHandler(c *gin.Context) {
	lister := service.NewCosLister()
	files, err := lister.List()
	if err != nil {
		c.JSON(500, gin.H{
			"code":  500,
			"msg":   "Failed to retrieve file list from COS",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "文件列表获取成功",
		"data": files,
	})
}

// CopyFileHandler 处理文件拷贝请求
func CopyFileHandler(c *gin.Context) {
	// 获取请求参数
	srcBucket := c.DefaultQuery("srcBucket", os.Getenv("COS_BUCKET"))
	srcObject := c.DefaultQuery("srcObject", "")
	destBucket := c.DefaultQuery("destBucket", "")
	destObject := c.DefaultQuery("destObject", "")
	srcRegion := c.DefaultQuery("srcRegion", os.Getenv("COS_REGION"))
	destRegion := c.DefaultQuery("destRegion", os.Getenv("COS_REGION"))

	// 如果未提供 destBucket，则使用 srcBucket 作为默认值
	if destBucket == "" {
		destBucket = srcBucket
	}

	// 验证必要的参数
	if srcObject == "" || destObject == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少必需的参数: srcObject 或 destObject",
		})
		return
	}

	// 创建 CosCopier 实例并调用拷贝文件方法
	cosCopier := service.NewCosCopier()
	err := cosCopier.CopyFile(srcBucket, srcObject, destBucket, destObject, srcRegion, destRegion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("文件拷贝失败: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件拷贝成功",
	})
}

func MoveFileHandler(c *gin.Context) {
	// 获取请求参数
	srcBucket := c.DefaultQuery("srcBucket", os.Getenv("COS_BUCKET"))
	srcObject := c.DefaultQuery("srcObject", "")
	destBucket := c.DefaultQuery("destBucket", "")
	destObject := c.DefaultQuery("destObject", "")
	srcRegion := c.DefaultQuery("srcRegion", os.Getenv("COS_REGION"))
	destRegion := c.DefaultQuery("destRegion", os.Getenv("COS_REGION"))

	// 验证必要的参数
	if srcObject == "" || destObject == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少必需的参数: srcObject 或 destObject",
		})
		return
	}

	// 如果没有传 destBucket，默认使用 srcBucket
	if destBucket == "" {
		destBucket = srcBucket
	}

	// 创建 CosCopier 实例并调用 MoveFile 方法
	cosCopier := service.NewCosCopier()
	err := cosCopier.MoveFile(srcBucket, srcObject, destBucket, destObject, srcRegion, destRegion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("文件移动失败: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "文件移动成功",
	})
}
