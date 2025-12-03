package main

import (
	"fmt"
	"github.com/SyncTank/grattor/internal"
	"os"
)

func main() {
	main := internal.Init(os.Args)
	fmt.Println(main)
}
