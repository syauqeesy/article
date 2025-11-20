package main

import (
	"log"
	"os"

	"ahmadsyauqi.dev/article/foundation"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go [http|migration]")
	}

	foundationType := os.Args[1]
	arguments := make([]string, 0, 2)
	err := foundation.Boot(foundationType, arguments)
	if err != nil {
		log.Fatal(err.Error())
	}
}
