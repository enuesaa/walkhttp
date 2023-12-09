# designdoc
- instant web server
- frontend と backend を仲介するものにしたい
- 開発の途上で Web API を用意できないために mock server を立てる時があるが、それを容易にしたい
- あと feature flag だとかで api を実際に呼ぶか mock を返すか切り替えられるようにしたい

## Feature Plan: Instant Web Server
- ローカルにあるファイルを serve できる
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
