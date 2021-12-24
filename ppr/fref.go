package ppr

import (
	"time"

	"github.com/kindalus/emis_ppr/registos"
)

func GerarFREF(cfg Config, ctx Contexto, pagamentos []Pagamento) ([]registos.Registo, error) {
	var err error
	quantidade := len(pagamentos)

	reg := make([]registos.Registo, quantidade+2)

	reg[0], err = gerarRegistoHeaderFREF(cfg, ctx)

	for i, pagamento := range pagamentos {
		reg[i+1], err = gerarRegistoReferencia(ctx, pagamento)
	}

	reg[quantidade+1], err = gerarRegistoTrailerFREF(quantidade)

	if err != nil {
		return nil, err
	}

	registos.Registos(reg).Sort()
	return reg, nil
}

func gerarRegistoHeaderFREF(config Config, ctx Contexto) (registos.Registo, error) {
	hoje := time.Now()

	return registos.Gerar(
			registos.NewCampoTipoRegistoHeader(),
			registos.NewCampoNomeFicheiroFREF(),
			registos.NewCampoInstituicao(config.EntidadeNegocio()),
			registos.NewCampoInstituicaoEGR(),
			registos.NewCampoDataProcessamento(hoje, ctx.Repositorio.ProximoNumeroSequencia(hoje)),
			registos.NewCampoIdUltimoFicheiro(ctx.Repositorio.UltimoFicheiro()),
			registos.NewCampoNumeroEntidade(config.IdEntidade()),
			registos.NewFiller(6)),
		nil
}

func gerarRegistoTrailerFREF(quantidade int) (registos.Registo, error) {
	return registos.Gerar(
			registos.NewCampoTipoRegistoTrailer(),
			registos.NewCampoNumerico(8, int64(quantidade)),
			registos.NewFiller(45)),
		nil
}

func gerarRegistoReferencia(contexto Contexto, pagamento Pagamento) (registos.Registo, error) {

	referencia := contexto.GeradorReferencia.GerarReferencia()

	return registos.Gerar(
			registos.NewCampoTipoRegistoDetalheFREF(),
			registos.NewCampoCodigoProcessamentoAdicionarAlterar(),
			registos.NewCampoReferencia(referencia),
			registos.NewCampoCodigoRefenciaComum(1),
			registos.NewCampoData(pagamento.dataLimitePagamento),
			registos.NewFiller(20)),
		nil
}
