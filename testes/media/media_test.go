package media

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valorRecebido := Media(7.2, 9.9, 6.1, 5.9)
	if valorRecebido != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valorRecebido)
	}
}
