package ppr_test

import (
	"testing"

	"github.com/kindalus/emis_ppr/ppr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_New_Config(t *testing.T) {
	config, err := ppr.NewConfig("00976", "09000976", "99")

	require.Nil(t, err)
	require.NotEqual(t, config, ppr.Config{})

	assert.Equal(t, config.IdEntidade(), "00976")
	assert.Equal(t, config.EntidadeNegocio(), "09000976")
	assert.Equal(t, config.TipoProduto(), "99")
}

func Test_New_Config_Tipo_Negocio_Deve_Ter_2_Caracters_Numericos(t *testing.T) {
	config1, err1 := ppr.NewConfig("OO976", "09000976", "9")
	config2, err2 := ppr.NewConfig("OO976", "09000976", "990")
	config3, err3 := ppr.NewConfig("OO976", "09000976", "9X")

	require.NotNil(t, err1)
	require.NotNil(t, err2)
	require.NotNil(t, err3)

	assert.Equal(t, config1, ppr.Config{})
	assert.Equal(t, config2, ppr.Config{})
	assert.Equal(t, config3, ppr.Config{})
}

func Test_New_Config_Id_Entidade_Deve_Ter_5_Caracters_Numericos(t *testing.T) {
	config1, err1 := ppr.NewConfig("0097", "09000976", "99")
	config2, err2 := ppr.NewConfig("009764", "09000976", "99")
	config3, err3 := ppr.NewConfig("OO976", "09000976", "99")

	require.NotNil(t, err1)
	require.NotNil(t, err2)
	require.NotNil(t, err3)

	assert.Equal(t, config1, ppr.Config{})
	assert.Equal(t, config2, ppr.Config{})
	assert.Equal(t, config3, ppr.Config{})
}

func Test_New_Config_Entidade_Negocio_Deve_Ter_8_Caracters_Numericos(t *testing.T) {
	config1, err1 := ppr.NewConfig("00976", "9000976", "99")
	config2, err2 := ppr.NewConfig("00976", "9909000976", "99")
	config3, err3 := ppr.NewConfig("OO976", "x9000976", "99")

	require.NotNil(t, err1)
	require.NotNil(t, err2)
	require.NotNil(t, err3)

	assert.Equal(t, config1, ppr.Config{})
	assert.Equal(t, config2, ppr.Config{})
	assert.Equal(t, config3, ppr.Config{})
}
