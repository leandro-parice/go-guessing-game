package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da advinhação!")
	fmt.Println(
		"Um número aleatório será sorteado. \n" +
			"Tente acertar. O número é um inteiro entre 0 e 100.",
	)
	fmt.Println("Boa sorte!")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Printf("Qual o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O valor do chute deve ser um número inteiro.")
			return
		}

		chutes[i] = chuteInt // armazena o chute no array

		switch {
		case chuteInt < x:
			fmt.Println("Errado! O seu chute é menor que o número sorteado.")
		case chuteInt > x:
			fmt.Println("Errado! O seu chute é maior que o número sorteado.")
		default:
			fmt.Printf("Parabéns! Você acertou o número sorteado. "+
				"Suas tentativas foram: %v.", chutes[:i+1])
			return
		}
	}

	fmt.Printf("Suas tentativas foram: %v. O número sorteado era %d. Tente novamente!", chutes, x)
}
