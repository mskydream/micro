package main

import (
	"fmt"
	"log"

	"github.com/mskydream/micro/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)
	// database.Init()
}
