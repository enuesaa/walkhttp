# designdoc

- 1セッションでなんでもできなきゃ、とこだわる必要はない
- サブコマンドを定義して、それぞれがプロンプトを立ち上げられればいい

## Commands
```bash
walkhttp get -c walkhttp.json
walkhttp post -c walkhttp.json
walkhttp post -c walkhttp.json -v
walkhttp put -c walkhttp.json
walkhttp delete -c walkhttp.json
walkhttp search -c walkhttp.json
walkhttp -c walkhttp.json # serve web console
```
