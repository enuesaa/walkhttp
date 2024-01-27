//go:build dev

package controlweb

func init() {
	go RunDevCmd()
}

var Serve = ServeDev
