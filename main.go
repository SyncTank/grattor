package main

import (
	"fmt"
	"github.com/SyncTank/grattor/internal"
)

func main() {
	GConfig, err := internal.ReadConfig()
	internal.Check("main - config check", err)
	fmt.Println(GConfig)

}
