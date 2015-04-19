package main

import (
	"log"

	"os"

	"path/filepath"

	"github.com/dynport/dgtk/cli"
	"github.com/jgroeneveld/routes"
)

func main() {
	router := cli.NewRouter()

	router.Register("generate", &generate{}, "generate routes")

	switch err := router.RunWithArgs(); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		log.Fatalf("ERROR: %s", err)
	}
}

type generate struct {
	InFile  string `cli:"arg required"`
	OutFile string `cli:"arg required"`
}

func (g *generate) Run() error {
	in, err := os.Open(g.InFile)
	if err != nil {
		return err
	}
	defer in.Close()

	rd, err := routes.ParseRoutes(in)
	if err != nil {
		return err
	}

	_ = os.Remove(g.OutFile)

	out, err := os.Create(g.OutFile)
	if err != nil {
		return err
	}
	defer out.Close()

	abs, err := filepath.Abs(g.OutFile)
	if err != nil {
		return err
	}

	err = routes.Write(abs, rd, out)
	if err != nil {
		return err
	}

	return nil
}
