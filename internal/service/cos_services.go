package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosUploader struct {
	client *cos.Client
}

type CosDownloader struct {
	client *cos.Client
}

type CosDeleter struct {
	client *cos.Client
}

func NewCosUploader() *CosUploader {
	return &CosUploader{
		client: NewCosClient(),
	}
}

func NewCosDownloader() *CosDownloader {
	return &CosDownloader{
		client: NewCosClient(),
	}
}

func NewCosDeleter() *CosDeleter {
	return &CosDeleter{
		client: NewCosClient(),
	}
}

func NewCosClient() *cos.Client {
	// 创建一个通用的 COS 客户端
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", os.Getenv("COS_BUCKET"), os.Getenv("COS_REGION")))
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRETID"),
			SecretKey: os.Getenv("SECRETKEY"),
		},
	})

	return client
}

// Upload 上传文件到腾讯云 COS
func (u *CosUploader) Upload(fileData multipart.File, objectName string) (string, error) {
	// 上传文件流
	_, err := u.client.Object.Put(context.Background(), objectName, fileData, nil)
	if err != nil {
		return "", err
	}
	// 返回文件的远程 URL
	return fmt.Sprintf("https://%s.cos.%s.myqcloud.com/%s", os.Getenv("COS_BUCKET"), os.Getenv("COS_REGION"), objectName), nil
}

// Download 从 COS 下载文件
func (d *CosDownloader) Download(objectName string) ([]byte, error) {
	resp, err := d.client.Object.Get(context.Background(), objectName, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete 方法删除指定的对象
func (d *CosDeleter) Delete(objectName string) error {
	_, err := d.client.Object.Delete(context.Background(), objectName, nil)
	if err != nil {
		if cos.IsNotFoundError(err) {
			return fmt.Errorf("resource not found: %v", objectName)
		}
		// if e, ok := cos.IsCOSError(err); ok {
		// 	return fmt.Errorf("COS error - Code: %v, Message: %v, Resource: %v, RequestId: %v", e.Code, e.Message, e.Resource, e.RequestID)
		// }
		// return fmt.Errorf("delete object error: %v", err)
		if e, ok := cos.IsCOSError(err); ok {
			if e.Code == "AccessDenied" {
				return fmt.Errorf("access denied. Please check COS permissions for DeleteObject operation")
			}
			return fmt.Errorf("COS error - Code: %v, Message: %v, Resource: %v, RequestId: %v", e.Code, e.Message, e.Resource, e.RequestID)
		}
	}
	return nil
}
