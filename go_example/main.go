package main

import (
	"go_example/server"
)

func main() {
	server.NewBuilder().
		SetPort("8080").
		SetReadTimeout(10).
		SetWriteTimeout(10).
		BuildAndServe()
}
