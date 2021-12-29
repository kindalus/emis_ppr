package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/kindalus/emis_ppr/pkg/ppr"
)

func main() {

	sandFREF := flag.Bool("fref", false, "Corre em sandbox (FREF)")
	sss := flag.Int("s", 1, "Sequência do Ficheiro")
	ultimoFicheiro := flag.String("u", "00000000000", "Id do último ficheiro")

	flag.Parse()

	if *sandFREF {
		RunFREF(*sss, *ultimoFicheiro)
		return
	}

	RunFSEC(*sss, *ultimoFicheiro)

}

func RunFSEC(sequencia int, ultimoFicheiro string) {
	registos, err := ppr.GerarFSECFacturas(makeConfig(), makeContexto(sequencia, ultimoFicheiro), makeFacturas())

	if err != nil {
		panic(err)
	}

	for _, registo := range registos {
		fmt.Println(registo)
	}
}

func RunFREF(sequencia int, ultimoFicheiro string) {
	registos, err := ppr.GerarFREF(makeConfig(), makeContexto(sequencia, ultimoFicheiro), makePagamentos())

	if err != nil {
		panic(err)
	}

	for _, registo := range registos {
		fmt.Println(registo)
	}
}

type repositorioFicheiros struct {
	ultimoFicheiro string
	sequencia      int
}

func (r repositorioFicheiros) UltimoFicheiro() string {
	return r.ultimoFicheiro
}

func (r repositorioFicheiros) ProximoNumeroSequencia(data time.Time) int {
	return r.sequencia
}

func makeConfig() ppr.Config {
	config, _ := ppr.NewConfig("00976", "09000976", "99")

	return config
}

func makeContexto(sequencia int, ultimoFicheiro string) ppr.Contexto {
	return ppr.Contexto{
		GeradorReferencia: ppr.NewGeradorReferencia(),
		Repositorio: repositorioFicheiros{
			sequencia:      sequencia,
			ultimoFicheiro: ultimoFicheiro,
		},
	}
}

func makeFacturas() []ppr.Factura {
	agora := time.Now()
	facturas := make([]ppr.Factura, 1)

	facturas[0], _ = ppr.NewFactura(
		agora.AddDate(0, 1, 0),
		1,
		100.000,
		"Factura de teste")

	return facturas
}

func makePagamentos() []ppr.Pagamento {
	hoje := time.Now()
	pagamentos := make([]ppr.Pagamento, 4)

	pagamentos[0], _ = ppr.NewPagamento(hoje.AddDate(0, 1, 0))
	pagamentos[1], _ = ppr.NewPagamento(hoje.AddDate(0, 0, 15))
	pagamentos[2], _ = ppr.NewPagamento(hoje.AddDate(0, 0, 7))
	pagamentos[3], _ = ppr.NewPagamento(hoje.AddDate(0, 2, 0))

	return pagamentos
}
