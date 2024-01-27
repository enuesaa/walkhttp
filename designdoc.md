# designdoc
- http client
- statemachine 風味に http request をしたりコマンドを実行できる
- 途中で何か失敗したらそのログなどトレースできるようにしたい

## development plan
- [cli] add command `control`
- [cli] add command `work`
- [db] ~/.walkin/control.db にSQLite3のDBファイルを置く