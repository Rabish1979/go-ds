package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	myPlaylist := createLinkList()
	myPlaylist.insertAtFront(20)
	myPlaylist.insertAtFront(30)
	myPlaylist.insertAtFront(40)
	myPlaylist.Print()
	fmt.Println("hello package")
	fmt.Println(reverse.String("Hello"))
}
