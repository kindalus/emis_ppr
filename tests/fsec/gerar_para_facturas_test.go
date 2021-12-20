package fsec_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kindalus/emis_pps/fsec"

	"github.com/stretchr/testify/assert"
)

func Test_Deve_Comecar_Com_Registo_Header(t *testing.T) {
	hoje := time.Now().Format("20060102")
	primeiroRegisto := fmt.Sprintf("0FSEC0900097650000000%v0010000000000000976", hoje)

	registo, _ := fsec.GerarParaFacturas(makeConfig(), makeContexto(), []fsec.Factura{})

	assert.Equal(t, primeiroRegisto, registo[0][:48])
}

func Test_Deve_Acabar_Com_Registo_Trailer(t *testing.T) {
	ultimoRegisto := "900000002"

	registo, _ := fsec.GerarParaFacturas(makeConfig(), makeContexto(), makeFacturas())

	assert.Equal(t, ultimoRegisto, registo[3][:9])
}

func Test_Os_Registos_Devem_Ter_Tamanho_Fixo_536(t *testing.T) {

	registos, _ := fsec.GerarParaFacturas(makeConfig(), makeContexto(), makeFacturas())

	for _, registo := range registos {
		assert.Len(t, registo, 536)

	}
}

func Test_Os_Dois_Ultimos_Digitos_Dos_Campos_4203_4206_Representam_As_Casas_Decimais(t *testing.T) {

	facturas := make([]fsec.Factura, 2)
	agora := time.Now()

	facturas[0], _ = fsec.NewFactura(agora.AddDate(0, 1, 0), 1, 2345.75, "")
	facturas[1], _ = fsec.NewFactura(agora.AddDate(0, 5, 0), 1, 65789.91, "")

	registos, _ := fsec.GerarParaFacturas(makeConfig(), makeContexto(), facturas)

	assert.Equal(t, "0000000234575", registos[1][90:90+13])
	assert.Equal(t, "0000006578991", registos[2][90:90+13])
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

type geradorReferencia struct{}

func (g geradorReferencia) GerarReferencia() string {
	return "000001234567890"
}
