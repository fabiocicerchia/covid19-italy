package main

import (
    "log"
    "os"

    getoptshandler "cli/pkg/utils/getoptshandler"
    reportfactory "cli/pkg/factory"
    reportformatter "cli/pkg/utils/reportformatter"

    "github.com/urfave/cli/v2"
)

func main() {
    var province string
    var region string
    var lastWeek bool = false
    var lastMonth bool = false

    app := &cli.App{
        Name: "covid19-trend",
        Usage: "COVID-19 ITALY - TRENDS",
        Flags: []cli.Flag{
            &cli.StringFlag{
              Name:        "region",
              Usage:       "Show the trends for the selected region",
              Destination: &region,

            },
            &cli.StringFlag{
              Name:        "province",
              Usage:       "Show the trends for the selected province",
              Destination: &province,

            },
            &cli.BoolFlag{
              Name:        "last-week",
              Usage:       "Show the trends for the last week",
              Destination: &lastWeek,
            },
            &cli.BoolFlag{
              Name:        "last-month",
              Usage:       "Show the trends for the last month",
              Destination: &lastMonth,
            },
        },
        Action: func(c *cli.Context) error {
            o := &getoptshandler.GetoptsHandler{}
            o.HandleOptions(region, province, lastWeek, lastMonth)

            service := reportfactory.Create(o.GetOptions())
            service.ProcessData()

            reportformatter.Output(service.GetUserChoice().GetSelectedValue(), service.GetData())

            return nil
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
