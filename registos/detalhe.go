package registos

import (
	"fmt"
	"time"
)

func NewCampoTipoRegistoDetalheFactura() Campo {
	return NewCampo(1, 'N', "3")
}

func NewCampoCodigoProcessamentoAdicionarAlterar() Campo {
	return NewCampo(2, 'N', "80")
}

func NewCampoCodigoProcessamentoRemover() Campo {
	return NewCampo(2, 'N', "82")
}

func NewCampoReferencia() Campo {
	valor := ""
	return NewCampo(15, 'N', valor)
}

func NewCampoIndicadorProduto(produto int) Campo {

	mapaProduto := "1"
	for i := 1; i < 64; i++ {
		mapaProduto = fmt.Sprintf("%v0", mapaProduto)
	}

	return NewCampo(64, 'N', mapaProduto)
}

func NewCampoData(data time.Time) Campo {
	textoData := data.Format("20060102")
	return NewCampo(8, 'N', textoData)
}

func NewCampoMontante(montante float64) Campo {
	return NewCampoDecimal(13, montante)
}

func NewCampoCodigoCliente(codigoCliente int64) Campo {
	return NewCampoNumerico(10, codigoCliente)
}
