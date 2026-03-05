# GoAPILib - AI Agent Context

このファイルを最初に読むこと。
go-server-template はこのファイルのコンテキストに基づいて実装する。

## プロジェクト概要

Minecraft Paper/Spigot サーバーで Go を使ってプラグインロジックを書けるようにするシステム。
Java プラグイン (GoAPI.jar + GoLoader.jar) が橋渡しをし、
ユーザーは Go コードだけでゲームロジックを実装できる。

## このリポジトリの役割

go-server-template はユーザーがプラグインを書くテンプレート。
- main.go は変更しない (モジュール登録の1行追加のみ)
- modules/ 以下に機能ごとのパッケージを作る
- goapi-sdk を import して On/Command/RunLater を呼ぶ

## ディレクトリ構成

