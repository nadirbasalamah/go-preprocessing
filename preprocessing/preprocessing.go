//Package preprocessing gives simple text preprocessing operation.
package preprocessing

//Preprocessing function returns a preprocessed text
func Preprocessing(xs []string) []string {
	data := xs

	for i := 0; i < len(data); i++ {
		if isSymbolExists(data[i]) {
			data[i] = ""
		}
	}

	for i := 0; i < len(data); i++ {
		if r, f := isDuplicate(data, data[i]); r == true && f > 1 {
			data[i] = ""
		}
	}

	return data
}

//isSymbolExists is a utility function for checking symbol that exists in a string
func isSymbolExists(s string) bool {
	var result bool
	symbols := []string{"`", "~", ",", ".", "?", "!", "/", ";", "&", "'"}

	for i := 0; i < len(symbols); i++ {
		if s == symbols[i] {
			result = true
		}
	}
	return result
}

//isDuplicate is a utility function for checking duplicate words
func isDuplicate(xs []string, s string) (bool, int) {
	var result bool
	var found int
	for i := 0; i < len(xs); i++ {
		if s == xs[i] {
			result = true
			found++
		}
	}
	return result, found
}
