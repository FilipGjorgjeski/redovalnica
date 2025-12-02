package main

import (
	"fmt"
	"os"

	"github.com/FilipGjorgjeski/redovalnica/redovalnica"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "redovalnica",
		Usage: "Upravljanje ocen preko CLI",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "stOcen", Usage: "Minimalno stevilo ocen za pozitivno oceno", Value: 3},
			&cli.IntFlag{Name: "minOcena", Usage: "Najmanjsa mozna ocena", Value: 1},
			&cli.IntFlag{Name: "maxOcena", Usage: "Najvecja mozna ocena", Value: 5},
		},
		Commands: []*cli.Command{
			{
				Name:  "dodaj",
				Usage: "Dodaj oceno (argument: ocena)",
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() < 1 {
						return fmt.Errorf("Podajte oceno kot argument")
					}
					grade := ctx.Args().Get(0)
					// simple parse
					var g int
					_, err := fmt.Sscanf(grade, "%d", &g)
					if err != nil {
						return fmt.Errorf("Ocena mora biti celo stevilo")
					}
					ok := redovalnica.DodajOceno(g, ctx.Int("minOcena"), ctx.Int("maxOcena"))
					if !ok {
						return fmt.Errorf("Ocena izven meja [%d,%d]", ctx.Int("minOcena"), ctx.Int("maxOcena"))
					}
					fmt.Println("Ocena dodana.")
					return nil
				},
			},
			{
				Name:  "izpisi",
				Usage: "Izpisi vse ocene",
				Action: func(ctx *cli.Context) error {
					all := redovalnica.IzpisVsehOcen()
					fmt.Println("Ocene:", all)
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpisi koncni uspeh",
				Action: func(ctx *cli.Context) error {
					msg := redovalnica.IzpisiKoncniUspeh(ctx.Int("stOcen"))
					fmt.Println(msg)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
