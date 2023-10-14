# designdoc
## Features
- instant web server
- serve files info with graohql
- serve files info with rest api

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある
- api gateway ライクにレスポンスとか変えられる

### Usecases
- API Mock
- HTML/CSS 確認

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
            // this is optional setting to append content-type header to response.
            "Content-Type": "application/json"
        },
    }
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
