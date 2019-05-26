package main

import (
	"deploy-tools/handle"
	"fmt"
)

func main() {

	fmt.Println("hello, deploy-tools!")
	fmt.Println()

	handle.InitCfg("./config.yml")
}

// handle the configuration
