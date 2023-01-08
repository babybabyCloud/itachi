package cli

import (
	"os"

	"gitee.com/babybabycloud/itachi/configuration"
	"github.com/urfave/cli/v2"
)

func NewCheck() *cli.Command {
	check := cli.Command{
		Name:        "check",
		Usage:       "itachi check [file...]",
		Description: "Check if there is any error in the configuration files",
		Action:      CheckAction,
	}
	return &check
}

// CheckAction check the configuration files
func CheckAction(ctx *cli.Context) error {
	for _, fileName := range ctx.Args().Slice() {
		data, err := os.ReadFile(fileName)
		if err != nil {
			// TODO handle this error
			return err
		}

		config := configuration.NewConfiguration(data)

		check(config)
	}

	return nil
}

func check(config *configuration.Configuration) error {
	config.Start()
	return nil
}
