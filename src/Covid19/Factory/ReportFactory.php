<?php

namespace Covid19\Factory;

use League\Csv\Reader;

class ReportFactory
{
    public static function create(array $options)
    {
        $csv = Reader::createFromPath(__DIR__ . '/../../..' . $options['file'], 'r');
        $csv->setHeaderOffset(0);

        $service = new \Covid19\Service\ReportCasesService(
            $csv->getRecords(),
            new \Covid19\ValueObject\UserChoice($options['type'], $options['userChoice']),
            $options['lastDays']
        );

        return $service;
    }
}
