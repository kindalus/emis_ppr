package registos

import (
	"fmt"
	"strconv"
)

const (
	TipoNumerico     = "N"
	TipoAlfanumerico = "A"
)

type Campo struct {
	comprimento int
	tipo        rune
	valor       string
}

func NewCampo(comprimento int, tipo rune, valor string) Campo {

	mascara := "%-" + strconv.Itoa(comprimento) + "s"
	return Campo{comprimento: comprimento, tipo: tipo, valor: fmt.Sprintf(mascara, valor)}
}

func NewCampoNumerico(comprimento int, valor int64) Campo {

	mascara := "%0" + strconv.Itoa(comprimento) + "d"

	valorDoCampo := fmt.Sprintf(mascara, valor)

	return NewCampo(comprimento, 'N', valorDoCampo)
}

func NewCampoDecimal(comprimento int, valor float64) Campo {

	return NewCampoNumerico(comprimento, int64(valor*100))
}

func (c Campo) ToString() string {
	return c.valor
}
