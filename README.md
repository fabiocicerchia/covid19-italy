# COVID-19 SPREADING IN ITALY

[![Pull Requests](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?longCache=true)](https://github.com/fabiocicerchia/covid19-italy/pulls)
[![MIT License](https://img.shields.io/badge/License-MIT-lightgrey.svg?longCache=true)](LICENSE)

**What's COVID-19?**

> Coronavirus disease 2019 (COVID-19) is an infectious disease caused by severe acute respiratory syndrome coronavirus 2
> (SARS-CoV-2). The disease was first identified in 2019 in Wuhan, China, and has spread globally, resulting in the 2019–20
> coronavirus pandemic. Common symptoms include fever, cough and shortness of breath. Muscle pain, sputum production and sore
> throat are less common symptoms.
> https://en.wikipedia.org/wiki/Coronavirus_disease_2019

**How can I help prevent the spreading?**

Since the virus is quite aggressive and the contamination trends are exponential, the recommended way to prevent from
spreading even more is: **[#StayTheFuckHome](https://covid19.fabiocicerchia.it/)**.

This is not an official source so please do get informed with your local authorities in regards to what measures have
been put in place to contain the virus. More details on the
[Coronavirus disease (COVID-19) outbreak](https://www.who.int/emergencies/diseases/novel-coronavirus-2019).

**Want lots of charts?**
[Coronavirus Italia - Google Data Studio](https://datastudio.google.com/reporting/91350339-2c97-49b5-92b8-965996530f00) by [FMossotto](https://twitter.com/FMossotto).

This repo is based on the official data released by the [Protezione Civile](https://github.com/pcm-dpc) under the
[CC-BY-4.0 license](https://github.com/pcm-dpc/COVID-19/blob/master/LICENSE) and the geography data by
[Openpolis](https://github.com/openpolis).

![Screenshot](/screenshot.png)

## Requirements

 - PHP 7.0+

## Install

```
$ ./bin/setup.sh
```

## Run

```
$ go run main.go --help
NAME:
   covid19-trend - COVID-19 ITALY - TRENDS

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value    Show the trends for the selected region
   --province value  Show the trends for the selected province
   --last-week       Show the trends for the last week (default: false)
   --last-month      Show the trends for the last month (default: false)
   --help, -h        show help (default: false)
```

```
$ go run main.go --province roma
COVID-19 TOTAL CASES IN ROMA
2020-02-24:  0
2020-02-25:  3
2020-02-26:  3 ->   +0%
2020-02-27:  3 ->   +0%
2020-02-28:  3 ->   +0%
2020-02-29:  6 -> +100%
2020-03-01:  6 ->   +0%
2020-03-02:  7 ->  +16%
2020-03-03: 14 -> +100%
2020-03-04: 29 -> +107%
2020-03-05: 42 ->  +44%
2020-03-06: 49 ->  +16%
2020-03-07: 71 ->  +44%
2020-03-08: 77 ->   +8%
2020-03-09: 91 ->  +18%
2020-03-10: 76 ->  -16%
```

```
$ go run main.go --region lazio
COVID-19 TOTAL CASES IN LAZIO
2020-02-24:   3
2020-02-25:   3 ->   +0%
2020-02-26:   3 ->   +0%
2020-02-27:   3 ->   +0%
2020-02-28:   3 ->   +0%
2020-02-29:   6 -> +100%
2020-03-01:   6 ->   +0%
2020-03-02:   7 ->  +16%
2020-03-03:  14 -> +100%
2020-03-04:  30 -> +114%
2020-03-05:  44 ->  +46%
2020-03-06:  54 ->  +22%
2020-03-07:  76 ->  +40%
2020-03-08:  87 ->  +14%
2020-03-09: 102 ->  +17%
2020-03-10: 116 ->  +13%
```

## Run web interface

```
$ php -S 127.0.0.1:8001 -t web
```

Then you can browse application on `http://127.0.0.1:8001/` (you can use a different port, if you prefer)

## CONTRIBUTE

Any contribution is very welcomed!

**NOTE:** I've tried to be as much practical as possible, without going too deep with design-patterns and best-practices.
At this stage I mainly focus on the functionalities, although I do refactor as often as possible.
