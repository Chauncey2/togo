# Go 学习计划

这份计划面向有 Python 后端经验的开发者。目标不是从零学习编程，而是把已有后端能力迁移到 Go，并形成能提升竞争力、支撑副业或第二职业的作品集。

建议节奏：

- **14 天**：完成 Go 核心语法、并发、HTTP、测试与数据库实战
- **4 周**：完成一个可展示的后端项目，并补齐工程化能力
- **8-12 周**：形成求职/接单/独立产品都能复用的 Go 技术资产

## 你为什么适合学 Go

你已经有 7 年 Python 后端经验，数据库、API、缓存、异步任务、部署、排障这些后端基本功不需要重学。学习 Go 的重点是掌握 Go 的语言模型和工程习惯。

Go 能补足 Python 后端的几个竞争力短板：

- 更适合高并发、网络服务、基础设施、云原生工具
- 编译型语言，部署简单，单二进制交付友好
- 标准库强，工程一致性高，适合团队协作
- 在微服务、DevOps、网关、任务调度、数据处理工具中应用广
- 对副业友好，可以做 CLI 工具、API 服务、SaaS 后端、自动化平台

## 总体学习目标

- 掌握 Go 的核心语言机制：`struct`、`interface`、`error`、`context`、`goroutine`、`channel`
- 能独立完成一个可运行、可测试、可部署的 Go 后端服务
- 能解释 Python 与 Go 在错误处理、并发、类型系统、工程结构上的差异
- 能产出一个可用于简历、面试或副业展示的作品集项目

## 学习原则

- 少看长视频，多写代码。每天至少提交一次代码。
- 先用标准库，再引入框架。理解 `net/http` 后再用 Gin/Fiber。
- 先做完整闭环，再追求架构。一个能跑通的项目比十个半成品更有价值。
- 学 Go 不要照搬 Python。尤其不要过度封装、动态万能参数、异常式流程控制。
- 每周都要有可展示成果：代码、README、测试、Docker、压测记录、复盘文章。

---

## 第 0 阶段：环境与学习资源准备（半天）

### 目标

让本地环境能稳定运行 Go 项目，并准备好后续学习入口。

### 要做的事

- 安装 Go：<https://go.dev/dl/>
- 确认版本：`go version`
- 初始化模块：`go mod init example.com/to-go`
- 熟悉常用命令：`go run`、`go build`、`go test`、`go fmt`、`go vet`
- 准备一个 GitHub/Gitee 仓库，用于持续提交学习成果

### 核心链接

- Go 官方学习入口：<https://go.dev/learn/>
- Go 安装文档：<https://go.dev/doc/install>
- Go Modules Reference：<https://go.dev/ref/mod>
- 标准库文档：<https://pkg.go.dev/std>

### 输出物

- 本地能运行 `main.go`
- 项目已初始化 `go.mod`
- 远程仓库已创建并完成第一次提交

---

## 第 1 阶段：14 天 Go 后端冲刺计划

这 14 天是主线任务。你不需要把所有概念背熟，重点是每天有明确输出物。

### Day 1：语法速通与工程入口

学习内容：

- 包、模块、变量、常量、函数、控制流
- `go run`、`go build`、`go fmt`
- Go 项目的最小结构

学习链接：

- A Tour of Go：<https://go.dev/tour/>
- Go by Example：<https://gobyexample.com/>
- Effective Go：<https://go.dev/doc/effective_go>

输出物：

- 写一个 CLI 小程序：接收姓名、年龄、城市参数，输出格式化结果
- 写一段学习笔记：Go 的包和 Python module 有什么不同

---

### Day 2：基础类型、切片、映射与字符串

学习内容：

- `array`、`slice`、`map`、`string`
- 值传递与引用语义的区别
- `range` 的使用和常见坑

学习链接：

- Go Spec：<https://go.dev/ref/spec>
- Go by Example - Slices：<https://gobyexample.com/slices>
- Go by Example - Maps：<https://gobyexample.com/maps>

输出物：

- 写一个日志分析小工具：读取字符串列表，统计错误类型出现次数
- 对比 Python `list/dict` 与 Go `slice/map` 的差异

---

### Day 3：结构体、方法、接口

学习内容：

- `struct` 与方法
- 接口的隐式实现
- 组合优于继承
- 小接口原则

学习链接：

- Effective Go - Interfaces and methods：<https://go.dev/doc/effective_go>
- Go FAQ：<https://go.dev/doc/faq>
- Go by Example - Interfaces：<https://gobyexample.com/interfaces>

输出物：

- 实现一个 `UserRepository` 接口
- 提供内存版实现，并写一个简单调用示例

---

### Day 4：错误处理与资源释放

学习内容：

- `error` 返回值
- `defer`
- 错误包装：`fmt.Errorf("...: %w", err)`
- 不用异常控制业务流程

学习链接：

- Error handling and Go：<https://go.dev/blog/error-handling-and-go>
- errors 包文档：<https://pkg.go.dev/errors>
- fmt 包文档：<https://pkg.go.dev/fmt>

输出物：

- 给 Day 3 的 repository 增加错误返回
- 写一段对比：Python exception 与 Go error 的差异

---

### Day 5：并发基础

学习内容：

- `goroutine`
- `channel`
- `select`
- worker pool 基础

学习链接：

- Go Concurrency Patterns: Pipelines：<https://go.dev/blog/pipelines>
- Go by Example - Goroutines：<https://gobyexample.com/goroutines>
- Go by Example - Channels：<https://gobyexample.com/channels>
- Go by Example - Worker Pools：<https://gobyexample.com/worker-pools>

输出物：

- 实现一个并发任务执行器：输入 N 个任务，限制最多 3 个并发执行

---

### Day 6：`context` 与请求生命周期

学习内容：

- `context.Context`
- 超时、取消、父子上下文
- 在 HTTP、DB、RPC 调用中传递 context

学习链接：

- context 包文档：<https://pkg.go.dev/context>
- Go blog - Context：<https://go.dev/blog/context>

输出物：

- 给 Day 5 的任务执行器增加超时取消
- 总结：为什么后端服务必须重视 `context`

---

### Day 7：HTTP 服务与 JSON API

学习内容：

- `net/http`
- 请求解析与响应序列化
- REST API 基础
- 中间件模式

学习链接：

- net/http：<https://pkg.go.dev/net/http>
- encoding/json：<https://pkg.go.dev/encoding/json>
- Gin 文档（可选）：<https://gin-gonic.com/>

输出物：

- 用标准库实现一个 TODO API：创建、查询、更新、删除
- 暂时使用内存存储

---

### Day 8：测试与表驱动测试

学习内容：

- `testing`
- 表驱动测试
- `httptest`
- benchmark 入门

学习链接：

- testing：<https://pkg.go.dev/testing>
- httptest：<https://pkg.go.dev/net/http/httptest>
- Go by Example - Testing：<https://gobyexample.com/testing>

输出物：

- 为 TODO API 写 handler 测试
- 至少覆盖成功、参数错误、资源不存在 3 类场景

---

### Day 9：数据库访问

学习内容：

- `database/sql`
- 连接池
- SQL 超时与 context
- repository 分层

学习链接：

- database/sql：<https://pkg.go.dev/database/sql>
- PostgreSQL 驱动 pgx：<https://github.com/jackc/pgx>
- MySQL 驱动 go-sql-driver：<https://github.com/go-sql-driver/mysql>

输出物：

- 将 TODO API 从内存存储迁移到 PostgreSQL 或 MySQL
- 保留 repository 接口，替换具体实现

---

### Day 10：项目结构与分层

学习内容：

- `cmd/`
- `internal/`
- handler/service/repository 分层
- 配置、日志、路由初始化

学习链接：

- Go Modules Reference：<https://go.dev/ref/mod>
- Standard Go Project Layout（只做参考，不要照抄）：<https://github.com/golang-standards/project-layout>

建议结构：

```text
.
├── cmd/api
├── internal/config
├── internal/http
├── internal/service
├── internal/repository
└── internal/model
```

输出物：

- 重构 TODO API 项目结构
- 每层职责清晰，测试能继续运行

---

### Day 11：配置、日志与中间件

学习内容：

- 环境变量配置
- 结构化日志
- recover 中间件
- request id

学习链接：

- log/slog：<https://pkg.go.dev/log/slog>
- Zap（可选）：<https://github.com/uber-go/zap>
- Viper（可选）：<https://github.com/spf13/viper>

输出物：

- 给 API 加上统一日志、panic recover、请求耗时记录
- 配置通过环境变量读取

---

### Day 12：Docker 与本地运行环境

学习内容：

- 多阶段 Dockerfile
- Docker Compose
- app + db 一键启动

学习链接：

- Docker 文档：<https://docs.docker.com/>
- Compose 文档：<https://docs.docker.com/compose/>

输出物：

- 添加 Dockerfile
- 添加 `docker-compose.yml`
- README 中写清楚本地启动方式

---

### Day 13：并发安全、竞态检查、性能分析

学习内容：

- `sync.Mutex`、`sync.RWMutex`
- `sync/atomic`
- `go test -race`
- `pprof`

学习链接：

- sync：<https://pkg.go.dev/sync>
- sync/atomic：<https://pkg.go.dev/sync/atomic>
- pprof：<https://pkg.go.dev/net/http/pprof>
- Diagnostics：<https://go.dev/doc/diagnostics>

输出物：

- 对项目跑一次 race 检查
- 对一个接口做简单压测，并记录优化前后的指标

---

### Day 14：项目验收与复盘

学习内容：

- 复盘 Go 的核心差异
- 整理项目 README
- 准备作品集展示材料

学习链接：

- Go 官方文档总入口：<https://go.dev/doc/>
- 标准库总览：<https://pkg.go.dev/std>

输出物：

- 发布项目 v0.1
- 写一篇《Python 后端转 Go：14 天实践总结》
- README 包含功能说明、架构图、启动方式、测试方式、性能记录

---

## 第 2 阶段：4 周能力提升路线

14 天冲刺之后，继续用 2 周补齐生产能力。这里不建议继续刷语法，而是围绕一个项目持续增强。

### 第 3 周：生产级后端能力

重点：

- 认证鉴权：JWT、Session、API Key
- 数据库迁移：schema migration
- 缓存：Redis
- 限流：令牌桶或滑动窗口
- 统一错误码
- OpenAPI/Swagger 文档

资源：

- Redis Go client：<https://github.com/redis/go-redis>
- JWT Go：<https://github.com/golang-jwt/jwt>
- OpenAPI：<https://www.openapis.org/>
- Swagger UI：<https://swagger.io/tools/swagger-ui/>

输出物：

- 给 TODO API 增加用户系统
- 支持登录、鉴权、用户隔离数据
- 生成 API 文档

### 第 4 周：可观测、部署与稳定性

重点：

- 健康检查
- Prometheus 指标
- 日志追踪字段
- 优雅关闭
- CI 流程
- 部署到云服务器或容器平台

资源：

- Prometheus Go client：<https://github.com/prometheus/client_golang>
- GitHub Actions：<https://docs.github.com/actions>
- Go Graceful Shutdown 示例参考：<https://pkg.go.dev/net/http#Server.Shutdown>

输出物：

- 项目支持 `/healthz`、`/metrics`
- GitHub Actions 自动执行测试
- 服务可部署到一台云服务器或容器平台

---

## 第 3 阶段：竞争力与第二职业方向

学 Go 的最终目标不是“会语法”，而是把它变成职业杠杆。可以从下面 3 条路径中选一条主攻。

### 路线 A：求职竞争力增强

适合目标：

- 后端岗位升级
- 转向高并发、微服务、基础架构、云原生方向
- 简历增加 Go 技术栈

建议作品：

- 高并发短链服务
- 任务调度平台
- API 网关简化版
- 分布式限流服务

需要展示：

- QPS/延迟测试数据
- 并发安全说明
- 数据库与缓存设计
- 服务部署说明
- 关键技术取舍

### 路线 B：副业接单能力

适合目标：

- 接企业内部工具、API 服务、自动化工具
- 做轻量 SaaS 后端
- 给已有 Python 系统补 Go 高性能模块

建议作品：

- 企业内部审批/任务系统后端
- 数据同步工具
- Webhook 转发平台
- 文件处理/批量任务平台

需要展示：

- Docker 一键部署
- 清晰 API 文档
- 管理后台或示例前端
- 配置简单，维护成本低

### 路线 C：独立产品或第二职业

适合目标：

- 做自己的 SaaS
- 做开发者工具
- 做自动化平台
- 做内容 + 工具组合产品

建议作品：

- API 监控平台
- 日志告警小系统
- 个人知识库后端
- AI 工具调用网关
- 面向中小企业的轻量 CRM/工单系统

需要展示：

- 用户注册与登录
- 计费或额度系统
- 后台管理
- 部署文档
- 最小可用产品（MVP）演示

---

## 推荐作品集项目：Go TaskFlow

如果只做一个项目，建议做一个「任务与工作流后端系统」。它既贴合后端经验，又能覆盖 Go 的核心能力。

### 项目定位

一个支持用户、任务、标签、状态流转、Webhook 通知、定时任务的后端服务。

### 核心功能

- 用户注册、登录、JWT 鉴权
- 任务 CRUD
- 标签与筛选
- 任务状态流转
- 定时任务提醒
- Webhook 回调
- Redis 缓存热点查询
- PostgreSQL 持久化
- Docker Compose 一键启动
- Prometheus 指标

### 技术覆盖

- HTTP API
- `context`
- `database/sql` 或 `pgx`
- Redis
- 中间件
- 单元测试
- 集成测试
- Docker
- CI
- 观测指标

### 简历描述示例

可以在简历中这样描述：

```text
使用 Go 独立实现任务工作流后端系统，支持用户鉴权、任务状态流转、Webhook 通知、Redis 缓存、Prometheus 指标与 Docker 一键部署。项目采用 handler/service/repository 分层，核心接口覆盖单元测试与集成测试，并使用 context 管理请求超时和取消。
```

---

## Python 后端转 Go 的重点差异

### 1. 错误处理

Python 常用异常；Go 使用显式 `error` 返回。  
Go 的优势是调用链更清楚，代价是代码更啰嗦。

重点习惯：

- 不吞错误
- 返回错误时带上下文
- 使用 `errors.Is`、`errors.As` 判断错误类型
- 业务错误与系统错误区分

### 2. 并发模型

Python 后端常见并发方式包括多进程、线程池、asyncio、Celery。  
Go 的 `goroutine` 更轻量，但仍然需要控制数量、超时和退出。

重点习惯：

- 不无限制启动 goroutine
- 使用 `context` 控制生命周期
- 使用 worker pool 控制并发
- 用 `go test -race` 检查竞态

### 3. 类型系统

Python 动态类型灵活，但大型项目容易依赖约定。  
Go 类型系统更显式，适合长期维护。

重点习惯：

- 用明确结构体表达数据
- 小接口定义在使用方
- 少用 `any`，除非确实需要泛型/动态数据
- 不为了“复用”过早抽象

### 4. 工程风格

Go 更强调简单、直接、可读。  
不要把 Python/Django/FastAPI 的复杂分层完整搬过来。

重点习惯：

- 先用标准库
- 保持目录结构简单
- 文件名和包名表达清楚
- 可测试性优先于框架感

---

## 必读资源清单

### 官方资源

- Go Learn：<https://go.dev/learn/>
- A Tour of Go：<https://go.dev/tour/>
- Effective Go：<https://go.dev/doc/effective_go>
- Go Documentation：<https://go.dev/doc/>
- Go Spec：<https://go.dev/ref/spec>
- Standard library：<https://pkg.go.dev/std>
- Go Blog：<https://go.dev/blog/>

### 后端重点包

- `net/http`：<https://pkg.go.dev/net/http>
- `encoding/json`：<https://pkg.go.dev/encoding/json>
- `context`：<https://pkg.go.dev/context>
- `database/sql`：<https://pkg.go.dev/database/sql>
- `testing`：<https://pkg.go.dev/testing>
- `httptest`：<https://pkg.go.dev/net/http/httptest>
- `sync`：<https://pkg.go.dev/sync>
- `log/slog`：<https://pkg.go.dev/log/slog>

### 实战练习

- Go by Example：<https://gobyexample.com/>
- Exercism Go Track：<https://exercism.org/tracks/go>
- LeetCode：<https://leetcode.com/>

### 常用库

- Gin：<https://gin-gonic.com/>
- Fiber：<https://gofiber.io/>
- GORM：<https://gorm.io/>
- pgx：<https://github.com/jackc/pgx>
- go-redis：<https://github.com/redis/go-redis>
- Cobra：<https://github.com/spf13/cobra>
- Viper：<https://github.com/spf13/viper>
- Zap：<https://github.com/uber-go/zap>
- GolangCI-Lint：<https://golangci-lint.run/>

### 推荐书籍

- 《The Go Programming Language》
- 《100 Go Mistakes and How to Avoid Them》
- 《Concurrency in Go》

---

## 每周验收标准

### 第 1 周

- 能写基本 Go 程序
- 理解 `struct`、`interface`、`error`、`context`
- 能写简单并发程序
- 能用 `net/http` 写 API

### 第 2 周

- 完成一个带数据库的 API 服务
- 有基础单元测试
- 有清晰项目结构
- 能用 Docker 启动

### 第 3 周

- 增加用户鉴权、缓存、限流、错误码
- API 文档基本完整
- 服务具备基本生产形态

### 第 4 周

- 有 CI
- 有健康检查与监控指标
- 有性能测试记录
- README 能让别人独立运行项目

---

## 最小学习闭环

每天结束前完成这 5 件事：

1. 提交代码
2. 运行测试
3. 写 3-5 行学习笔记
4. 记录一个 Python 与 Go 的差异
5. 写明第二天要做什么

每周结束前完成这 4 件事：

1. 整理 README
2. 补一个测试
3. 写一个复盘小节
4. 把项目推送到远程仓库

## 不建议做的事

- 不要一开始就学太多框架
- 不要只看视频不写项目
- 不要先纠结完美目录结构
- 不要把 Go 写成 Java 或 Python
- 不要忽视测试、超时、日志、部署

## 推荐下一步

从当前项目开始，先完成这 3 个小目标：

1. 用 Go 写一个最小 HTTP API
2. 给 API 增加测试
3. 把它改造成一个可持续扩展的 TaskFlow 项目
