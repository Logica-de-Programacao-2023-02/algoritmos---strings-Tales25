//Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	//i recebe o tamanho menos um porque em WORLD o tamanho é 5 com indices de 0 a 4 (5 numeros; 0, 1, 2, 3, 4)
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))

	}

}
