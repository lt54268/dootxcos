package config

import (
	"dootxcos/internal/model"
	"os"
)

// LoadConfig 从环境变量加载配置信息
func LoadCosConfig() *model.Config {
	return &model.Config{
		Port:         os.Getenv("PORT"),          // 从环境变量读取端口
		CosRegion:    os.Getenv("COS_REGION"),    // 从环境变量读取区域
		CosEndpoint:  os.Getenv("COS_ENDPOINT"),  // 从环境变量读取 Endpoint
		CosBucket:    os.Getenv("COS_BUCKET"),    // 从环境变量读取 Bucket
		CosSecretId:  os.Getenv("COS_SECRETID"),  // 从环境变量读取 AccessKeyId
		CosSecretKey: os.Getenv("COS_SECRETKEY"), // 从环境变量读取 AccessKeySecret
	}
}
