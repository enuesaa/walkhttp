# designdoc
- instant proxy server for web api to transform request and response.
- BFFライク(簡単なパスの変換をできる)
- Web API ではリクエストボディが大きく呼ぶのが大変な時があるが、これを変数だとかテンプレートだとかで簡略化できる

## development plan
- [cli] add command `paths add`
- [cli] add command `paths list`
- [cli] add command `serve`
- [cli] add command `paths transform-request --path /aa`
- [cli] add command `paths transform-response --path /bb`
- [http] use go-plugin to customize request header dynamically. this is useful for auth headers.
- [db] save to sqlite3
