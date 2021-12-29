package registos

import (
	"sort"
	"strings"
)

func Gerar(campos ...Campo) Registo {
	valoresCampos := make([]string, len(campos))

	for i, v := range campos {
		valoresCampos[i] = v.ToString()
	}

	return Registo(strings.Join(valoresCampos, ""))
}

type Registos []Registo

func (r Registos) Len() int {
	return len(r)
}

func (r Registos) Less(i, j int) bool {

	if r[i].TipoRegisto() == "0" {
		return true
	}

	if r[i].TipoRegisto() == "9" {
		return false
	}

	return strings.Compare(r[i].Referencia(), r[j].Referencia()) < 0
}

func (r Registos) Swap(i, j int) {
	vi := r[i]

	r[i] = r[j]
	r[j] = vi
}

func (r Registos) Sort() {
	sort.Sort(r)
}

type Registo string

func (r Registo) String() string {
	return string(r)
}

func (r Registo) TipoRegisto() string {
	return r.String()[:1]
}

func (r Registo) campoSeTipoRegisto(comeco int, final int, tipoRegisto string) string {
	if r.TipoRegisto() != tipoRegisto {
		return ""
	}

	return string(r[comeco:final])
}

func (r Registo) Referencia() string {
	return r.campoSeTipoRegisto(3, 3+15, "3")

}

func (r Registo) MontanteMaximo() string {
	return r.campoSeTipoRegisto(90, 90+13, "3")
}
