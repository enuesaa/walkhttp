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
walkin up --config walkin.yaml
```

## Config
```yaml
paths:
  - path: '/api/*'
    behavior: proxy
    proxyConfig:
      url: https://localhost:3002/aaa/
    responseHeaders:
      Content-Type: application/json
  - path: '/files/*'
    behavior: readLocalFiles
  - path: '/*'
    behavior: proxy
    proxyConfig:
      url: https://localhost:3001

onUp:
  frontend:
    - cd ./frontend
    - pnpm dev --port 3001
  backend:
    - cd ./backend
    - go run . --port 3002
```
