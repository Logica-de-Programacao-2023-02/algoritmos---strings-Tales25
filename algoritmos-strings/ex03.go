//Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro
//caractere especificado pelo usuário. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var phrase, oldChar, newChar string

	//para poder ler uma string como uma string
	scanner := bufio.NewScanner(os.Stdin)

	//lendo, de fato, a string
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	phrase = scanner.Text() //atribuindo a frase lida a variavel phrase

	//lendo o caractere a ser substituido
	fmt.Print("Escolha um caractere: ")
	scanner.Scan()
	oldChar = scanner.Text() //atribuindo o caractere a variavel oldChar

	//lendo o caractere que vai substitui-lo
	fmt.Print("Escolha um caractere: ")
	scanner.Scan()
	newChar = scanner.Text() //atribuindo o caractere a variavel newChar

	//parecido com Replace, ReplaceAll substitui todas as ocorrencias de um determinado caractere
	//(string a ser alterada, oq vc quiser substituir, oq irá substituir)
	phrase = strings.ReplaceAll(phrase, oldChar, newChar)

	fmt.Printf("A frase após a substituição ficou:\n%s", phrase)

}
