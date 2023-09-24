//Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
)

//JEITO NA LÓGICA DA TABELA ASC (TALVEZ DESSE ERRO COM CARACTERES ESPECIAIS, FUNCIONA EM FRASES SIMPLES)

func main() {

	var str string
	var isUpper bool
	var i int

	//para poder ler uma string
	scanner := bufio.NewScanner(os.Stdin)

	//coletando a string digitada
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	//atribuindo a string digitada para a var str
	str = scanner.Text()

	//loop para verificar se é maiusculo ou minisculo caractere por caractere
	for i = 0; i < len(str); i++ {

		//o carctere comeca como maiusculo (isUpper verdadeiro)
		isUpper = true

		//o valor do caractere segundo a tabela ascii for maior do que 96 (letras minusculas)...
		if str[i] > 96 {

			//o caractere nao esta maiusculo (isUpper falso)
			isUpper = false

		}

		//se for maiscula (isUpper verdadeiro)
		if isUpper {

			//printe a string que tenha o valor str[i] segundo a tabela ascii
			fmt.Print(string(str[i]))

			//se nao for maiuscula (isUpper falso)
		} else {

			//printe a string que tenha o valor str[i]-32 (para transformar a letra maiuscula em minuscula)...
			//segundo a tabela ascii
			fmt.Print(string(str[i] - 32))

		}

	}

}

##################################################################################################################################
//MELHOR MANEIRA

func main() {

	var str string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	fmt.Print("A frase, em caixa alta, ficou: ", strings.ToUpper(str))

}
