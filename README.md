# go-tinyurl

用golang实现的短链接demo级项目。

## 使用方法

使用前请确保redis运行在`127.0.0.1:6379`，且密码为空。  
~~当然也可以修改 ./store/store_service.go 中连接redis相关的配置。~~

执行：

```bash
go mod download
go run main.go
```

## 接口说明



### 1. 创建短网址

- **URL**: `/create-short-url`

- **Method**: POST

- Request Body

  : JSON 格式，包含以下字段：

    - `long_url` (string): 必需，用户要缩短的长网址。
    - `user_id` (string): 必需，用户的唯一标识。

#### 示例请求

```http
POST /create-short-url HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "long_url": "https://example.com",
    "user_id": "12345"
}
```

#### Response

- **Status code**: 200 OK

```json
{
    "message": "short url created successfully",
    "short_url": "http://localhost:8080/abc123"
}
```

### 2. 重定向到原始网址

- **URL**: `/:shortUrl`

- **Method**: GET

- 路径参数

  :

    - `shortUrl` (string): 用户提供的短网址部分，用于查找原始网址。

#### 示例请求

```http
GET /abc123 HTTP/1.1
Host: localhost:8080
```

#### Response

- **Status code**: 302 Found (重定向)
- **Location**: 原始网址。例如：`https://example.com`