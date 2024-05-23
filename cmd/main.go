package main

import (
	"fmt"

	"github.com/gcsequino/GoOrch/internal/node"
	"github.com/gcsequino/GoOrch/internal/registry"
)

func main() {
	fmt.Println("Hello from main!")
	node.Hello()
	registry.Hello()
}
