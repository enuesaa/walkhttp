# designdoc
- A CLI Tool to debug web api.
- 簡単なパスの変換だとかできる (BFFライクな)
- Web API ではリクエストボディが大きく呼ぶのが大変な時があるが、これを変数だとかテンプレートだとかで簡略化できる
- cache & queue

## development plan
- [http] invoke http url with command line
- [http] create http invoke template
- [http] save histories
- [web] lookup histories
- [http] create pool to wait invoke
- [mock] replay http reponse for mock
