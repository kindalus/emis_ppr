package registos

import "strings"

func Gerar(campos ...Campo) string {
	valoresCampos := make([]string, len(campos))

	for i, v := range campos {
		valoresCampos[i] = v.ToString()
	}

	return strings.Join(valoresCampos, "")
}
