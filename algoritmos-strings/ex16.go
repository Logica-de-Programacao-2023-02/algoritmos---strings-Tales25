//Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isSubstring(str1, str2 string) bool {

	if strings.Contains(str1, str2) {
		return true
	}

	return false

}

func main() {

	var str1, str2 string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str1 = scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	str2 = scanner.Text()

	if isSubstring(str1, str2) {
		fmt.Printf("'%s' é substring de '%s'", str2, str1)

	} else {
		fmt.Printf("'%s' não é substring de '%s'", str2, str1)

	}

}
