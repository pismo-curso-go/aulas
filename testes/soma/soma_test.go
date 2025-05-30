package soma

import "testing"

const mensagem = "esperado %d, obteve %d"

type TestCase struct {
	nome     string
	a, b     int
	esperado int
}

func TestSoma(t *testing.T) {
	casos := []TestCase{
		{"Testar a soma de números positivos", 2, 3, 5},
		{"Testar a soma de números negativos", -1, -1, -2},
	}
	t.Logf("Massa: %v", casos)

	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := Soma(caso.a, caso.b)
			if resultado != caso.esperado {
				t.Errorf(mensagem, caso.esperado, resultado)
			}
		})
	}
}
