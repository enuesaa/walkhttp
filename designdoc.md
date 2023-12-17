# designdoc
- web api の mock を簡単に作れるもの
- feature flag だとかで api を実際に呼ぶか mock を返すか切り替えられるようにしたい
- local api gateway みたいなイメージ
- cache & queue
- serve, invalidate

## development plan
- [http] invoke http url with command line
- [http] create http invoke template
- [http] save histories
- [web] lookup histories
- [http] create pool to wait invoke
- [mock] replay http reponse for mock
