package registos

func NewCampoTipoRegistoTrailer() Campo {
	return NewCampo(1, 'N', "9")
}

func NewCampoNumeroRegistosDetalhe(quantity int) Campo {

	return NewCampoNumerico(8, int64(quantity))
}
