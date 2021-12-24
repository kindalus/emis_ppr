package ppr

import (
	"fmt"
	"strconv"
)

type Config struct {
	idEntidade      string
	entidadeNegocio string
	tipoProduto     string
}

func NewConfig(id string, entidade string, tipo string) (Config, error) {

	if errId := validarIDEntidade(id); errId != nil {
		return Config{}, errId
	}

	if errID := validarTipoProduto(tipo); errID != nil {
		return Config{}, errID
	}

	if errId := validarEntidadeNegocio(entidade); errId != nil {
		return Config{}, errId
	}

	return Config{idEntidade: id, entidadeNegocio: entidade, tipoProduto: tipo}, nil
}

func validarIDEntidade(id string) error {
	return validar(id, 5, "Id da Entidade")
}

func validarTipoProduto(tipo string) error {
	return validar(tipo, 2, "Tipo de Produto")
}

func validarEntidadeNegocio(tipo string) error {
	return validar(tipo, 8, "Entidade Negocio")
}

func validar(id string, tamanho int, nomeCampo string) error {
	if len(id) != tamanho {
		return fmt.Errorf("O tamanho do <%v> eh diferente de %v", nomeCampo, tamanho)
	}

	if !ehNumerico(id) {
		return fmt.Errorf("O valor do <%v> nao eh numerico", nomeCampo)
	}

	return nil
}

func ehNumerico(valor string) bool {
	_, err := strconv.ParseInt(valor, 10, 64)

	return err == nil
}

func (c Config) IdEntidade() string {
	return c.idEntidade
}

func (c Config) EntidadeNegocio() string {
	return c.entidadeNegocio
}

func (c Config) TipoProduto() string {
	return c.tipoProduto
}
