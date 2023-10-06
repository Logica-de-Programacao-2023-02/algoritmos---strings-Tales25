//Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeVowels(str string) string {

	str = strings.ToLower(str)

	str = strings.ReplaceAll(str, "a", "")
	str = strings.ReplaceAll(str, "e", "")
	str = strings.ReplaceAll(str, "i", "")
	str = strings.ReplaceAll(str, "o", "")
	str = strings.ReplaceAll(str, "u", "")

	return str

}

func main() {

	var str, newStr string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	newStr = removeVowels(str)

	fmt.Printf("Revomendo as vogais, temos:\n%s", newStr)

}
