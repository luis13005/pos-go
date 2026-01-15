package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4", "teste5", "teste6", "teste7", "teste8", "teste9", "teste10"}
	evento = append(evento[:4], evento[5:]...)
	fmt.Println(evento)
}
