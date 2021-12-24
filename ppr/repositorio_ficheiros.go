package ppr

import "time"

type RepositorioFicheiros interface {
	UltimoFicheiro() string
	ProximoNumeroSequencia(data time.Time) int
}
