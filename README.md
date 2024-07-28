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
  --port value                Serve port (default: 3000)
  --workspace value, -w value workspace file path (default: "walkhttp.json")
  --help, -h                  show help
  --version, -v               print the version
```
