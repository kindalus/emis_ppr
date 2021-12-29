package registos

import (
	"fmt"
	"time"
)

const instituicaoDestinoEGR = "50000000"

func NewCampoTipoRegistoHeader() Campo {
	return NewCampo(1, 'N', "0")
}

func NewCampoNomeFicheiroFSEC() Campo {
	return NewCampo(4, 'A', "FSEC")
}

func NewCampoNomeFicheiroFREF() Campo {
	return NewCampo(4, 'A', "FREF")
}

func NewCampoInstituicao(entidade string) Campo {
	return NewCampo(8, 'N', entidade)
}

func NewCampoInstituicaoEGR() Campo {
	return NewCampoInstituicao(instituicaoDestinoEGR)
}

func NewCampoDataProcessamento(data time.Time, sequencia int) Campo {
	textoSequencia := fmt.Sprintf("%d", 1000+sequencia)[1:]
	textoData := data.Format("20060102")
	valor := textoData + textoSequencia

	return NewCampo(11, 'N', valor)
}

func NewCampoIdUltimoFicheiro(valor string) Campo {
	if valor == "" {
		return NewCampo(11, 'N', "00000000000")
	}
	return NewCampo(11, 'N', valor)
}

func NewCampoNumeroEntidade(valor string) Campo {
	return NewCampo(5, 'N', valor)
}
