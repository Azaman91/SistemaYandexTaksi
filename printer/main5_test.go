package printer

import (
	"testing"
)

func TestSum(t *testing.T) {
	vv := Sum()
	vi := a + b

	if vv != vi {
		t.Fatalf("Ожидание %d, вывод %d", vi, vv)
	}
}
