# designdoc
## Features
- instant web server

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある
- local api gateway みたいなイメージ

## Commands
```bash
walkin up
walkin up --proxy 'path=/*,url=https://example.com' --read-local-files 'path=/*'
walkin up --config walkin.json
```

## Config
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
