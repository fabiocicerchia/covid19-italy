#!/usr/bin/env php
<?php

require __DIR__ . '/vendor/autoload.php';

use League\Csv\Reader;

define('DATA_FOLDER', __DIR__ . '/data/dati-province/');

$provincia = strtolower($_SERVER['argv'][1]);

$data = [];

$csvFiles = scandir(DATA_FOLDER);
foreach($csvFiles as $csvFile) {
    if (preg_match('/^dpc-covid19-ita-province-\d{8}\.csv$/', $csvFile) !== 1) continue;
    $csv = Reader::createFromPath(DATA_FOLDER . $csvFile, 'r');
    $csv->setHeaderOffset(0);

    foreach ($csv as $record) {
        if (strtolower($record['denominazione_provincia']) === $provincia) {
            $data[$record['data']] = $record['totale_casi'];
        }
    }
}

ksort($data);
$maxPadding = strlen(max($data));

printf('TOTAL CASES IN THE PROVINCE OF %s' . PHP_EOL, strtoupper($provincia));
$previous = 0;
foreach($data as $dataora => $casi) {
    $data = substr($dataora, 0, 10);
    printf('%s: %'.$maxPadding.'d', $data, $casi);
    if ($previous > 0) {
        $perc = (100 / $previous * $casi) - 100;
        printf(' -> %+4d%%', $perc);
    }
    printf(PHP_EOL);
    $previous = $casi;
}
