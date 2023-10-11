//Solicite ao usuário uma string e imprima a mesma string invertida.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str, reversedStr string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	reversedStr = reverseString(str)
	fmt.Printf("A frase invertida ficou:\n%s", reversedStr)
}

func reverseString(str string) string {
	var reversedStr string
	reversedStr = ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	return reversedStr
}
