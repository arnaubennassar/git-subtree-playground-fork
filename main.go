package main

import (
	"fmt"

	"github.com/arnaubennassar/git-subtree-playground-fork/hello"
)

func main() {
	fmt.Println("hello from the fork repo")
	hello.SayHello()
	hello.SayHelloFromInternalDependency()
	hello.SayHelloFromForkDep()
	hello.SayHelloToConflicts()
}
