package questoes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func PrimeiraQuestao() {
	println("PRIMEIRA QUESTAO:")
	fmt.Print("Informe um numero: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	if isFibonacci(num) {
		fmt.Print(strconv.Itoa(num) + " e um fibonacci")
	} else {
		fmt.Print(strconv.Itoa(num) + " nao e um fibonacci")
	}
}

func SegundaQuestao() {
	println("SEGUNDA QUESTAO:")
	fmt.Print("Informe uma string: ")
	input, _ := reader.ReadString('\n')
	println(CountA(input))
}

func TerceiraQuestao() {
	println("TERCEIRA QUESTAO:")
	println("o valor da soma da terceira questao e: " + strconv.Itoa(TerceiraAlgoritmo()))
}

func QuartaQuestao() {
	println("QUARTA QUESTAO:")
	println("a = 9, numeros primos")
	println("b = 128, potencia de 2")
	println("c = 49, quadrado dos numeros em ordem crescente")
	println("d = 100, 4n^2")
	println("e = 13, sequencia de Fibonacci")
	println("f = 200, numeros que comecam com D")
}

func QuintaQuestao() {
	print("ligo o primeiro interruptor, e depois de um tempo desligo ele, entao ligo o segundo,\n" +
		"e vou para a primeira sala, se a lampada estiver quente, ela pertence ao primeiro interruptor,\n" +
		"se estiver acessa, pertence ao segundo, e se estiver apagada pertence ao terceiro." +
		"supondo que ela esteja apagada e fria, vou entao para a segunda sala, se a luz estiver acesa,\n" +
		"pertence ao segundo interruptor, se estiver apagada, pertence ao primeiro. supondo que ela esteja\n" +
		"apagada, entao ela pertence ao primeiro, e por eliminação, a terceira sala pertencera ao segundo interruptor")
}
