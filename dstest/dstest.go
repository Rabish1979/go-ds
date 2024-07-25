package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	myPlaylist := createLinkList()
	fmt.Println("hello package")
	fmt.Println(reverse.String("Hello"))
}
