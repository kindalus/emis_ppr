package ppr_test

import (
	"testing"
	"time"

	"github.com/kindalus/emis_ppr/pkg/ppr"
	"github.com/stretchr/testify/assert"
)

func Test_Os_Registos_Devem_Ter_Tamanho_Fixo_De_54_Characteres(t *testing.T) {

	registos, _ := ppr.GerarFREF(makeConfig(), makeContexto(), makePagamentos())

	for _, registo := range registos {
		assert.Len(t, registo.String(), 54)
	}
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
