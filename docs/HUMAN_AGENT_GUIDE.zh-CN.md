# 人类维护指南（简体中文）

> [!WARNING]
> GoAPI 和 GoLoader（goloder）目前尚未公开。

## 目的

本指南面向维护或扩展此模板的开发者。
可以把它看作 `.goapi/AGENT.md` 的人类版本说明。

## 仓库职责

- 提供一个 Go 服务进程，用于处理桥接过来的 Minecraft 事件
- 将功能逻辑按模块拆分到 `modules/` 目录
- 保持 `main.go` 作为稳定入口

## 基本流程

1. 复制此模板仓库
2. 在 `modules/` 下创建新功能目录
3. 在 `Register()` 中注册事件和命令
4. 在 `main.go` 增加一行模块注册
5. 构建并部署 `server` 二进制

```bash
go build -o server .
```

## 推荐模块结构

```text
modules/
└── myfeature/
    ├── register.go
    ├── command.go
    └── events.go
```

- `register.go`：仅放注册代码
- 其他文件：放具体实现

## main.go 规则

- 将 `main.go` 视为基础设施层
- 不要把运行时初始化拆到功能模块里
- 仅追加模块注册调用

## 修改检查清单

- 新模块可正常编译
- `main.go` 已添加注册行
- 没有引入无关基础设施改动
- `server` 二进制可成功生成

## 协作建议

- 倾向于小而清晰的模块边界
- 处理器逻辑保持可预测、可测试
- 对不直观行为在实现附近补充注释
