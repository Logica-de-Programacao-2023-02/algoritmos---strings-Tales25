//Escreva um programa que receba uma string e remova todos os espaços em branco.
//Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//JEITO LÓGICO - TABELA ASCII (POR ALGUM MOTIVO NÃO FUNCIONA COM ACENTOS E NÃO ME IMPORTO)

/*func main() {

	var str string
	var isSpace bool

	//para poder ler uma string
	scanner := bufio.NewScanner(os.Stdin)

	//lendo de fato a string e atribuindo ela a variavel str
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	//loop para percorrer os indices de todos os caracteres
	for i := 0; i < len(str); i++ {

		isSpace = false

		//se o valor do caractere segundo a tabela ascii for igual a 32 (valor do espaço)...
		if str[i] == 32 {

			isSpace = true

		}

		//se não for espaço (isSpace = false; !isSpace)
		if !isSpace {

			//printe a string (letra/caractere) de identificação str[i] normalmente (pq não é um espaço)
			fmt.Print(string(str[i]))

			//caso o contrario (se for espaço; isSpace = true)
		} else {

			//printe a string referente ao valor str[i]-32 (pq é um espaço e substraindo 32 equivale a nulo)
			fmt.Print(string(str[i] - 32))

		}

	}

}*/

//MELHOR JEITO USANDO FUNÇÕES/COMANDOS

func main() {

	var str string

	//para poder ler uma string
	scanner := bufio.NewScanner(os.Stdin)

	//lendo de fato a string e atribuindo ela a variavel str
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	//a função replace pega a string deseja (str), o que você quer substituir (" ", espaços em branco),...
	//o que você usara como substituto ("", nada, nulo, "espaço vazio") e a quantidade de vezes que...
	//isso será feito (-1, um negativo faz infinitamente, um positivo faria apenas no primeiro espaço)
	str = strings.Replace(str, " ", "", -1)

	//print da string sem espaços
	fmt.Print(str)

}
