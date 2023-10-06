//Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str string

	scanner := bufio.NewScanner(os.Stdin)
	wordsCounter := 1

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	for i := 0; i < len(str); i++ {

		if string(str[i]) == " " {
			wordsCounter++

		}

	}

	fmt.Printf("A frase contém %d palavras", wordsCounter)

}
