# designdoc
- instant web server

## Feature Plan: Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある
- local api gateway みたいなイメージ

### Commands
```bash
walkin serve --proxy 'path=/*,url=https://example.com' --read-local-files 'path=/*'
```

### Config
```json
{
  "paths": {
    "/api/*": {
      "behavior": "proxy",
      "proxyConfig": {
        "url": "https://localhost:3002/aaa/"
      },
      "responseHeaders": {
        "Content-Type": "application/json"
      }
    },
    "/files/*": {
      "behavior": "readLocalFiles"
    },
    "/*": {
      "behavior": "proxy",
      "proxyConfig": {
        "url": "https://localhost:3001"
      }
    }
  }
}
```

### dev command
```bash
go run . up --proxy "path=/*,url=https://example.com" --read-local-files "path=/aaa/*"
```

## Another Feature Plan: Statemachine with http client.
- あらかじめフロー(statemachine)を決めておく
- フロー通りにアクションを進める
- アクションでできることは Http のエンドポイントを叩くこと
- 最後にレポートのような画面がある
- Step Functions のようなプロダクト向けではなく、あくまで開発用

### usecase
- API Test
- 必然的にファイルダウンロードもできなきゃダメか（ログ残すには）
- curl や postman といった http client でできることのみ
