//Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func changeVowels(str string) string {

	str = strings.ToLower(str)

	str = strings.ReplaceAll(str, "a", "*")
	str = strings.ReplaceAll(str, "e", "*")
	str = strings.ReplaceAll(str, "i", "*")
	str = strings.ReplaceAll(str, "o", "*")
	str = strings.ReplaceAll(str, "u", "*")

	return str

}

func main() {

	var str, newStr string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	newStr = changeVowels(str)

	fmt.Printf("Revomendo as vogais, temos:\n%s", newStr)

}
