//Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite qualquer coisa: ")
	scanner.Scan()
	str = scanner.Text()

	if containsOnlyNumbers(str) {
		fmt.Print("Há apenas números na frase")
	} else {
		fmt.Print("Não há apenas números na frase")
	}
}

func containsOnlyNumbers(str string) bool {
	for _, char := range str {
		//aspas duplas comparam com string, duplas comparam com alguma coisa de rune
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
