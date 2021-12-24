package sandbox

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kindalus/emis_pps/ppr"
)

func RunFSEC() {
	registos, err := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), makeFacturas())

	if err != nil {
		panic(err)
	}

	for _, registo := range registos {
		fmt.Println(registo)
	}
}

func RunFREF() {
	registos, err := ppr.GerarFREF(makeConfig(), makeContexto(), makePagamentos())

	if err != nil {
		panic(err)
	}

	for _, registo := range registos {
		fmt.Println(registo)
	}
}

type geradorReferencia struct{}

func (g geradorReferencia) GerarReferencia() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	referencia := random.Intn(999_999_999_999)

	return fmt.Sprintf("%015d", referencia)
}

func makeConfig() ppr.Config {
	config, _ := ppr.NewConfig("00976", "09000976", "99")

	return config
}

func makeContexto() ppr.Contexto {
	return ppr.Contexto{
		GeradorReferencia: geradorReferencia{},
	}
}

func makeFacturas() []ppr.Factura {
	agora := time.Now()
	facturas := make([]ppr.Factura, 1)

	facturas[0], _ = ppr.NewFactura(
		agora.AddDate(0, 1, 0),
		5,
		1000.000,
		"Factura de teste")

	// facturas[1], _ = ppr.NewFactura(
	// 	agora.AddDate(0, 0, 15),
	// 	1,
	// 	200.00,
	// 	"Entregas em 72 horas")

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
