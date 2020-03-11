#!/usr/bin/env php
<?php

require __DIR__ . '/vendor/autoload.php';

$formatter = new \Covid19\Utils\ReportFormatter;
$getopts = new \Covid19\Utils\GetoptsHandler;

$getopts->handleOptions();
$service = new \Covid19\Service\ReportCasesService(
    __DIR__ . '/data',
    $getopts->getOptions(),
);
$service->processData();
$formatter->output($service);
