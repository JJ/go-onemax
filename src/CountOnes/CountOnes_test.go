package CountOnes

import (
	"testing"
)


func TestShort(t *testing.T) {
	t.Log("Test 8 bits")
	if ones_8bits := count("10101010"); ones_8bits != 4 {
		t.Error("Esperaba 4, pero obtuve ", ones_8bits)
	}
}

func TestLonger(t *testing.T) {
	t.Log("Test 16 bits")
	if ones := count("1010101010101010"); ones != 8 {
		t.Error("Esperaba 8, pero obtuve ", ones)
	}
}

func TestZero(t *testing.T) {
	t.Log("Test 0")
	if cero := count("00000000000000000000000000000000000000000"); cero != 0 {
		t.Error("Esperaba 8, pero obtuve ", cero)
	}	
}
