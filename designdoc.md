# designdoc
## Features
- instant web server
- configure to serve files info with graohql or rest api

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある
- api gateway ライクにレスポンスとか変えられる

## Commands
```bash
walkin serve --config ./config.json
```

## Config file format
### (a) only one config file.
```json
// like openapi
{
    "paths": {
        "/": {
            "Content-Type": "application/json"
        },
    },
    "default": {
        "respone": {
            "headers": {
                // this is optional setting to append content-type header to response.
                "Content-Type": "application/json"
            }
        }
    },
    "rules": [
        {
            "if": {
                "path": "/"
            }
        }
    ]
}
```

### (b) multiple config files exists like file-based routing
```json
// /aaa/config.json match route `/aaa/`
{
    "request": {
        "method": "POST"
    },
    "response": {
        "headers": {
            "Content-Type": "application/json"
        }
    }
}
```
