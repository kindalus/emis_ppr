package fsec

import "time"

type Factura struct {
	dataLimitePagamento time.Time
	codigoCliente       int64
	valor               float64
	texto               string
}

func NewFactura(dataLimitePagamento time.Time,
	codigoCliente int64,
	valor float64,
	texto string) (Factura, error) {

	return Factura{
		dataLimitePagamento: dataLimitePagamento,
		codigoCliente:       codigoCliente,
		valor:               valor,
		texto:               texto,
	}, nil
}
