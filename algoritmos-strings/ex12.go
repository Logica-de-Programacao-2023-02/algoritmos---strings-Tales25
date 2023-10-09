//Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(str string) bool {

	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	j := len(str)

	for i := 0; i < len(str); i++ {
		j--
		if str[i] != str[j] {
			return false

		}

	}

	return true

}

func main() {

	var str string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite umsa frase: ")
	scanner.Scan()
	str = scanner.Text()

	if isPalindrome(str) {
		fmt.Print("É um palíndromo")

	} else {
		fmt.Print("Não é um palíndromo")

	}

}
