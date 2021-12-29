package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type RepositorioFicheiros struct {
	mock.Mock
}

func (m *RepositorioFicheiros) UltimoFicheiro() string {
	args := m.Called()
	return args.String(0)
}

func (m *RepositorioFicheiros) ProximoNumeroSequencia(data time.Time) int {
	args := m.Called()
	return args.Int(0)
}
