/* CountOnes - provides different ways of counting ones in a string
   with 0s and ones 
*/

package CountOnes

import (
	"fmt"
	"strconv"
)

var unos_8bits = make(map[string]int)

// Initialize 8 bit values for caching
func init() {
	var i int64
	for i = 0; i < 256; i++ {
		this_string := fmt.Sprintf("%08s", strconv.FormatInt(i,2) )
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
func Count( cadena string ) int {
	unos := 0
//	fmt.Printf( "Cadena longitud %d ", len( cadena ))
	if len( cadena ) > 8  {
		for _, element := range cadena {
			if ( element == '1' ) {
				unos++
			}
		}
	} else {
		if len(cadena) < 8 {
			cadena = fmt.Sprintf("%08s", cadena )
		}
//		fmt.Println( cadena )
		unos = unos_8bits[ cadena ]
	}
	return unos
}

// Valid only for multiples of 8, counts the number of ones
// Internal for the time being. 
func recursive_count( cadena string ) int {
	if len(cadena) <= 8 {
		return unos_8bits[ cadena ]
	}
	return recursive_count( cadena[:len(cadena)/2] ) + recursive_count(  cadena[len(cadena)/2:] )
}
