package main

import (
	"flag"

	"github.com/kindalus/emis_pps/sandbox"
)

func main() {
	sandFSEC := flag.Bool("ss", false, "Corre em sandbox (FSEC)")
	sandFREF := flag.Bool("sr", false, "Corre em sandbox (FREF)")
	//sss := flag.Int("s", 1, "Sequência do Ficheiro")
	//ultimoFicheiro := flag.String("u", "00000000000", "Id dio último ficheiro")

	flag.Parse()

	if *sandFSEC {
		sandbox.RunFSEC()
	}

	if *sandFREF {
		sandbox.RunFREF()
	}

}
