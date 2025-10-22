# 个人博客系统后端

基于 Go + Gin + GORM + MySQL 开发的个人博客系统后端。

## 功能特性

- ✅ 用户注册和登录
- ✅ JWT 认证授权
- ✅ 文章 CRUD 操作
- ✅ 评论功能
- ✅ 权限控制
- ✅ 错误处理
- ✅ 日志记录

## 技术栈

- Go 1.25+
- Gin Web 框架
- GORM ORM
- MySQL 数据库
- JWT 认证

## 项目结构

```
blog-system/
├── main.go                 # 应用入口
├── go.mod                  # 依赖管理
├── config/
│   └── database.go        # 数据库配置
├── models/
│   └── models.go          # 数据模型
├── handlers/
│   ├── auth.go            # 认证处理器
│   ├── posts.go           # 文章处理器
│   └── comments.go        # 评论处理器
├── middleware/
│   └── auth.go            # 认证中间件
├── routes/
│   └── routes.go          # 路由配置
└── utils/
    └── logger.go          # 日志工具
```

## 环境配置

1. 创建 MySQL 数据库：
```sql
CREATE DATABASE blog_system;
```

2. 设置环境变量：
```bash
export DB_DSN="root:your_password@tcp(localhost:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
export PORT="8080"
```

## 安装和运行

1. 克隆项目：
```bash
git clone <repository-url>
cd go_gin
```

2. 安装依赖：
```bash
go mod tidy
```

3. 运行应用：
```bash
go run main.go
```

## API 接口

### 认证接口
- POST `/api/v1/auth/register` - 用户注册
- POST `/api/v1/auth/login` - 用户登录

### 文章接口
- GET `/api/v1/posts` - 获取所有文章
- GET `/api/v1/posts/:id` - 获取单篇文章
- POST `/api/v1/posts` - 创建文章（需认证）
- PUT `/api/v1/posts/:id` - 更新文章（需认证且仅作者）
- DELETE `/api/v1/posts/:id` - 删除文章（需认证且仅作者）

### 评论接口
- GET `/api/v1/posts/:post_id/comments` - 获取文章评论
- POST `/api/v1/posts/:post_id/comments` - 创建评论（需认证）

## 部署说明

1. 构建应用：
```bash
go build -o blog-system main.go
```

2. 运行：
```bash
./blog-system
```
<code_end>

这个完整的个人博客系统后端具有以下功能和特点：

1. **完整的项目结构**：模块化设计，清晰的目录结构
2. **数据库模型**：用户、文章、评论三表关联，包含完整字段定义
3. **用户认证**：JWT token机制，密码加密存储
4. **文章管理**：完整的CRUD操作，支持权限控制
5. **评论系统**：支持评论创建和查询
6. **错误处理**：统一的错误响应格式
7. **日志记录**：系统运行日志和错误日志
8. **API路由**：RESTful风格接口设计

所有功能模块都已实现，可以直接运行和部署。
