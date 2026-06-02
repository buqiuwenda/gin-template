# gin-template

基于 **Gin** 的 Web API 服务脚手架：使用 **Cobra** 启动、**Google Wire** 依赖注入、**Protobuf** 定义接口与参数，业务按分层架构组织。

## 技术栈

| 能力 | 选型 |
|------|------|
| HTTP 框架 | [gin-gonic/gin](https://github.com/gin-gonic/gin) |
| CLI 启动 | [spf13/cobra](https://github.com/spf13/cobra) v1.10.2 |
| 依赖注入 | [google/wire](https://github.com/google/wire) |
| API 契约 | Protocol Buffers（`protoc` + `protoc-gen-go`） |
| 配置 | YAML（`gopkg.in/yaml.v3`） |

## 分层说明

```
transport (Gin Handler) → application (用例) → domain (实体/仓储接口) → data (仓储实现)
         ↑
    api/*.proto（请求/响应 DTO）
```

| 层级 | 目录 | 职责 |
|------|------|------|
| 契约层 | `api/` | `.proto` 定义接口参数与 DTO；`api/gen/go` 为生成代码 |
| 入口层 | `cmd/` | Cobra 命令：`web` 启动 HTTP，`job` 后台任务 |
| 传输层 | `internal/transport/http` | Gin 路由、绑定 JSON、调用 application |
| 应用层 | `internal/application` | 用例编排，不含 HTTP/DB 细节 |
| 领域层 | `internal/domain` | 实体、值对象、仓储 **interface** |
| 基础设施层 | `internal/data` | 仓储实现、MySQL/Redis 等 |
| 服务层 | `internal/server` | Gin Engine、路由注册、Server 生命周期 |
| 横切 | `internal/middleware`、`internal/meta`、`internal/config` | 中间件、统一响应、配置 |

## 目录结构

```
gin-template/
├── api/                              # API 契约（Protobuf）
│   ├── v1/
│   │   └── user/
│   │       └── user.proto            # 用户相关请求/响应定义
│   └── gen/go/                       # protoc 生成代码（make proto）
│       └── v1/user/user.pb.go
├── cmd/                              # 程序入口（Cobra）
│   ├── main.go
│   ├── root.go                       # 根命令
│   ├── web/
│   │   ├── http.go                   # web 子命令：启动 HTTP 服务
│   │   ├── wire.go                   # Wire 注入定义（wireinject）
│   │   └── wire_gen.go               # Wire 生成（make wire）
│   └── job/
│       └── once.go                   # job 子命令：一次性/定时任务
├── configs/
│   └── config.example.yaml           # 配置示例（复制为 config.yaml 使用）
├── internal/
│   ├── app/                          # 应用组装与运行（Run / 优雅退出）
│   │   ├── app.go
│   │   └── provider.go
│   ├── config/                       # 配置加载
│   │   ├── config.go
│   │   └── provider.go
│   ├── server/                       # HTTP Server（Gin + 路由）
│   │   ├── http.go
│   │   ├── router.go
│   │   └── provider.go
│   ├── transport/http/               # 传输层（Gin Handler）
│   │   ├── provider.go
│   │   └── v1/user/handler.go
│   ├── application/                  # 应用层（用例 Service）
│   │   ├── provider.go
│   │   └── user/service.go
│   ├── domain/                       # 领域层（实体 + 仓储接口）
│   │   └── user/
│   │       ├── user.go
│   │       └── repository.go
│   ├── data/                         # 基础设施（仓储实现、DB）
│   │   ├── data.go
│   │   ├── provider.go
│   │   └── user/repo.go
│   ├── middleware/
│   │   ├── jwt/jwt.go
│   │   └── recovery/recovery.go
│   ├── meta/                         # 统一响应、错误码
│   │   └── response.go
│   └── utils/                        # 项目内工具函数
│       └── string.go
├── pkg/                              # （可选）可被外部项目引用的公共库
├── third_party/                      # （可选）protoc 第三方 .proto
├── scripts/
│   ├── gen_proto.sh                  # 生成 api/gen/go
│   └── gen_wire.sh                   # 生成 cmd/web/wire_gen.go
├── deployments/                      # （可选）K8s / Helm 等部署清单
├── Dockerfile
├── Makefile
├── go.mod
└── README.md
```

## 快速开始

```bash
# 依赖
go mod tidy

# 生成 Wire（修改 Provider 后需重新执行）
make wire

# 生成 Protobuf（需安装 protoc、protoc-gen-go）
make proto

# 编译 & 启动 Web 服务
make run
# 或
go run ./cmd web -c configs/config.example.yaml
```

### 示例接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/healthz` | 健康检查 |
| POST | `/api/v1/users` | 创建用户（body 对齐 `CreateUserRequest`） |
| GET | `/api/v1/users/:id` | 查询用户 |

## 扩展新业务

1. 在 `api/v1/<biz>/` 增加 `.proto`，执行 `make proto`
2. 在 `internal/domain/<biz>/` 定义实体与 `Repository` 接口
3. 在 `internal/data/<biz>/` 实现仓储
4. 在 `internal/application/<biz>/` 编写 Service
5. 在 `internal/transport/http/v1/<biz>/` 编写 Handler 并注册路由
6. 在对应 `provider.go` 加入 `wire.NewSet`，执行 `make wire`

## 依赖注入（Wire）

注入入口：`cmd/web/wire.go` → 生成 `cmd/web/wire_gen.go`。

Provider 分布在各层 `provider.go`，由 `InitializeApp` 一次性组装：

`config` → `data` → `application` → `transport` → `server` → `app`

## License

See [LICENSE](LICENSE).
