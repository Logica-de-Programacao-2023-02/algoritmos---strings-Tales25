//Peça ao usuário para digitar uma string e retorne a mesma string com as letras em
//maiúsculo.

package main

import (
	"bufio"
	"fmt"
	"os"
)

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
