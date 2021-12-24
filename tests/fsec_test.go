package ppr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kindalus/emis_ppr/mocks"
	"github.com/kindalus/emis_ppr/ppr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Deverá estar ordenado de forma crescente por número de referência
// (tipos de registo 1 e registo 3).
func Test_Registos_Devem_Estar_Ordenados_Por_Numero_Numero_De_Referencia(t *testing.T) {
	registos, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), makeOitoFacturas())

	lastRef := ""

	for _, registo := range registos {
		if registo.TipoRegisto() == "1" {
			continue
		}

		if registo.TipoRegisto() == "9" {
			break
		}

		assert.GreaterOrEqual(t, registo.Referencia(), lastRef)

		lastRef = registo.Referencia()
	}
}

// O primeiro ficheiro que a entidade envia deve ter o campo
// “identificação do último ficheiro” – atributo 2399 preenchido a zeros.
func Test_O_Primeiro_Ficheiro_Deve_Ter_O_Campo_Ultimo_Ficheiro_Preenchido_A_Zeros(t *testing.T) {
	hoje := time.Now().Format("20060102")
	primeiroRegisto := fmt.Sprintf("0FSEC0900097650000000%v0010000000000000976", hoje)

	registo, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), []ppr.Factura{})

	assert.Equal(t, primeiroRegisto, registo[0].String()[:48])
}

// Em cada ficheiro que produz, a Entidade deve referir a identificação do ficheiro
// anteriormente enviado, de modo a que possam ser controladas eventuais falhas de envio.
func Test_Ficheiro_Deve_Referir_Ficheiro_Anterior(t *testing.T) {

	mockRepo := new(mocks.RepositorioFicheiros)
	mockRepo.On("ProximoNumeroSequencia", mock.Anything).Return(1)
	mockRepo.On("UltimoFicheiro").Return("20211213001")

	ctx := makeContexto()
	ctx.Repositorio = mockRepo

	registos, _ := ppr.GerarFSECFacturas(makeConfig(), ctx, []ppr.Factura{})

	assert.Equal(t, "20211213001", registos[0].String()[32:32+11])

}

// Em cada ficheiro que produz, a Entidade deve referir a identificação do ficheiro
// anteriormente enviado, de modo a que possam ser controladas eventuais falhas de envio.
func Test_Id_Do_Ficheiro_Composto_Por_Data_E_Sequencia(t *testing.T) {

	idFicheiro := time.Now().Format("20060102") + "004"

	mockRepo := new(mocks.RepositorioFicheiros)
	mockRepo.On("ProximoNumeroSequencia", mock.Anything).Return(4)
	mockRepo.On("UltimoFicheiro").Return("00000000000")

	ctx := makeContexto()
	ctx.Repositorio = mockRepo

	registos, _ := ppr.GerarFSECFacturas(makeConfig(), ctx, makeFacturas())

	assert.Equal(t, idFicheiro, registos[0].String()[21:21+11])

}

func Test_Deve_Comecar_Com_Registo_Header(t *testing.T) {
	hoje := time.Now().Format("20060102")
	primeiroRegisto := fmt.Sprintf("0FSEC0900097650000000%v0010000000000000976", hoje)

	registo, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), []ppr.Factura{})

	assert.Equal(t, primeiroRegisto, registo[0].String()[:48])
}

func Test_Deve_Acabar_Com_Registo_Trailer(t *testing.T) {
	ultimoRegisto := "900000002"

	registos, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), makeFacturas())

	assert.Equal(t, ultimoRegisto, string(registos[3][:9]))
}

func Test_Os_Registos_Devem_Ter_Tamanho_Fixo_De_536_Characteres(t *testing.T) {

	registos, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), makeFacturas())

	for _, registo := range registos {
		assert.Len(t, registo.String(), 536)

	}
}

func Test_Os_Dois_Ultimos_Digitos_Dos_Campos_4203_4206_Representam_As_Casas_Decimais(t *testing.T) {

	facturas := make([]ppr.Factura, 1)
	agora := time.Now()

	facturas[0], _ = ppr.NewFactura(agora.AddDate(0, 1, 0), 1, 2345.75, "")

	registos, _ := ppr.GerarFSECFacturas(makeConfig(), makeContexto(), facturas)

	assert.Equal(t, "0000000234575", registos[1].MontanteMaximo())
}

func makeFacturas() []ppr.Factura {
	agora := time.Now()
	facturas := make([]ppr.Factura, 2)

	facturas[0], _ = ppr.NewFactura(
		agora.AddDate(0, 1, 0),
		5,
		23.50,
		"Entregas em 24 horas")

	facturas[1], _ = ppr.NewFactura(
		agora.AddDate(0, 0, 15),
		1,
		200.00,
		"Entregas em 72 horas")

	return facturas
}

func makeOitoFacturas() []ppr.Factura {
	agora := time.Now()
	facturas := make([]ppr.Factura, 8)

	facturas[0], _ = ppr.NewFactura(agora.AddDate(0, 1, 00), 5, 0023.50, "")
	facturas[1], _ = ppr.NewFactura(agora.AddDate(0, 0, 15), 1, 0200.00, "")
	facturas[2], _ = ppr.NewFactura(agora.AddDate(0, 0, 15), 1, 2010.00, "")
	facturas[3], _ = ppr.NewFactura(agora.AddDate(0, 0, 10), 1, 1200.00, "")
	facturas[4], _ = ppr.NewFactura(agora.AddDate(0, 0, 15), 1, 1000.00, "")
	facturas[5], _ = ppr.NewFactura(agora.AddDate(0, 0, 15), 1, 0830.00, "")
	facturas[6], _ = ppr.NewFactura(agora.AddDate(0, 0, 15), 1, 0750.00, "")
	facturas[7], _ = ppr.NewFactura(agora.AddDate(0, 0, 01), 1, 0089.99, "")

	return facturas
}
