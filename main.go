package main

import (
	_ "fmt"
	"github.com/SyncTank/grattor/internal"
	"os"
)

func main() {
	internal.Init(os.Args)
}
