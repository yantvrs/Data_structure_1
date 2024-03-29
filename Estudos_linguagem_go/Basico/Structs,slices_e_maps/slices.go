package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 6, 7, 1}

	var s []int = primes[1:4]
	fmt.Println(s)
}

/*

Slices
Uma matriz tem um tamanho fixo. Uma slice, por outro lado, é dinamicamente redimensionada, uma visão flexível dos elementos de uma matriz. Na prática, as slices são muito mais comuns do que as matrizes.

O tipo []T é uma slice com elementos do tipo T.

Uma slice é formada pela especificação de dois indices, um limite menor e maior, separados por dois pontos:

a[low : high]
Este seleciona um intervalo meio-aberto que inclui o primeiro elemento, mas exclui o último.

A expressão a seguir cria uma slice que inclui os elementos de 1 à 3 de a:

a[1:4]
*/
