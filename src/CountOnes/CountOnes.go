package CountOnes

import (
	"fmt"
	"strconv"
)

var unos_8bits = make(map[string]int)

// Initialize 8 bit values
func init() {
	var i int64
	for i = 0; i < 256; i++ {
		this_string := fmt.Sprintf("%-8s", strconv.FormatInt(i,2) )
		unos := 0
		for j:= 0; j < 8; j++ {
			if  this_string[j] == '1'  {
				unos++
			}
		}
		unos_8bits[this_string] = unos
	}
}

// Check out https://gobyexample.com/command-line-arguments
func count( cadena string ) int {
	unos := 0
	fmt.Printf( "Cadena longitud %d ", len( cadena ))
	if len( cadena ) > 8  {
		for _, element := range cadena {
			if ( element == '1' ) {
				unos++
			}
		}
	} else {
		unos = unos_8bits[ cadena ]
	}
	return unos
}
