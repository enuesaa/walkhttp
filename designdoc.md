# designdoc
- instant web server which deals with frontend and backend.
- web api の mock を簡単に作れるもの
- feature flag だとかで api を実際に呼ぶか mock を返すか切り替えられるようにしたい
- local api gateway みたいなイメージ

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
