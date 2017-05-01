package CountOnes

// Check out https://gobyexample.com/command-line-arguments
func count( cadena string) int {
	unos := 0
	for _, element := range cadena {
		if ( element == '1' ) {
			unos++
		} 
	}
	return unos
}
