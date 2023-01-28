package main

import "github.com/dotdancer/gogofly/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
