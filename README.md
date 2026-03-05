# go-server-template

GoAPILib を使って Minecraft プラグインを Go で書くためのテンプレートです。

## セットアップ

1. このリポジトリを「Use this template」でコピーする
2. `modules/example/` を参考に自分のモジュールを作る
3. `main.go` にモジュールを登録する
4. ビルドする

## ビルド

```bash
go build -o server .
# -> server バイナリを plugins/GoLoader/server に配置
```

## モジュールの追加

```
modules/
└── myfeature/
    ├── register.go   <- goapi.On / goapi.Command をここに書く
    └── *.go          <- ハンドラ実装
```

main.go に1行追加:
```go
myfeature.Register()
```

## API リファレンス

- [goapi-sdk](https://github.com/mituba4154/goapi-sdk)
- [GoAPILib (Java)](https://github.com/mituba4154/GoAPI)
