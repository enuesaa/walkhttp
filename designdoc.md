# designdoc
- http client
- 途中で何か失敗したらそのログなどトレースできるようにしたい
- call http api from browser or server (golang)
- http request がコアでミニマムの実行単位. api gateway に登録すれば再利用でき、また state machine 風味にロジックを入れられる

## development plan
- [db] ~/.walkin/logs にログファイルを置く
- [cli] reuse with history