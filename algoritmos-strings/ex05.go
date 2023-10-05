//Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
//Informe ao usuário se é ou não.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var str string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite qualquer número: ")
	scanner.Scan()
	str = scanner.Text()

	//preciso criar outra variável (i) para, caso seja um float, receber o novo valor como float
	//e err como uma variavel de error
	i, err := strconv.ParseFloat(str, 64)

	//caso o erro seja diferente de nulo, no caso se o erro existir não é float
	if err != nil {
		//uso a str pq o i não existe caso o erro exista
		fmt.Print(str, " não é válido em ponto flutuante")

	} else {
		fmt.Print(i, " é válido em ponto flutuante")

	}

}
