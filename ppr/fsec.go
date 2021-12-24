package ppr

import (
	"time"

	"github.com/kindalus/emis_ppr/registos"
)

func GerarFSECFacturas(cfg Config, ctx Contexto, facturas []Factura) ([]registos.Registo, error) {
	var err error
	reg := make([]registos.Registo, len(facturas)+2)

	reg[0], err = gerarRegistoHeaderFSEC(cfg, ctx)

	for i, factura := range facturas {
		reg[i+1], err = gerarRegistoFactura(ctx, factura)
	}

	reg[len(facturas)+1], err = gerarRegistoTrailerFSEC(len(facturas))

	if err != nil {
		return nil, err
	}

	registos.Registos(reg).Sort()
	return reg, nil
}

func gerarRegistoFactura(contexto Contexto, factura Factura) (registos.Registo, error) {
	referencia := contexto.GeradorReferencia.GerarReferencia()
	numeroLinhasTexto := (len(factura.texto) / 40) + 1

	return registos.Gerar(registos.NewCampoTipoRegistoDetalheFactura(),
		registos.NewCampoCodigoProcessamentoAdicionarAlterar(),
		registos.NewCampoReferencia(referencia),
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

func gerarRegistoTrailerFSEC(numeroFacturas int) (registos.Registo, error) {
	return registos.Gerar(registos.NewCampoTipoRegistoTrailer(),
		registos.NewCampoNumeroRegistosDetalhe(numeroFacturas),
		registos.NewFiller(527)), nil

}

func gerarRegistoHeaderFSEC(cfg Config, ctx Contexto) (registos.Registo, error) {
	hoje := time.Now()

	return registos.Gerar(registos.NewCampoTipoRegistoHeader(),
		registos.NewCampoNomeFicheiroFSEC(),
		registos.NewCampoInstituicao(cfg.EntidadeNegocio()),
		registos.NewCampoInstituicaoEGR(),
		registos.NewCampoDataProcessamento(hoje, ctx.Repositorio.ProximoNumeroSequencia(hoje)),
		registos.NewCampoIdUltimoFicheiro(ctx.Repositorio.UltimoFicheiro()),
		registos.NewCampoNumeroEntidade(cfg.IdEntidade()),
		registos.NewFiller(488)), nil
}
