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
  walkhttp [command]

Available Commands:
  ctl         Serve web console

Flags:
      --help      Show help information
      --version   Show version

Use "walkhttp [command] --help" for more information about a command.
```

## Development Plan
- [x] [ctlweb] create `ctl` subcommand to serve web console
- [ ] [prompt] create `get`, `post`, `put`, `delete`, `options` subcommand. These commands start prompt with received method.
- [ ] [prompt] implement rich prompt using go-prompt or huh.
