# designdoc
- batch manager
- statemachine 風味に http request をしたりコマンドを実行できる
- config ファイルに定義した順番で実行する
- 途中で何か失敗したらそのログなどトレースできるようにしたい

## development plan
- [cli] add command `validate`
- [cli] add command `plan`
- [cli] add command `run --log trace.log`
- [cli] add command `trace`

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
