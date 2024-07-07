# walkhttp
A CLI tool to call http endpoint with browser or prompt.

> [!NOTE]  
> Work in progress.. `walkhttp` is currently under development.

## Usage [Experimental]
```console
$ walkhttp --help
A CLI tool to call http endpoint with browser or prompt.

USAGE:
  walkhttp [flags] command [command options] [arguments...]

COMMANDS:
   get      make http GET request
   post     make http POST request
   put      make http PUT request
   delete   make http DELETE request
   options  make http OPTIONS request

FLAGS:
  --port value              Serve port (default: 3000)
  --config value, -c value  config file path (default: "walkhttp.json")
  --help, -h                show help
  --version, -v             print the version
```

## Development Plan
- [x] [prompt] create `get`, `post`, `put`, `delete`, `options` subcommand. These commands start prompt with received method.
- [ ] [prompt] implement rich prompt using go-prompt or huh.

### Background
- 「1セッションでなんでもできる」必要はない
- サブコマンドを定義して、それぞれがプロンプトを立ち上げられればいい

### Commands
```bash
walkhttp get -c walkhttp.json
walkhttp post -c walkhttp.json
walkhttp post -c walkhttp.json -v
walkhttp put -c walkhttp.json
walkhttp delete -c walkhttp.json
walkhttp search -c walkhttp.json
walkhttp -c walkhttp.json # serve web console
```
