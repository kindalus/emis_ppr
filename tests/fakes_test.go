package ppr_test

import (
	"math/rand"

	"github.com/kindalus/emis_pps/mocks"
	"github.com/kindalus/emis_pps/ppr"
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
		GeradorReferencia: geradorReferencia{},
		Repositorio:       repo,
	}
}

type geradorReferencia struct{}

func (g geradorReferencia) GerarReferencia() string {
	referencias := []string{
		"000001234567890",
		"000401234567891",
		"000485836456867",
		"000920629512510",
		"000997970227015",
		"000525870304065",
		"000333885905807",
		"000191897259696",
	}

	return referencias[rand.Intn(len(referencias))]
}
