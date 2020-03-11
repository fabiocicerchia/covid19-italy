<?php

namespace Covid19\Utils;

use Covid19\Service\ReportCasesService;

class ReportFormatter
{
    public function output(ReportCasesService $service)
    {
        $data = $service->getData();

        printf('COVID-19 CASES IN %s' . PHP_EOL, strtoupper($service->getUserChoice()->value()));

        if (empty($data)) {
            printf('No cases found' . PHP_EOL);
            return;
        }

        $maxPadding = strlen(max($data));

        $sum = $previous = 0;
        foreach ($data as $dataora => $casi) {
            $sum += $casi;
            $data = (\DateTime::createFromFormat('Y-m-d H:i:s', $dataora))->format('Y-m-d');
            printf('%s: %' . $maxPadding . 'd', $data, $casi);
            if ($previous > 0) {
                $perc = (100 / $previous * $casi) - 100;
                printf(' -> %+4d%%', $perc);
            }
            printf(PHP_EOL);
            $previous = $casi;
        }

        printf('TOTAL CASES: %d' . PHP_EOL, $sum);
    }
}
