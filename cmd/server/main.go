package main

import "github.com/luis13005/pos-go/configs"

func main() {
	confg := configs.LoadConfig(".")
	println(confg.DBDriver)
}
