# walkhttp
A CLI tool to call http endpoint with browser or prompt.

> [!NOTE]  
> Work in progress.. `walkhttp` is currently under development.

## Usage [Experimental]
```console
$ walkhttp --help
A CLI tool to call http endpoint with browser or prompt.

Usage:
  walkhttp [flags]

Flags:
      --help       Show help information
      --port int   port (default 3000)
      --version    Show version
```

## Development Plan
- [x] [ctlweb] create `ctl` subcommand to serve web console
- [ ] [prompt] create `get`, `post`, `put`, `delete`, `options` subcommand. These commands start prompt with received method.
- [ ] [prompt] implement rich prompt using go-prompt or huh.
