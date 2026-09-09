package hello_world

import "fmt"

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name + "!"
}

func PrintHelloWorld(name string) {
	fmt.Println(HelloWorld(name))
}
