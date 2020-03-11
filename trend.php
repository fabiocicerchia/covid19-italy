#!/usr/bin/env php
<?php

require __DIR__ . '/vendor/autoload.php';

use League\Csv\Reader;

function readCsv($input, $key) {
    $data = [];

    $csvFiles = scandir(DATA_FOLDER);
    foreach($csvFiles as $csvFile) {
        if (preg_match('/^dpc-covid19-ita-(province|regioni)-\d{8}\.csv$/', $csvFile) !== 1) continue;
        $csv = Reader::createFromPath(DATA_FOLDER . $csvFile, 'r');
        $csv->setHeaderOffset(0);

        $data = array_merge($data, iterator_to_array($csv));
    }

    return $data;
}

function processData($input, $key) {
    $csv = readCsv($input, $key);

    $data = [];
    foreach ($csv as $record) {
        if (strtolower($record[$key]) === $input) {
            $data[$record['data']] = $record['totale_casi'];
        }
    }

    // TODO: MOVE THIS LOGIC IN THE CSV PROCESSING TO REDUCE THE DATA
    if (isset($options['last-week'])) $data = array_slice($data, 0, 7);
    elseif (isset($options['last-month'])) $data = array_slice($data, 0, 31);

    ksort($data);

    return $data;
}

function printTotalCases($input, $data) {
    $maxPadding = strlen(max($data));
    printf('TOTAL CASES IN %s' . PHP_EOL, strtoupper($input));
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
}

function handleOptions($input) {
    $shortopts = 'p:r:h';
    $longopts = ['province:', 'region:', 'help', 'last-week', 'last-month'];
    $options = getopt($shortopts, $longopts);

    $help = ($options['h'] ?? $options['help'] ?? null) !== null;
    if ($help) {
      echo 'COVID-19 ITALY - TRENDS' . PHP_EOL;
      echo ' ./trend.php [--province x|--region y] [--last-week|--last-month]' . PHP_EOL;
      exit;
    }

    $showProvince = strtolower($options['p'] ?? $options['province'] ?? null);
    $showRegion   = strtolower($options['r'] ?? $options['region'] ?? null);

    if ($showRegion) {
        define('DATA_FOLDER', __DIR__ . '/data/dati-regioni/');
        $key = 'denominazione_regione';
        $input = $showRegion;
    } elseif ($showProvince) {
        define('DATA_FOLDER', __DIR__ . '/data/dati-province/');
        $key = 'denominazione_provincia';
        $input = $showProvince;
    }

    return [$input, $key];
}

list($input, $key) = handleOptions(strtolower($_SERVER['argv'][1]));

$data = processData($input, $key);
printTotalCases($input, $data);
