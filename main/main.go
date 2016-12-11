package main

import (
	"fmt"
	"github.com/rjansen/migi"
)

var (
	cfg = migi.Get()
)

func init() {
	fmt.Printf("main.init config=%+v\n", cfg)
}

func main() {
	fmt.Printf("config.main config=%+v\n", cfg)
}
