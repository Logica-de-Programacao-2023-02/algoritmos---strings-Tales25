//Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var str, currentChar, newChar string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	fmt.Print("Escolha uma letra atual: ")
	scanner.Scan()
	currentChar = scanner.Text()

	fmt.Print("Escolha uma nova letra: ")
	scanner.Scan()
	newChar = scanner.Text()

	str = strings.ReplaceAll(str, currentChar, newChar)

	fmt.Printf("A nova frase ficou:\n%s", str)

}
