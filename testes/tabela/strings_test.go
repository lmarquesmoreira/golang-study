package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) -  indices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		text     string
		part     string
		expected int
	}{
		{"Coder é show", "Coder", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"lucas", "a", 3},
	}

	for _, teste := range testes {
		t.Logf("--> %v", teste)
		actual := strings.Index(teste.text, teste.part)

		if actual != teste.expected {
			t.Errorf(msgIndex, teste.text, teste.part, teste.expected, actual)
		}
	}
}
