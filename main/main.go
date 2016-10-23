package main

import (
	"farm.e-pedion.com/repo/config"
	"fmt"
)

var (
	cfg = config.Get()
)

func init() {
	fmt.Printf("main.init config=%+v\n", cfg)
}

func main() {
	fmt.Printf("config.main config=%+v\n", cfg)
}
