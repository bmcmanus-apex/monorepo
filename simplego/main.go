package main

import (
	"fmt"

	"github.com/bmcmanus-apex/monorepo/simplego/fortune"
)

func main() {
	fmt.Println("Hello, Bazel! 💚")
	fmt.Println(fortune.Get())
}
