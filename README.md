# 腾讯云OSS
运行：go run main.go

## 一、上传接口（POST）
http://127.0.0.1:3030/upload

Body：file

返回示例：
```
{
    "code": 200,
    "msg": "上传成功",
    "url": "https://cloud-xxx.oss-cn-xxx.aliyuncs.com/大模型测评报告.docx"
}
```

## 二、下载接口（GET）
http://127.0.0.1:3030/download

参数：objectName

返回示例：返回文件，浏览器自动跳转下载

## 三、删除接口（DELETE）
http://127.0.0.1:3030/delete

参数：objectName

返回示例：
```
{
    "code": 200,
    "msg": "File deleted successfully"
}
```

## 四、获取文件列表接口（GET）
http://127.0.0.1:3030/list

参数：无

返回示例：
```
{
    "code": 200,
    "data": [
        {
            "key": "10.14会议纪要.docx",
            "size": 14077,
            "last_modified": "2024-11-04 03:29:47"
        },
        {
            "key": "大模型测评报告.docx",
            "size": 16530,
            "last_modified": "2024-11-04 06:30:45"
        }
    ],
    "msg": "File list retrieved successfully"
}
```

### 说明：
阿里云OSS对象存储，上传同名文件会自动覆盖旧文件
