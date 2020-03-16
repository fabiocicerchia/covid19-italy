package clisetup

import (
	"github.com/urfave/cli/v2"

	reportfactory "cli/pkg/factory"
	getoptshandler "cli/pkg/utils/getoptshandler"
	reportformatter "cli/pkg/utils/reportformatter"
)

func Config() cli.App {
	return cli.App{
		Name:  "covid19-trend",
		Usage: "COVID-19 ITALY - TRENDS",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "region",
				Usage: "Show the trends for the selected region",
			},
			&cli.StringFlag{
				Name:  "province",
				Usage: "Show the trends for the selected province",
			},
			&cli.BoolFlag{
				Name:  "last-week",
				Usage: "Show the trends for the last week",
			},
			&cli.BoolFlag{
				Name:  "last-month",
				Usage: "Show the trends for the last month",
			},
		},
		Action: func(c *cli.Context) error {
			o := new(getoptshandler.GetoptsHandler)
			err := o.HandleOptions(*c)
			if err != nil {
				panic(err)
			}

			service := reportfactory.Create(o.GetOptions())
			service.ProcessData()

			reportformatter.Output(service.GetUserChoice().GetSelectedValue(), service.GetData())

			return nil
		},
	}
}
