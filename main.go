package main

import (
	"github.com/SyncTank/grattor/internal"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	ctx := internal.Init(os.Args)
	log.Println(ctx.Cfg)

}
