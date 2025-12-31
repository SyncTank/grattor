package main

import (
	"github.com/SyncTank/grattor/internal"
	"log"
	"os"
)

func main() {
	ctx := internal.Init(os.Args)
	log.Println(ctx.Cfg)
}
