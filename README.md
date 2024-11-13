# 腾讯云OSS
运行：go run main.go

API文档：http://127.0.0.1:9090/swagger/index.html

## 一、上传接口（POST）
http://127.0.0.1:9090/api/v1/upload

Body：file

返回示例：
```
{
    "code": 200,
    "msg": "上传成功",
    "url": {
        "content-length": 10089,
        "etag": "\"55axxxxxxxxxxxxxxxxxxxxxxxxx\"",
        "last-modified": "2024-11-11T06:08:26Z"
    }
}
```

## 二、下载接口（GET）
http://127.0.0.1:9090/api/v1/download

参数：objectName

返回示例：返回文件，浏览器自动跳转下载

## 三、删除接口（DELETE）
http://127.0.0.1:9090/api/v1/delete

参数：objectName

返回示例：
```
{
    "code": 200,
    "msg": "文件删除成功"
}
```

## 四、获取文件列表接口（GET）
http://127.0.0.1:9090/api/v1/list

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
    "msg": "文件列表获取成功"
}
```

## 五、拷贝接口（POST）
http://127.0.0.1:9090/api/v1/copy

参数：srcBucket、srcObject、destBucket、destObject

返回示例：同一个桶不需要传destBucket参数
```
{
  "code": 200,
  "msg": "文件拷贝成功"
}
```

## 六、移动接口（POST）
http://127.0.0.1:9090/api/v1/move

参数：srcBucket、srcObject、destBucket、destObject

返回示例：同一个桶不需要传destBucket参数
```
{
  "code": 200,
  "msg": "文件移动成功"
}
```

### 说明：
腾讯云COS对象存储，上传同名文件会自动覆盖旧文件
