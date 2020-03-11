#!/usr/bin/env php
<?php

require __DIR__ . '/vendor/autoload.php';

$service = new \Covid19\Service\ReportCasesService(__DIR__ . '/data');

$service->handleOptions();
$data = $service->processData();
$service->printTotalCases($data);
