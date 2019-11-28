package main

import "fmt"

const (
	hello       = "你好"
	exclamation = "!"
)

var world string

func 世界() {
	world = "世界"
}

func init() {
	世界()
}

func main() {
	fmt.Println(hello + world + exclamation)
}
