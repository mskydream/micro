package main

import (
	"log"

	"github.com/mskydream/micro/config"
	"github.com/mskydream/micro/database"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	database.Init(&c)
}
