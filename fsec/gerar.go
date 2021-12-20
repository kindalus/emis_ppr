package fsec

import (
	"time"

	"github.com/kindalus/emis_pps/registos"
)

func GerarParaFacturas(config Config, contexto Contexto, facturas []Factura) ([]string, error) {
	var err error

	numeroFacturas := len(facturas)
	numeroRegistos := numeroFacturas + 2
	indexTrailer := numeroFacturas + 1

	registos := make([]string, numeroRegistos)

	registos[0], err = gerarRegistoHeader(config)

	for i, factura := range facturas {
		registos[i+1], err = gerarRegistoFactura(contexto, factura)
	}

	registos[indexTrailer], err = gerarRegistoTrailer(numeroFacturas)

	if err != nil {
		return nil, err
	}

	return registos, nil
}

func gerarRegistoFactura(contexto Contexto, factura Factura) (string, error) {
	referencia := contexto.GeradorReferencia.GerarReferencia()
	numeroLinhasTexto := (len(factura.texto) / 40) + 1

	return registos.Gerar(registos.NewCampoTipoRegistoDetalheFactura(),
		registos.NewCampoCodigoProcessamentoAdicionarAlterar(),
		registos.NewCampo(15, 'N', referencia),
		registos.NewCampoIndicadorProduto(0),
		registos.NewCampoData(factura.dataLimitePagamento),
		registos.NewCampoMontante(factura.valor),
		registos.NewCampoData(time.Now()),
		registos.NewCampoMontante(factura.valor),
		registos.NewCampoCodigoCliente(factura.codigoCliente),
		registos.NewCampoNumerico(2, int64(numeroLinhasTexto)),
		registos.NewCampo(400, 'A', factura.texto),
	), nil
}

func gerarRegistoTrailer(numeroFacturas int) (string, error) {
	return registos.Gerar(registos.NewCampoTipoRegistoTrailer(),
		registos.NewCampoNumeroRegistosDetalhe(numeroFacturas),
		registos.NewFiller(527)), nil

}

func gerarRegistoHeader(config Config) (string, error) {
	hoje := time.Now()
	sequencia := 1
	ultimoFicheiro := "00000000000"

	return registos.Gerar(registos.NewCampoTipoRegistoHeader(),
		registos.NewCampoNomeFicheiroFSEC(),
		registos.NewCampoInstituicao(config.entidadeNegocio),
		registos.NewCampoInstituicaoEGR(),
		registos.NewCampoDataProcessamento(hoje, sequencia),
		registos.NewCampoIdUltimoFicheiro(ultimoFicheiro),
		registos.NewCampoNumeroEntidade(config.IdEntidade()),
		registos.NewFiller(488)), nil
}
