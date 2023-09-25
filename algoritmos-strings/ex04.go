//Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str1, str2 string

	//para poder ler as strings
	scanner := bufio.NewScanner(os.Stdin)

	//coletando as strings
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str1 = scanner.Text()

	fmt.Print("Digite outra frase: ")
	scanner.Scan()
	str2 = scanner.Text()

	if str1 == str2 {

		fmt.Print("As strings são iguais")

	} else {

		fmt.Print("As strings são diferentes")

	}

}
