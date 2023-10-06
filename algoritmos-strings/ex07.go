//Solicite ao usuário uma string e informe se ela contém pelo menos um número.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var str string
	var isNum bool

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	for i := 0; i < len(str); i++ {

		if str[i] >= 48 && str[i] <= 57 {
			isNum = true
			break

		}

	}

	if isNum {
		fmt.Print("Há pelo menos um número na frase")

	} else {
		fmt.Print("Não há números na frase")

	}

}
