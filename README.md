# walkhttp
A CLI tool to proxy HTTP request for rough observability.

## Usage
```console
$ walkhttp --origin https://example.com
┌─────────────────────────────────────────────────────────────────
│ walkhttp
│
│ Web console: http://localhost:3000/_
│ Origin URL: https://example.com
│
│ Try `curl http://localhost:3000/` and open web console.
└─────────────────────────────────────────────────────────────────
```

## モチベーション
- api gateway ライクなもの
- リバースプロキシ的な立ち位置で、HTTP リクエスト/レスポンスをロギングすることで、可観測性を上げたい
- あくまで開発用のもの。手っ取り早くトレースしたいときに使えるもの
