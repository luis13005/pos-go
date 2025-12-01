package main

import "github.com/luis13005/apis/configs"

func main() {
	confg := configs.LoadConfig(".")
	println(confg.DBDriver)
}
