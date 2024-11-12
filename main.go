package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/keypairs"
)

func main() {

	bytes, err := keypairs.DeriveNodeAddress("n9KHn8NfbBsZV5q8bLfS72XyGqwFt5mgoPbcTV4c6qKiuPTAtXYk")
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes, len(bytes))
}
