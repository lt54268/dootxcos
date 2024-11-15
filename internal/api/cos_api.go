package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"dootxcos/internal/service"

	"github.com/gin-gonic/gin"
)

// UploadHandler 上传文件接口
// @Summary 上传文件
// @Description 处理文件上传请求
// @Tags 文件管理
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传的文件"
// @Success 200 {object} model.UploadResponse "上传成功"
// @Failure 400 {object} map[string]interface{} "文件解析失败"
// @Failure 400 {object} map[string]interface{} "文件打开失败"
// @Failure 400 {object} map[string]interface{} "上传失败"
// @Router /api/v1/upload [post]
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
	info, err := uploader.Upload(fileData, file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": fmt.Sprintf("上传失败: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "上传成功", "data": info})
}

// DownloadFileHandler 下载文件接口
// @Summary 下载文件
// @Description 处理文件下载请求
// @Tags 文件管理
// @Accept json
// @Produce octet-stream
// @Param object_Name query string true "文件对象名"
// @Success 200 {file} []byte "文件数据"
// @Failure 400 {object} map[string]interface{} "参数缺失或文件下载失败"
// @Router /api/v1/download [get]
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

	c.Data(http.StatusOK, "application/octet-stream", data)
}

// DeleteFileHandler 删除文件接口
// @Summary 删除文件
// @Description 处理文件删除请求
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param objectName query string true "文件对象名"
// @Success 200 {object} map[string]interface{} "文件删除成功"
// @Failure 400 {object} map[string]interface{} "参数缺失"
// @Failure 500 {object} map[string]interface{} "文件删除失败"
// @Router /api/v1/delete [delete]
func DeleteFileHandler(c *gin.Context) {
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

// ListFilesHandler 获取文件列表接口
// @Summary 获取文件列表
// @Description 获取指定目录下的文件列表，并支持分页查询
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param prefix query string false "文件前缀"
// @Param marker query string false "分页查询标记，继续上次查询的位置"
// @Param maxKeys query int false "每页返回的文件数，最大值为1000，默认为1000"
// @Success 200 {object} map[string]interface{} "文件列表获取成功"
// @Failure 400 {object} map[string]interface{} "Invalid maxKeys parameter"
// @Failure 500 {object} map[string]interface{} "获取文件列表失败"
// @Router /api/v1/list [get]
func ListFilesHandler(c *gin.Context) {
	// 获取请求参数 Prefix、Marker、maxKeys
	prefix := c.DefaultQuery("prefix", "") // 默认为 ""
	marker := c.DefaultQuery("marker", "")
	maxKeysParam := c.DefaultQuery("maxKeys", "1000")
	maxKeys, err := strconv.Atoi(maxKeysParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"msg":   "Invalid maxKeys parameter",
			"error": err.Error(),
		})
		return
	}

	lister := service.NewCosLister()
	files, nextMarker, err := lister.List(prefix, marker, maxKeys)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  500,
			"msg":   "Failed to retrieve file list from COS",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":       200,
		"msg":        "文件列表获取成功",
		"data":       files,
		"nextMarker": nextMarker,
	})
}

// CopyFileHandler 拷贝文件接口
// @Summary 拷贝文件
// @Description 处理文件拷贝请求，支持在不同存储桶之间拷贝文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param srcBucket query string false "源存储桶" default(os.Getenv("COS_BUCKET"))
// @Param srcObject query string true "源文件对象名"
// @Param destBucket query string false "目标存储桶" default(srcBucket)
// @Param destObject query string true "目标文件对象名"
// @Param srcRegion query string false "源存储桶区域" default(os.Getenv("COS_REGION"))
// @Param destRegion query string false "目标存储桶区域" default(os.Getenv("COS_REGION"))
// @Success 200 {object} map[string]interface{} "文件拷贝成功"
// @Failure 400 {object} map[string]interface{} "缺少必需的参数"
// @Failure 500 {object} map[string]interface{} "文件拷贝失败"
// @Router /api/v1/copy [post]
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

// MoveFileHandler 移动文件接口
// @Summary 移动文件
// @Description 处理文件移动请求，支持在同一存储桶或不同存储桶之间移动文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param srcBucket query string false "源存储桶" default(os.Getenv("COS_BUCKET"))
// @Param srcObject query string true "源文件对象名"
// @Param destBucket query string false "目标存储桶" default(srcBucket)
// @Param destObject query string true "目标文件对象名"
// @Param srcRegion query string false "源存储桶区域" default(os.Getenv("COS_REGION"))
// @Param destRegion query string false "目标存储桶区域" default(os.Getenv("COS_REGION"))
// @Success 200 {object} map[string]interface{} "文件移动成功"
// @Failure 400 {object} map[string]interface{} "缺少必需的参数"
// @Failure 500 {object} map[string]interface{} "文件移动失败"
// @Router /api/v1/move [post]
func MoveFileHandler(c *gin.Context) {
	// 获取请求参数
	srcBucket := c.DefaultQuery("srcBucket", os.Getenv("COS_BUCKET"))
	srcObject := c.DefaultQuery("srcObject", "")
	destBucket := c.DefaultQuery("destBucket", "")
	destObject := c.DefaultQuery("destObject", "")
	srcRegion := c.DefaultQuery("srcRegion", os.Getenv("COS_REGION"))
	destRegion := c.DefaultQuery("destRegion", os.Getenv("COS_REGION"))

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
