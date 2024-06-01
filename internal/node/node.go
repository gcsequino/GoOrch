package node

import (
	"fmt"
)

func Hello(registryAddr string) {
	fmt.Printf("Connecting to registry at %s\n", registryAddr)
}
