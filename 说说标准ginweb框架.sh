project/
├── cmd/
│   └── server/
│       └── main.go           # 程序入口
│
├── internal/
│   ├── api/                  # HTTP Handler（Controller层）
│   │   ├── user.go
│   │   └── router.go
│   │
│   ├── service/              # 业务逻辑层
│   │   └── user_service.go
│   │
│   ├── repo/                 # 数据库访问层 类似Spring中的Mapper
│   │   └── user_repo.go
│   │
│   ├── model/                # 数据模型（struct）
│   │   └── user.go
│   │
│   ├── middleware/           # 中间件
│   │   ├── auth.go
│   │   └── logger.go
│   │
│   ├── infra/                # 基础设施（DB、Redis、配置...）
│   │   ├── db.go
│   │   ├── redis.go
│   │   └── config.go
│   │
│   └── app/                  # 依赖注入（可选）
│       └── wire.go           # 用 google/wire 自动注入
│
├── pkg/                      # 通用工具函数 (非业务逻辑)
│   ├── response/
│   ├── logger/
│   └── errors/
│
├── configs/                  # 配置文件
│   └── config.yaml
│
├── migrations/               # 数据库迁移（migrate工具）
│
├── test/                     # 单元测试
│
├── go.mod
└── Dockerfile
