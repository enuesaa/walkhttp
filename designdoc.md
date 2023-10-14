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

### (c) only one config file with multiple rules.
```json
{
    "rules": [
        {
            "path": "/*",
            "behavior": "passThrough",
            "accessWithoutExtention": true, // like `/users/aaa.json` or `/users/aaa/`
            "accessWithoutTrailingSlash": true, // like `/users/aaa.json` or `/users/aaa`. if accessWithoutExtention is false, this also do not work.
            "responseHeaders": {
                "Content-Type": "application/json"
            },
            "responseStatus": "200"
        },
        {
            "path": "/users",
            "behavior": "listFiles",
            "listFilesDirectory": "./users",
            "accessWithoutTrailingSlash": true,
            // todo re-consider
            "listFilesBodyBase": {
                "items": "LIST_FILES_ITEMS",
                "page": 1
            },
        },
        {
            "path": "/users-graph",
            "behavior": "graphFiles",
            "graphFilesDirectory": "./users",
        },
    ]
}
```