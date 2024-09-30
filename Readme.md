### 1) Dado a sequência de Fibonacci, onde se inicia por 0 e 1 e o próximo valor sempre será a soma dos 2 valores anteriores (exemplo: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...), escreva um programa na linguagem que desejar onde, informado um número, ele calcule a sequência de Fibonacci e retorne uma mensagem avisando se o número informado pertence ou não a sequência.

IMPORTANTE: Esse número pode ser informado através de qualquer entrada de sua preferência ou pode ser previamente definido no código;

#### [resposta](questoes/primeiraFibonacci.go)

### 2) Escreva um programa que verifique, em uma string, a existência da letra ‘a’, seja maiúscula ou minúscula, além de informar a quantidade de vezes em que ela ocorre.

IMPORTANTE: Essa string pode ser informada através de qualquer entrada de sua preferência ou pode ser previamente definida no código;

#### [resposta](questoes/segundaContarA.go)

### 3) Observe o trecho de código abaixo: int INDICE = 12, SOMA = 0, K = 1; enquanto K < INDICE faça { K = K + 1; SOMA = SOMA + K; } imprimir(SOMA);

### Ao final do processamento, qual será o valor da variável SOMA?

####resposta: 
66
[calculo](questoes/terceiraCalculo.go)

### 4) Descubra a lógica e complete o próximo elemento:
#### a) 1, 3, 5, 7, ___      
#### b) 2, 4, 8, 16, 32, 64, ____
#### c) 0, 1, 4, 9, 16, 25, 36, ____
#### d) 4, 16, 36, 64, ____
#### e) 1, 1, 2, 3, 5, 8, ____
#### f) 2,10, 12, 16, 17, 18, 19, ____


#### respostas:
a = 9, numeros primos

b = 128, potencia de 2

c = 49, quadrado dos numeros em ordem crescente

d = 100, 4n²

e = 13, sequencia de Fibonacci

f = 200, numeros que comecam com "D"


### 5) Você está em uma sala com três interruptores, cada um conectado a uma lâmpada em salas diferentes. Você não pode ver as lâmpadas da sala em que está, mas pode ligar e desligar os interruptores quantas vezes quiser. Seu objetivo é descobrir qual interruptor controla qual lâmpada. Como você faria para descobrir, usando apenas duas idas até uma das salas das lâmpadas, qual interruptor controla cada lâmpada? 

#### resposta:
ligo o primeiro interruptor, e depois de um tempo desligo ele, entao ligo o segundo, e vou para a primeira sala, se a lampada estiver quente, ela pertence ao primeiro interruptor, se estiver acessa, pertence ao segundo, e se estiver apagada pertence ao terceiro. supondo que ela esteja apagada e fria, vou entao para a segunda sala, se a luz estiver acesa, pertence ao segundo interruptor, se estiver apagada, pertence ao primeiro. supondo que ela esteja apagada, entao ela pertence ao primeiro, e por eliminação, a terceira sala pertencera ao segundo interruptor