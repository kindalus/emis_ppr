package ppr_test

import (
	"github.com/kindalus/emis_ppr/mocks"
	"github.com/kindalus/emis_ppr/ppr"
	"github.com/stretchr/testify/mock"
)

func makeConfig() ppr.Config {

	config, _ := ppr.NewConfig("00976", "09000976", "99")

	return config
}

func makeContexto() ppr.Contexto {

	repo := new(mocks.RepositorioFicheiros)
	repo.On("UltimoFicheiro").Return("00000000000")
	repo.On("ProximoNumeroSequencia", mock.Anything).Return(1)

	return ppr.Contexto{
		GeradorReferencia: ppr.NewGeradorReferencia(),
		Repositorio:       repo,
	}
}
