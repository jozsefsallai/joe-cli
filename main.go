package main

import (
	"log"
	"os"

	"github.com/jozsefsallai/joe-cli/commands"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func meta() {
	app.Name = "joe"
	app.Usage = "Joe's own CLI for everyday tasks."
	app.Author = "@jozsefsallai"
	app.Version = "0.0.1"
}

func setup() {
	app.Commands = []cli.Command{
		commands.UploadCommand,
	}
}

func main() {
	meta()
	setup()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
