# go-server-template



Template for writing Minecraft plugin logic in Go.

## Languages

- [English](#english)
- [日本語](#日本語)
- [한국어](#한국어)
- [中文](#中文)

## Human Docs

- [Human Guide (English)](docs/HUMAN_AGENT_GUIDE.en.md)
- [Human Guide (日本語)](docs/HUMAN_AGENT_GUIDE.ja.md)
- [Human Guide (한국어)](docs/HUMAN_AGENT_GUIDE.ko.md)
- [Human Guide (中文)](docs/HUMAN_AGENT_GUIDE.zh-CN.md)

## English

Template repository for building a Go-based server module that runs behind the GoAPI/GoLoader bridge.

### Setup

1. Click `Use this template` to create your own repository.
2. Create your module by following `modules/example/`.
3. Add your module registration line in `main.go`.
4. Build the binary.

### Build

```bash
go build -o server .
# Place the "server" binary at plugins/GoLoader/server
```

### Add a module

```text
modules/
└── myfeature/
    ├── register.go   <- write goapi.On / goapi.Command registrations
    └── *.go          <- handler implementations
```

Add one line in `main.go`:

```go
myfeature.Register()
```

## 日本語

GoAPI/GoLoader ブリッジ上で動く Go 製サーバーモジュールのテンプレートです。

### セットアップ

1. `Use this template` でこのテンプレートを複製する
2. `modules/example/` を参考にモジュールを作成する
3. `main.go` にモジュール登録を1行追加する
4. バイナリをビルドする

### ビルド

```bash
go build -o server .
# "server" バイナリを plugins/GoLoader/server に配置
```

### モジュール追加

```text
modules/
└── myfeature/
    ├── register.go   <- goapi.On / goapi.Command の登録を書く
    └── *.go          <- ハンドラ実装
```

`main.go` に1行追加:

```go
myfeature.Register()
```

## 한국어

GoAPI/GoLoader 브리지 뒤에서 동작하는 Go 서버 모듈 템플릿입니다.

### 설정

1. `Use this template`로 저장소를 복제합니다.
2. `modules/example/`를 참고해 새 모듈을 만듭니다.
3. `main.go`에 모듈 등록 한 줄을 추가합니다.
4. 바이너리를 빌드합니다.

### 빌드

```bash
go build -o server .
# "server" 바이너리를 plugins/GoLoader/server 에 배치
```

### 모듈 추가

```text
modules/
└── myfeature/
    ├── register.go   <- goapi.On / goapi.Command 등록 작성
    └── *.go          <- 핸들러 구현
```

`main.go`에 한 줄 추가:

```go
myfeature.Register()
```

## 中文

这是一个用于编写 Go 服务器模块的模板，运行方式是通过 GoAPI/GoLoader 桥接。

### 初始化

1. 使用 `Use this template` 复制本仓库
2. 参考 `modules/example/` 创建自己的模块
3. 在 `main.go` 中添加模块注册代码
4. 构建二进制文件

### 构建

```bash
go build -o server .
# 将 "server" 二进制放到 plugins/GoLoader/server
```

### 添加模块

```text
modules/
└── myfeature/
    ├── register.go   <- 编写 goapi.On / goapi.Command 注册
    └── *.go          <- 处理逻辑实现
```

在 `main.go` 添加一行：

```go
myfeature.Register()
```
