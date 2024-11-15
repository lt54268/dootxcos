// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/copy": {
            "post": {
                "description": "处理文件拷贝请求，支持在不同存储桶之间拷贝文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "拷贝文件",
                "parameters": [
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_BUCKET\"",
                        "description": "源存储桶",
                        "name": "srcBucket",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "源文件对象名",
                        "name": "srcObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "srcBucket",
                        "description": "目标存储桶",
                        "name": "destBucket",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "目标文件对象名",
                        "name": "destObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_REGION\"",
                        "description": "源存储桶区域",
                        "name": "srcRegion",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_REGION\"",
                        "description": "目标存储桶区域",
                        "name": "destRegion",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件拷贝成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "缺少必需的参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件拷贝失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/delete": {
            "delete": {
                "description": "处理文件删除请求",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件对象名",
                        "name": "objectName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件删除成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "参数缺失",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件删除失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/download": {
            "get": {
                "description": "处理文件下载请求",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "下载文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件对象名",
                        "name": "object_Name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件数据",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "参数缺失或文件下载失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/list": {
            "get": {
                "description": "获取指定目录下的文件列表，并支持分页查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "获取文件列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文件前缀",
                        "name": "prefix",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "分页查询标记，继续上次查询的位置",
                        "name": "marker",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页返回的文件数，最大值为1000，默认为1000",
                        "name": "maxKeys",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件列表获取成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid maxKeys parameter",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "获取文件列表失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/move": {
            "post": {
                "description": "处理文件移动请求，支持在同一存储桶或不同存储桶之间移动文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "移动文件",
                "parameters": [
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_BUCKET\"",
                        "description": "源存储桶",
                        "name": "srcBucket",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "源文件对象名",
                        "name": "srcObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "srcBucket",
                        "description": "目标存储桶",
                        "name": "destBucket",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "目标文件对象名",
                        "name": "destObject",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_REGION\"",
                        "description": "源存储桶区域",
                        "name": "srcRegion",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "os.Getenv(\"COS_REGION\"",
                        "description": "目标存储桶区域",
                        "name": "destRegion",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "文件移动成功",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "缺少必需的参数",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "文件移动失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/upload": {
            "post": {
                "description": "处理文件上传请求",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件管理"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "上传成功",
                        "schema": {
                            "$ref": "#/definitions/model.UploadResponse"
                        }
                    },
                    "400": {
                        "description": "上传失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UploadResponse": {
            "type": "object",
            "properties": {
                "content-length": {
                    "type": "integer"
                },
                "etag": {
                    "type": "string"
                },
                "last-modified": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
