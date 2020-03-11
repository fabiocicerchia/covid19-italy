# COVID-19 ITALY - TRENDS

[![Pull Requests](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?longCache=true)](https://github.com/fabiocicerchia/covid19-italy/pulls)
[![MIT License](https://img.shields.io/badge/License-MIT-lightgrey.svg?longCache=true)](LICENSE)

**What's COVID-19?**

> Coronavirus disease 2019 (COVID-19) is an infectious disease caused by severe acute respiratory syndrome coronavirus 2
(SARS-CoV-2),[9] a virus closely related to the SARS virus.
> https://en.wikipedia.org/wiki/Coronavirus_disease_2019

**How can I help prevent the spreading?**

Since the virus is quite aggressive and the contamination trends are exponential, the recommended way to prevent from
spreading even more is: **STAYING AT HOME**.

This is not an official source so please do get informed with your local authorities in regards to what measures have
been put in place to contain the virus.

This repo is based on the official data released by the [Protezione Civile](https://github.com/pcm-dpc) under the
[CC-BY-4.0 license](https://github.com/pcm-dpc/COVID-19/blob/master/LICENSE).

## Requirements

 - PHP 7.0+

## Install

```
git submodule init && git submodule update
composer install
```

## Run

```
$ ./trend.php --help
COVID-19 ITALY - TRENDS
 ./trend.php [--province x|--region y] [--last-week|--last-month]
```

```
$ ./trend.php --province roma
TOTAL CASES IN ROMA
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
$ ./trend.php --region lazio
TOTAL CASES IN LAZIO
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

## CONTRIBUTE

Any contribution is very welcomed!

**NOTE:** I've tried to be as much practical as possible, without going too deep with design-patterns and best-practices.
Although I'll refactor asap as much as possible.

## TODO

 - Cleaning Code (OOP, etc)
 - Exporting data
 - API endpoint
 - One-liner installer
 - Packagist (?)
 - Docker (?)
 - CLI colours
