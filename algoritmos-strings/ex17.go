//Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func findUniqueLetters(str string) string {
	var uniqueLetters string
	var isUnique bool

	//inicio do loop, inicalmente presumimos que o caractere é unico
	for i := 0; i < len(str); i++ {
		isUnique = true

		for j := 0; j < len(str); j++ {
			//I!=J => para garantir uma comparacao de caracteres em posicoes diferentes (sem ser ele com ele mesmo)
			//STR[I]==STR[J] => verifica se o caractere da posicao I é igual ao da J
			//SE i!=j E str[i]==str[j]; então o caractere aparece mais de uma vez...
			if i != j && str[i] == str[j] {
				//...logo, isUnique recebe falso e quebro o loop
				isUnique = false
				break
			}
		}

		//se isUnique for true (só houver 1 caractere desse na string)
		if isUnique {
			//adiciono o caractere da posição I na string uniqueLetters (inicialmente vazia)
			uniqueLetters += string(str[i])
		}
	}

	return uniqueLetters
}

func main() {
	var str, uniqueLetters string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	uniqueLetters = findUniqueLetters(str)
	fmt.Printf("As letras únicas da frase são:\n%s", uniqueLetters)
}
