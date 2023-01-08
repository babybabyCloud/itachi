package main

import (
	"os"

	"github.com/urfave/cli/v2"

	check "gitee.com/babybabycloud/itachi/cli"
)

func main() {
	app := &cli.App{
		Name:                 "itachi",
		Usage:                "itachi subcommand [--option] [file...]",
		Version:              "0.0.1",
		Description:          "itachi is a HTTP test framework, you can use it to organize your test plan to a series of test cases.",
		DefaultCommand:       "help",
		EnableBashCompletion: false,
		Authors: []*cli.Author{
			{
				Name:  "babybabycloud",
				Email: "babybabycloud2020@gmail.com",
			},
		},
		Commands: []*cli.Command{
			check.NewCheck(),
		},
	}

	app.Run(os.Args)
}
