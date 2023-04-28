package cpf

import (
	"fmt"
	"math/rand"
	"strconv"
)

func inicializar() [11]int {
	var cpf [11]int

	for i := 0; i < 9; i++ {
		cpf[i] = rand.Intn(10)
	}

	return cpf
}

func somaFatores(cpf [11]int, limite int) int {

	i := 0
	sum := 0
	for limite >= 2 {
		sum += cpf[i] * limite
		limite--
		i++
	}

	return sum
}

func dv(sum int) int {
	resto := sum % 11

	if resto <= 1 {
		return 0
	}

	return 11 - resto
}

func escreve(cpf [11]int, formatado bool) {

	var cpfOutput string

	for i := 0; i < 11; i++ {
		cpfOutput += strconv.Itoa(cpf[i])
	}

	if formatado {
		cpfOutput = fmt.Sprintf("%s.%s.%s-%s", cpfOutput[0:3], cpfOutput[3:6], cpfOutput[6:9], cpfOutput[9:11])
	}

	fmt.Println(cpfOutput)
}

func GeraCPF(formatado bool) {

	cpf := inicializar()

	cpf[9] = dv(somaFatores(cpf, 10))
	cpf[10] = dv(somaFatores(cpf, 11))

	escreve(cpf, formatado)

}
