//Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	var word1, word2 string
	var areAnagrams bool
	var slice1 []string
	var slice2 []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma palavra: ")
	scanner.Scan()
	word1 = scanner.Text()

	fmt.Print("Digite outra palavra: ")
	scanner.Scan()
	word2 = scanner.Text()

	//para comparar todas as letras em minusculo e sem espaços
	word1 = strings.ToLower(strings.ReplaceAll(word1, " ", ""))
	word2 = strings.ToLower(strings.ReplaceAll(word2, " ", ""))

	if len(word1) != len(word2) {
		areAnagrams = false
		goto final
	}

	//para criar slices de string, onde cada caractere vira uma string
	//CHATGPT - Converte as strings em slices de caracteres
	slice1 = strings.Split(word1, "")
	slice2 = strings.Split(word2, "")

	//para botar o slice em ordem alfabetica
	//CHATGPT - Ordena os slices
	sort.Strings(slice1)
	sort.Strings(slice2)

	//CHATGPT - strings.Join é usado para concatenar os elementos de um slice de strings em uma única string
	word1 = strings.Join(slice1, "")
	word2 = strings.Join(slice2, "")
	
	if word1 == word2 {
		areAnagrams = true
	}

final:
	if areAnagrams {
		fmt.Print("São anagramas")
	} else {
		fmt.Print("Não são anagramas")
	}

}
