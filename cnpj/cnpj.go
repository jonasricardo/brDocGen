package cnpj

import (
	"fmt"
	"math/rand"
	"strconv"
)

func inicializar() [14]int {
	var cnpj [14]int

	for i := 0; i < 14; i++ {
		cnpj[i] = rand.Intn(10)
	}

	return cnpj
}

// dv - dígito 1 ou 2
func somaFatores(cpf [14]int, dv int) int {

	i := 0
	sum := 0
	limite := 4 + dv
	for limite >= 2 {
		sum += cpf[i] * limite
		limite--
		i++
	}

	limite = 9
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

func escreve(cnpj [14]int, formatado bool) {

	var cnpjOutput string

	for i := 0; i < 14; i++ {
		cnpjOutput += strconv.Itoa(cnpj[i])
	}

	if formatado {
		cnpjOutput = fmt.Sprintf("%s.%s.%s/%s-%s", cnpjOutput[0:2], cnpjOutput[2:5], cnpjOutput[5:8], cnpjOutput[8:12], cnpjOutput[12:14])
	}

	fmt.Println(cnpjOutput)
}

func GeraCNPJ(formatado bool) {

	cnpj := inicializar()

	cnpj[12] = dv(somaFatores(cnpj, 1))
	cnpj[13] = dv(somaFatores(cnpj, 2))

	escreve(cnpj, formatado)

}
