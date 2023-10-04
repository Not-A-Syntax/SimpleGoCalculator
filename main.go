package main

// Imports

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Map contendo as Strings dos Fatores

var operadores = map[string]bool{
	"*":    true,
	"/":    true,
	"+":    true,
	"-":    true,
	"^":    true,
	"log":  true,
	"root": true,
	"%":    true,
	"!":    true,
}

// Var para uso na fatoração

var s1 float64 = 1

// Função para Print

func Print(s string, args ...any) {
	fmt.Printf(s, args...)
}

// Função para leitura de um Input

func ReadInput(pergunta string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(pergunta + " ")
	readinput, _ := reader.ReadString('\n')
	readinput = strings.TrimSpace(readinput)
	return readinput
}

// Conversor de String para Float64

func ConvFloat64(input string) float64 {
	convinput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		Print("[ERROR] Input inválido")
	}
	return convinput
}

// Função Principal para fazer os calculos

func main() {

	// Var criada para atribuir e Printar o Resultado Futuramente.

	var resultado float64

	// Input para saber qual fator da operação a pessoa quer.

	fator := ReadInput("Qual o Fator da Operação?\n( '+' | '-' | '*' | '/' | '^' | 'log' | 'root' | '%' | '!'):")

	// If para verificar se é um fator válido. caso contrário mande uma mensagem de erro e interrompe o programa.

	if !operadores[fator] {
		Print("[ERROR] O fator '%s' não é válido para esta operação!", fator)
		return
	}

	// Input do Primeiro valor
	n1 := ConvFloat64(ReadInput("Primeiro Produto:"))

	// Switch para caso o primeiro valor seja fator ele não pedir um segundo valor.

	switch fator {
	case "!":

		// For em loop para fazer o Fatorial
		for i := 1; i < int(n1)+1; i++ {
			i2 := float64(i)
			s1 = s1 * i2
			resultado = s1
		}
		Print("O Resultado é: %s\n", fmt.Sprint(resultado))
		return
	}

	// Input do Segundo Valor

	n2 := ConvFloat64(ReadInput("Segundo Produto:"))

	// Switch para verificação de cada fator e calcular de acordo com o mesmo.

	switch fator {
	case "*":
		resultado = n1 * n2
	case "/":
		resultado = n1 / n2
	case "+":
		resultado = n1 + n2
	case "-":
		resultado = n1 - n2
	case "^":
		resultado = math.Pow(n1, n2)
	case "log":
		resultado = math.Log(n1) / math.Log(n2)
	case "root":
		resultado = math.Pow(n1, 1/n2)
	case "%":
		resultado = n1 / (100 * n2)
	default:
		Print("Argumentos Inválidos.")
	}

	// Print final para o Resultado da operação.

	Print("O Resultado é: %s\n", fmt.Sprint(resultado))

}
