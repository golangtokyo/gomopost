# gomopost

メッセージを POST するサンプル featuring gomobile

## ビルド方法

### 事前準備

* `go` コマンドが使える状態になっていること
* `gomobile` コマンドが使える状態になっていること
* `gomoible init` が済んでいること
* Android Studio 3.0.1 がインストールされていること

### コマンドラインで apk をビルドする

* トップディレクトリににて `make` を実行します。
  * cli 版 `gomopost` がビルドされます。
  * 傍らで、`gomobile bind` による AAR 生成が行われます。

### Android Studio で apk をビルドする

* 上述の `make` コマンドで gomopost.aar を生成しておきます。
* `android` ディレクトリを Android Studio に読み込ませます。
* ビルドして APK を生成し、実機に食わせるなりエミュレータに食わせるなりします。

## アプリの説明

[chatserver](https://github.com/golangtokyo/chatserver) にメッセージをポストするアプリです。
サーバの立て方、デプロイの仕方などは ↑ のリポジトリを参照してください。

### メッセージを POST する (cli 編)

* `gomopost/cmd/gomopost` 以下で `go build` すると、chatserver にメッセージをポストするための cli アプリがビルドされます。
* アプリが取れるオプションは以下です (`gomopost --help` で同内容を表示できます)。

```
NAME:
   gomopost - simple http client for gomobile instruction

USAGE:
   gomopost [options] [arguments...]

VERSION:
   0.1

OPTIONS:
   -a value, --address value  specify server address (scheme and port required) (default: "http://localhost:8080")
   -n value, --name value     specify name (default: "anonymous")
   -m value, --message value  specify message (default: "hello")
   --help, -h                 show help
   --version, -v              print the version
```

### メッセージをPOST する (Android アプリ編)

TBD

## LICENSE

MIT

