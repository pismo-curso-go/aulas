package multiplicar

import "testing"

func TestMultiplicar(t *testing.T) {

	t.Run("Multiplicação de dois positivos", func(t *testing.T) {
		resultado := Multiplicar(5, 5)
		esperado := 25.0
		if resultado != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
		}
	})

	t.Run("Multiplicação com zero", func(t *testing.T) {
		resultado := Multiplicar(10, 0)
		esperado := 0.0
		if resultado != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
		}
	})
	t.Run("Multiplicação de dois negativos", func(t *testing.T) {
		resultado := Multiplicar(-4, -5)
		esperado := 20.0
		if resultado != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
		}
	})
	t.Run("Multiplicação de positivo por negativo", func(t *testing.T) {
		resultado := Multiplicar(6, -2)
		esperado := -12.0
		if resultado != esperado {
			t.Errorf("Esperado %.2f, mas obteve %.2f", esperado, resultado)
		}
	})
}
