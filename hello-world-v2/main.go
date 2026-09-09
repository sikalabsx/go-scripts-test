package main

import (
	"flag"

	"github.com/sikalabsx/go-scripts-test/hello-world-v2/pkg/hello_world"
)

func main() {
	name := flag.String("name", "World", "Name to greet")

	flag.Parse()

	hello_world.PrintHelloWorld(*name)
}
