package ppr

import "time"

type Pagamento struct {
	dataLimitePagamento time.Time
	valor               float64
}

func NewPagamento(dataLimitePagamento time.Time,
) (Pagamento, error) {

	return Pagamento{
		dataLimitePagamento: dataLimitePagamento,
	}, nil
}
