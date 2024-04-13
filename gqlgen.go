//go:build tools

package main

import (
	_ "github.com/99designs/gqlgen"
)

//go:generate go run github.com/99designs/gqlgen generate
