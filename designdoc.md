# designdoc
- http client
- 途中で何か失敗したらそのログなどトレースできるようにしたい
- call http api from browser or server (golang)
- http request がコアで、ミニマムの実行単位であり、これを登録 (api gateway) して再利用できたり、あるいは state machine 風味にロジックを入れられればいいなあ

## development plan
- [cli] add command `serve`
- [db] ~/.walkin/logs にログファイルを置く
