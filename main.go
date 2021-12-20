package main

import (
	"flag"

	"github.com/kindalus/emis_pps/sandbox"
)

func main() {
	sandboxPtr := flag.Bool("sandbox", false, "Corre em sandbox")
	flag.Parse()

	if *sandboxPtr {
		sandbox.Run()
	}

}
