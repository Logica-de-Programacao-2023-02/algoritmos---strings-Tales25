//Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isDecreasing(str string) bool {

	var slice []string

	//separa a string em strings de caracteres
	slice = strings.Split(str, "")

	for i := 1; i < len(slice); i++ {
		//se o segundo for maior do que o primeiro nao esta crescente
		if slice[i] > slice[i-1] {
			return false

		}

	}

	return true

}

func main() {

	var str string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma sequência numérica: ")
	scanner.Scan()
	str = scanner.Text()

	if isDecreasing(str) {
		fmt.Print("É uma sequência numérica decrescente")

	} else {
		fmt.Print("Não é uma sequência numérica decrescente")

	}

}
