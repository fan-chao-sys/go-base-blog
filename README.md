# Go Base Blog - 个人博客系统后端

基于 Go 语言、Gin 框架和 GORM 库开发的博客系统后端，实现博客文章的完整 CRUD 操作、用户认证和评论功能。

## 📋 目录

- [功能特性](#功能特性)
- [技术栈](#技术栈)
- [项目架构](#项目架构)
- [运行环境](#运行环境)
- [安装步骤](#安装步骤)
- [配置说明](#配置说明)
- [启动项目](#启动项目)
- [API 接口文档](#api-接口文档)
- [项目结构](#项目结构)

## ✨ 功能特性

- ✅ 用户注册与登录（密码 BCrypt 加密）
- ✅ JWT Token 认证与授权
- ✅ 文章 CRUD 操作（创建、读取、更新、删除）
- ✅ 评论功能（发表和查看评论）
- ✅ 用户权限控制（仅作者可修改/删除自己的文章）
- ✅ 统一错误处理与日志记录
- ✅ 数据库自动迁移

## 🛠 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| **Go** | 1.25.2 | 编程语言 |
| **Gin** | v1.11.0 | Web 框架 |
| **GORM** | v1.31.0 | ORM 库 |
| **MySQL** | 8.0+ | 数据库 |
| **JWT** | v5.3.0 | 身份认证 |
| **BCrypt** | - | 密码加密 |

## 🏗 项目架构

本项目采用经典的**三层架构**设计：

```
┌─────────────────────────────────────────────────┐
│                   Client Layer                   │
│            (HTTP Requests/Responses)             │
└─────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────┐
│              API Layer (Controller)              │
│  • 接收请求、参数验证、调用 Service              │
│  • 文件: api/userApi.go, postApi.go 等          │
└─────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────┐
│           Service Layer (Business Logic)         │
│  • 业务逻辑处理、数据操作、日志记录              │
│  • 文件: server/userService.go, postService.go  │
└─────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────┐
│              Model Layer (Data Model)            │
│  • 数据模型定义、数据库映射                      │
│  • 文件: model/user.go, post.go, comment.go    │
└─────────────────────────────────────────────────┘
                        ↓
┌─────────────────────────────────────────────────┐
│                 Database (MySQL)                 │
│            Tables: users, posts, comments        │
└─────────────────────────────────────────────────┘
```

### 核心模块说明

#### 1. **Initialize 模块** (`initialize/`)
- **database.go**: 数据库连接与配置读取
- **enter.go**: Service 层依赖注入
- **gin.go**: 路由配置与中间件设置

#### 2. **API 模块** (`api/`)
- **userApi.go**: 用户相关接口（注册、登录、查询）
- **postApi.go**: 文章相关接口（CRUD）
- **commentApi.go**: 评论相关接口
- **enter.go**: API 层统一入口，接收 Service 实例

#### 3. **Service 模块** (`server/`)
- **userService.go**: 用户业务逻辑（密码加密、验证）
- **postService.go**: 文章业务逻辑
- **commentService.go**: 评论业务逻辑
- **logService.go**: 日志记录服务
- **enter.go**: Service 层统一初始化

#### 4. **Model 模块** (`model/`)
- **user.go**: 用户模型（User 结构体）
- **post.go**: 文章模型（Post 结构体）
- **comment.go**: 评论模型（Comment 结构体）
- **log.go**: 日志模型（Log 结构体）
- **response.go**: 统一响应格式

#### 5. **Middleware 模块** (`middleware/`)
- **token.go**: JWT Token 生成与验证中间件

#### 6. **Utils 模块** (`utils/`)
- **logger.go**: 日志工具（LogInfo, LogError）

## 💻 运行环境

### 必需环境

- **Go**: 1.25.2 或更高版本
- **MySQL**: 8.0 或更高版本
- **Git**: 用于克隆项目（可选）

### 推荐开发工具

- GoLand / VSCode
- Postman / Apifox（API 测试）
- MySQL Workbench / Navicat（数据库管理）

## 📦 安装步骤

### 1. 克隆项目（或下载源码）

```bash
git clone <repository-url>
cd go-base-blog
```

### 2. 安装 Go 依赖

```bash
go mod download
```

如果遇到网络问题，可配置代理：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go mod download
```

### 3. 创建 MySQL 数据库

```sql
CREATE DATABASE go_gorm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. 配置数据库连接

编辑 `config/db.toml` 文件：

```toml
[mysql]
host = "127.0.0.1"
port = "3306"
name = "root"           # MySQL 用户名
password = "root"       # MySQL 密码
charset = "utf8mb4"
database = "go_gorm"    # 数据库名

[server]
port = 8080            # 服务器端口
timeout = 30
```

## 🚀 启动项目

### 方式 1：直接运行

```bash
go run main.go
```

### 方式 2：编译后运行

```bash
# 编译
go build -o go-base-blog.exe main.go

# 运行
./go-base-blog.exe
```

### 启动成功标志

当看到以下输出时，表示启动成功：

```
🚀 启动用户管理系统...
2025/10/27 16:00:00 [INFO] 2025-10-27 16:00:00 - 数据库地址-ip:127.0.0.1, 端口=3306
2025/10/27 16:00:00 [INFO] 2025-10-27 16:00:00 - 服务器端口: 8080
2025/10/27 16:00:00 [INFO] 2025-10-27 16:00:00 - 数据库mysql连接组装地址: ...
2025/10/27 16:00:00 [INFO] 2025-10-27 16:00:00 - <<<<<<<<<<<<<<<<<<<<<<<<<< Gin 初始化
[GIN-debug] Listening and serving HTTP on :8080
```

服务器默认运行在 `http://localhost:8080`

## 📡 API 接口文档

### 公开接口（/public）

无需认证即可访问

| 方法 | 路径 | 说明 | 请求体示例 |
|------|------|------|-----------|
| POST | `/public/register` | 用户注册 | `{"username":"user1","password":"123456","email":"user@email.com"}` |
| POST | `/public/login` | 用户登录 | `{"username":"user1","password":"123456"}` |
| GET | `/public/getUser?uid=1` | 获取用户信息 | - |
| GET | `/public/getPostList` | 获取文章列表 | - |
| GET | `/public/getPost?pid=1` | 获取文章详情 | - |
| GET | `/public/getComList?pid=1` | 获取文章评论列表 | - |

### 认证接口（/private）

需要在 Header 中携带 Token：`Authorization: Bearer <token>`

| 方法 | 路径 | 说明 | 请求体示例 |
|------|------|------|-----------|
| POST | `/private/createPost` | 创建文章 | `{"title":"标题","content":"内容","author":1,"userid":1}` |
| POST | `/private/createCom` | 创建评论 | `{"content":"评论内容","postid":1,"userid":1}` |
| PUT | `/private/upPost` | 更新文章（仅作者） | `{"id":1,"title":"新标题","content":"新内容"}` |
| DELETE | `/private/delPost?pid=1` | 删除文章（仅作者） | - |

### 响应格式

#### 成功响应

```json
{
  "code": 200,
  "message": "操作成功",
  "data": { ... }
}
```

#### 失败响应

```json
{
  "code": 500,
  "message": "错误信息"
}
```

## 📁 项目结构

```
go-base-blog/
├── api/                      # API 控制器层
│   ├── commentApi.go        # 评论 API
│   ├── enter.go             # API 统一入口
│   ├── postApi.go           # 文章 API
│   └── userApi.go           # 用户 API
├── config/                   # 配置文件
│   └── db.toml              # 数据库配置
├── initialize/               # 初始化模块
│   ├── database.go          # 数据库初始化
│   ├── enter.go             # Service 依赖注入
│   └── gin.go               # 路由与中间件配置
├── middleware/               # 中间件
│   └── token.go             # JWT 认证中间件
├── model/                    # 数据模型层
│   ├── comment.go           # 评论模型
│   ├── log.go               # 日志模型
│   ├── post.go              # 文章模型
│   ├── response.go          # 统一响应模型
│   └── user.go              # 用户模型
├── server/                   # 业务逻辑层
│   ├── commentService.go    # 评论业务逻辑
│   ├── enter.go             # Service 统一初始化
│   ├── logService.go        # 日志服务
│   ├── postService.go       # 文章业务逻辑
│   └── userService.go       # 用户业务逻辑
├── utils/                    # 工具类
│   └── logger.go            # 日志工具
├── go.mod                    # Go 模块依赖
├── go.sum                    # 依赖版本锁定
├── main.go                   # 程序入口
└── README.md                 # 项目说明文档
```

## 🔧 数据库表结构

项目使用 GORM 的 AutoMigrate 功能自动创建和更新表结构。

### Users 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |
| deleted_at | datetime | 软删除时间 |
| user_name | varchar(20) | 用户名，唯一 |
| pass_word | varchar(255) | 加密后的密码 |
| email | varchar(255) | 邮箱，唯一 |

### Posts 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |
| deleted_at | datetime | 软删除时间 |
| title | varchar(255) | 文章标题 |
| content | text | 文章内容 |
| author | uint | 作者 ID |
| user_id | uint | 用户 ID（外键） |

### Comments 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |
| deleted_at | datetime | 软删除时间 |
| content | text | 评论内容 |
| post_id | uint | 文章 ID（外键） |
| user_id | uint | 用户 ID（外键） |

## 📝 开发规范

### 代码风格

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 函数和变量命名采用驼峰命名法

### 错误处理

- 统一使用 `model.FailWithMessage()` 返回错误
- 关键操作记录到 Log 表
- 使用 `utils.LogInfo()` 和 `utils.LogError()` 记录日志

### 安全建议

1. **生产环境修改配置**：
   - 修改数据库密码
   - 设置环境变量 `JWT_SECRET`
   - 启用 HTTPS

2. **数据验证**：
   - 所有输入数据已添加 binding 验证
   - 密码使用 BCrypt 加密

3. **权限控制**：
   - `/private/*` 路由需要 JWT 认证
   - 文章更新/删除需验证作者身份

## 🐛 常见问题

### 1. 数据库连接失败

**错误**: `数据库连接失败: Error 1045: Access denied`

**解决**: 检查 `config/db.toml` 中的用户名和密码是否正确

### 2. 端口被占用

**错误**: `bind: address already in use`

**解决**: 
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Linux/Mac
lsof -ti:8080 | xargs kill -9
```

### 3. 外键约束失败

**错误**: `Cannot add or update a child row: a foreign key constraint fails`

**解决**: 确保创建评论/文章时，用户 ID 和文章 ID 在数据库中真实存在

## 📄 许可证

本项目仅用于学习和研究目的。

## 👨‍💻 作者

Go Base Blog Development Team

---

**快速开始命令**：

```bash
# 1. 创建数据库
mysql -u root -p -e "CREATE DATABASE go_gorm CHARACTER SET utf8mb4;"

# 2. 配置 config/db.toml

# 3. 安装依赖
go mod download

# 4. 启动项目
go run main.go
```

**测试接口**：

```bash
# 注册用户
curl -X POST http://localhost:8080/public/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456","email":"test@email.com"}'

# 获取文章列表
curl http://localhost:8080/public/getPostList
```

🎉 **享受编程的乐趣！**

