package CountOnes

import (
	"testing"
)


func test16(t *testing.T) {
	t.Log("Test 16 bits")
	if ones := count("1010101010101010"); ones != 8 {
		t.Error("Esperaba 8, pero obtuve ", ones)
	}

	t.Log("Test 0")
	if cero := count("00000000000000000000000000000000000000000"); cero != 0 {
		t.Error("Esperaba 8, pero obtuve ", cero)
	}	
}
