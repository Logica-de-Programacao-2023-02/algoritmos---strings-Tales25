//Solicite ao usuário uma string e informe se ela é está em camelCase e quantas palavras possuí. CamelCase é quando a
//primeira letra de cada palavra é maiúscula, e as demais são minúsculas. Exemplo: "estaStringEstaEmCamelCase".

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	var words int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	//isCamel recebe o valor do primeiro item (bool) do return e words do segundo item (int)
	isCamel, words := isCamelCase(str)

	if isCamel {
		fmt.Printf("Está em Camel Case;\nPossui %d palavras;", words)
	} else {
		fmt.Print("Não está em Camel Case;")
	}
}

func isCamelCase(str string) (bool, int) {
	var words int
	words = 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && str[i+1] > 'Z' || str[0] > 'Z' {
			return false, 0
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			words++
		}
	}
	return true, words
}
