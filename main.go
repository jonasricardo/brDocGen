package main

import (
	"jonasricardo/brDocGen/app"
	"log"
	"os"
)

func main() {

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}
