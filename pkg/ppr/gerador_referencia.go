package ppr

import (
	"fmt"
	"math/rand"
	"time"
)

type GeradorReferencia interface {
	GerarReferencia() string
}

func NewGeradorReferencia() GeradorReferencia {
	return new(geradorReferencia)
}

type geradorReferencia struct{}

func (g geradorReferencia) GerarReferencia() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	referencia := random.Intn(999_999_999)

	return fmt.Sprintf("%015d", referencia)
}
