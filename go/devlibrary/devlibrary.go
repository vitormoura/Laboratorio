package main

import (
	"fmt"
	"github.com/vitormoura/Laboratorio/go/devlibrary/web"
)

func main() {
	fmt.Println("Listen to :4001")
	web.Start(4001)
}
