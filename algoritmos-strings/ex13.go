//Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isGrowing(str string) bool {

	var slice []string

	//separa a string em strings de caracteres
	slice = strings.Split(str, "")

	//ordena em ordem alfabetica e numerica pelo visto, MAS SE EU ORDENAR SEMPRE VAI ESTAR CRESCENTE
	//sort.Strings(slice)

	for i := 1; i < len(slice); i++ {
		//se o segundo for menor do que o primeiro nao esta crescente
		if slice[i] < slice[i-1] {
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

	if isGrowing(str) {
		fmt.Print("É uma sequência numérica crescente")

	} else {
		fmt.Print("Não é uma sequência numérica crscente")

	}

}
