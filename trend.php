#!/usr/bin/env php
<?php

use League\Csv\Reader;

require __DIR__ . '/vendor/autoload.php';

$formatter = new \Covid19\Utils\ReportFormatter;
$getopts = new \Covid19\Utils\GetoptsHandler;

$getopts->handleOptions();
$options = $getopts->getOptions();

// TODO: ADD FACTORY
$csv = Reader::createFromPath(__DIR__ . $options['file'], 'r');
$csv->setHeaderOffset(0);

$service = new \Covid19\Service\ReportCasesService(
    $csv->getRecords(),
    new \Covid19\ValueObject\UserChoice($options['type'], $options['userChoice']),
    $options['lastDays']
);
$service->processData();

$formatter->output($service);
