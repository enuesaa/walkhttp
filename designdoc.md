# designdoc
## Features
- instant web server
- configure to serve files info with graohql or rest api
- 例えば markdown の frontmatter を parse して JSON 形式で serve したい

### Instant Web Server
- ローカルにあるファイルを serve できる
- 開発しているとウェブサーバーが欲しい時がある

## Commands
```bash
walkin serve --config ./config.json
```

## config
```json
{
    "rules": {
        "/*": {
            "behavior": "readfile",
            "accessWithoutExtention": true, // like `/users/aaa.json` or `/users/aaa/`
            "accessWithoutTrailingSlash": true, // like `/users/aaa.json` or `/users/aaa`. if accessWithoutExtention is false, this also do not work.
        },
        "/*.json": {
            "behavior": "readfile",
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

## Development plan
- とりあえず上記仕様で実装する
- おそらく汎用性がないので、様子を見て仕様変更する

