package main

import (
	"farm.e-pedion.com/repo/config"
	"fmt"
)

func main() {
	fmt.Println("SetupErr =", config.Setup())
}
