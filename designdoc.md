# designdoc
- instant proxy server for web api to transform request and response.
- BFFライク(簡単なパスの変換をできる)
- Web API ではリクエストボディが大きく呼ぶのが大変な時があるが、これを変数だとかテンプレートだとかで簡略化できる

## development plan
- [cli] add command `new` ... configure url params, headers, body and finally save as template.
- [cli] add command `list` ... list templates. 
- [cli] add command `invoke <name> --aa bb`
- [cli] add command `transform-request <name> --set-header Content-Type 'application/json' --remove-header Aaa --set-header Accept $.var.aa`

- [cli] add command `serve`
- [db] save to sqlite3
- [http] use go-plugin to customize request header dynamically. this is useful for auth headers.

## how to call these web apis... 
### openai api
- https://qiita.com/yousan/items/3d53ee3922dffa191987
```console
$ curl https://api.openai.com/v1/chat/completions \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer ${API_KEY}' \
-d '{"model":"gpt-3.5-turbo","messages":[{"role":"user","content":"こんにちは"}]}'
```

### aws api
curl has flag to generate aws sigv4 headers.
- https://devops-blog.virtualtech.jp/entry/20230727/1690425829
- https://brandonstrohmeyer.medium.com/using-curl-to-call-aws-secrets-manager-api-198dbfc891e1

### contentful api
append query param named `access_token` 
- https://qiita.com/higebobo/items/877b1d2315c3be6b9636#contentful
