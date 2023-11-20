package main

import (
	"fmt"
	go_say_hello "github.com/wirawan-mahardika/go-say-hello/v2"
)

func main() {
	data := go_say_hello.SayHello("wirawan mahardika putra")
	fmt.Println(data)
}