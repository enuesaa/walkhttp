# walkhttp
A CLI tool to proxy HTTP requests for debugging.

[![ci](https://github.com/enuesaa/walkhttp/actions/workflows/ci.yaml/badge.svg)](https://github.com/enuesaa/walkhttp/actions/workflows/ci.yaml)

## Usage
```console
$ walkhttp -h
A CLI tool to proxy HTTP requests for debugging.

USAGE:
  walkhttp [flags] [arguments...]

FLAGS:
  --port value    Port (default: 3000)
  --origin value  Origin URL (default: https://example.com)
  --prompt        Start prompt and call a HTTP request (default: false)
  --help, -h      show help
  --version, -v   print the version

$ walkhttp --origin https://example.com
┌─────────────────────────────────────────────────────────────────
│ walkhttp
│
│ --origin https://example.com
│ --port   3000
│
│ Web console: http://localhost:3000/_
└─────────────────────────────────────────────────────────────────

 Open the web console and try `curl http://localhost:3000/`
```

## モチベーション
- api gateway ライクなもの
- リバースプロキシ的な立ち位置で、HTTP リクエスト/レスポンスを中継する
- リクエスト/レスポンスをロギングして可観測性を上げたい
- あくまで開発用のもの。手っ取り早くトレースしたいときに使える
- めっちゃ雑に実装している

## キャプチャ
![ctlweb](./docs/ctlweb.png)
![ctlweb-history](./docs/ctlweb-history.png)
