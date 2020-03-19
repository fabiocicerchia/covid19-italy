package clisetup

import (
	"strconv"

	"github.com/urfave/cli/v2"

	reportfactory "cli/pkg/factory/reportfactory"
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
			err := o.HandleOptions(c.String("region"), c.String("province"), c.Bool("last-month"), c.Bool("last-week"))
			if err != nil {
				panic(err)
			}

			options := o.GetOptions()
			lastDays, _ := strconv.Atoi(options["lastDays"])
			service := reportfactory.Create(options["file"], options["type"], options["userChoice"], lastDays)
			service.ProcessData()

			reportformatter.Output(service.GetUserChoice().GetSelectedValue(), service.GetData())

			return nil
		},
	}
}
