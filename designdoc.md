# designdoc
## Features
- instant web server
- configure to serve files info with graohql or rest api
- もう少しエディター寄りの開発用ツールにしたい

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある

## Commands
```bash
walkin serve --config ./config.json
```

## Config file format
only one config file with multiple rules.

```json
{
    "rules": {
        "/*": {
            "behavior": "passthrough",
            "accessWithoutExtention": true, // like `/users/aaa.json` or `/users/aaa/`
            "accessWithoutTrailingSlash": true, // like `/users/aaa.json` or `/users/aaa`. if accessWithoutExtention is false, this also do not work.
        },
        "/*.json": {
            "behavior": "passthrough",
            "responseHeaders": {
                "Content-Type": "application/json"
            },
        },
        "/files": {
            "behavior": "api"
        },
        "/graph": {
            "behavior": "graphql"
        }
    }
}
```
