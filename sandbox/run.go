package sandbox

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kindalus/emis_pps/fsec"
)

func Run() {
	registos, err := fsec.GerarParaFacturas(makeConfig(), makeContexto(), makeFacturas())

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

func makeConfig() fsec.Config {
	config, _ := fsec.NewConfig("00976", "09000976", "99")

	return config
}

func makeContexto() fsec.Contexto {
	return fsec.Contexto{
		GeradorReferencia: geradorReferencia{},
	}
}

func makeFacturas() []fsec.Factura {
	agora := time.Now()
	facturas := make([]fsec.Factura, 2)

	facturas[0], _ = fsec.NewFactura(
		agora.AddDate(0, 1, 0),
		5,
		23.50,
		"Entregas em 24 horas")

	facturas[1], _ = fsec.NewFactura(
		agora.AddDate(0, 0, 15),
		1,
		200.00,
		"Entregas em 72 horas")

	return facturas
}
