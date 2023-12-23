# designdoc
- instant proxy server for web api to transform request and response.
- BFFライク(簡単なパスの変換をできる)
- Web API ではリクエストボディが大きく呼ぶのが大変な時があるが、これを変数だとかテンプレートだとかで簡略化できる
- configure paths to proxy with web api url.

## development plan
- [http] invoke http url with command line
- [http] create http invoke template
- [http] save histories
- [web] lookup histories
- [http] create pool to wait invoke
- [mock] replay http reponse for mock
- [http] cache & queue