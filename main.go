package main

import (
	"fmt"
	"log"
	_ "log"

	"github.com/SyncTank/grattor/internal"
)

func main() {
	GConfig, err := internal.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(GConfig)

}
