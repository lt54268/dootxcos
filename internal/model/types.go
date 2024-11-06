package model

import (
	"mime/multipart"
)

// Config 用于存储配置信息
type Config struct {
	Port         string
	CosRegion    string
	CosEndpoint  string
	CosBucket    string
	CosSecretId  string
	CosSecretKey string
}

// Uploader 定义上传接口
type Uploader interface {
	Upload(file multipart.File, objectName string) (string, error)
}

// FileInfo 包含文件基本信息
type FileInfo struct {
	Key          string `json:"key"`
	Size         int64  `json:"size"`
	LastModified string `json:"last_modified"`
}
