package app

import (
	"fmt"
	"jonasricardo/brDocGen/cpf"

	"github.com/urfave/cli/v2"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "brDocGen"
	app.Usage = "Gerador de documentos brasileiros válidos"

	app.Commands = []*cli.Command{
		{
			Name:  "cpf",
			Usage: "Gera CPFs",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "formatado",
					Usage:   "retorna o CPF formatado com . e -",
					Aliases: []string{"f"},
				},
				&cli.IntFlag{
					Name:    "quantidade",
					Usage:   "Quantidade de documentos retornados",
					Aliases: []string{"q"},
					Value:   1,
				},
			},
			Action: geraCpf,
		},
		{
			Name:  "cnpj",
			Usage: "Gera CNPJs",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "formatado",
					Usage:   "retorna o CNPJ formatado com ./-",
					Aliases: []string{"f"},
				},
				&cli.IntFlag{
					Name:    "quantidade",
					Usage:   "Quantidade de documentos retornados",
					Aliases: []string{"q"},
					Value:   1,
				},
			},
			Action: teste,
		},
	}

	return app
}

func geraCpf(c *cli.Context) error {

	for i := 0; i < c.Int("quantidade"); i++ {
		cpf.GeraCPF(c.Bool("formatado"))
	}

	return nil
}

func teste(c *cli.Context) error {
	fmt.Println(c.Bool("formatado"))

	return nil

}
