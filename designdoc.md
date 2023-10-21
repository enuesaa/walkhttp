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
walkin serve --type static --response-header "Content-Type: application/json"
walkin serve --type restapi
walkin serve --type graphql
```

## Config file format
only one config file with multiple rules.

長ったらしい config file はユースケースに合わないので、なんとかしたい。

```json
{
    "rules": [
        {
            "path": "/*",
            "behavior": "readFiles",
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
        },
        {
            "path": "/users-graph",
            "behavior": "graphFiles",
            "graphFilesDirectory": "./users",
        },
    ]
}
```
