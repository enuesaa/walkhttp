# designdoc
## Features
- instant web server
- configure to serve files info with graohql or rest api

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある
- local api gateway みたいなイメージ

## Commands
```bash
walkin up
walkin up --config walkin.json
```

## Config
```json
{
    "paths": {
        "/*": {
            "behavior": "readLocalFiles"
        },
        "/*.json": {
            "behavior": "readLocalFiles",
            "responseHeaders": {
                "Content-Type": "application/json"
            },
        },
        "/files": {
            "behavior": "api"
        },
        "/something": {
            "behavior": "proxy",
            "url": "https://example.com"
        }
    }
}
```
