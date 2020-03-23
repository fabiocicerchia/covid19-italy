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

 - GO 1.14
 - PHP 7.0+

## Install

```
$ ./bin/setup.sh
```

## Run

```
$ ./bin/covid19-trend --help
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
$ ./bin/covid19-trend --province roma
COVID-19 NEW CASES IN ROMA

DATE            NEW_CASES       TOTAL_CASES     % INCREASE
2020-02-24         0               0
2020-02-25         3               3
2020-02-26         0               3               0%
2020-02-27         0               3               0%
2020-02-28         0               3               0%
2020-02-29         3               6            +100%
2020-03-01         0               6               0%
2020-03-02         1               7             +16%
2020-03-03         7              14            +100%
2020-03-04        15              29            +107%
2020-03-05        13              42             +44%
2020-03-06         7              49             +16%
2020-03-07        22              71             +44%
2020-03-08         6              77              +8%
2020-03-09        14              91             +18%
2020-03-10       -15              76             -17%
2020-03-11        23              99             +30%
2020-03-12        63             162             +63%
2020-03-13        56             218             +34%
2020-03-14        70             288             +32%
2020-03-15        66             354             +22%
2020-03-16        58             412             +16%
2020-03-17        74             486             +17%
2020-03-18       104             590             +21%
2020-03-19        88             678             +14%
2020-03-20        77             755             +11%
2020-03-21       138             893             +18%
2020-03-22       156            1049             +17% WORST PEAK

TOTAL CASES: 1049
```

```
$ ./bin/covid19-trend --region lazio
COVID-19 NEW CASES IN LAZIO

DATE            NEW_CASES       TOTAL_CASES     % INCREASE
2020-02-24         3               3
2020-02-25         0               3               0%
2020-02-26         0               3               0%
2020-02-27         0               3               0%
2020-02-28         0               3               0%
2020-02-29         3               6            +100%
2020-03-01         0               6               0%
2020-03-02         1               7             +16%
2020-03-03         7              14            +100%
2020-03-04        16              30            +114%
2020-03-05        14              44             +46%
2020-03-06        10              54             +22%
2020-03-07        22              76             +40%
2020-03-08        11              87             +14%
2020-03-09        15             102             +17%
2020-03-10        14             116             +13%
2020-03-11        34             150             +29%
2020-03-12        50             200             +33%
2020-03-13        77             277             +38%
2020-03-14        80             357             +28%
2020-03-15        79             436             +22%
2020-03-16        87             523             +19%
2020-03-17        84             607             +16%
2020-03-18       117             724             +19%
2020-03-19        99             823             +13%
2020-03-20       185            1008             +22%
2020-03-21       182            1190             +18%
2020-03-22       193            1383             +16% WORST PEAK

TOTAL CASES: 1383
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
