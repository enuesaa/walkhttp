//go:build dev

package ctlweb

func init() {
	go RunDevCmd()
}

var Serve = ServeDev
