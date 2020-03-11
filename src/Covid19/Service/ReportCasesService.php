<?php

namespace Covid19\Service;

use League\Csv\Reader;

class ReportCasesService
{
    const REGEX_DATA_FILES = '/^dpc-covid19-ita-(province|regioni)-\d{8}\.csv$/';

    protected $userChoice;
    protected $type;
    protected $dataFolder;
    protected $folder;
    protected $lastWeek = false;
    protected $lastMonth = false;

    public function __construct(string $dataFolder) {
        $this->dataFolder = $dataFolder;
    }

    public function handleOptions() {
        $options = getopt(
            '',
            ['province:', 'region:', 'help', 'last-week', 'last-month']
        );
    
        if (isset($options['help'])) {
            die($this->usageHelp());
        }
    
        $showProvince = strtolower($options['province'] ?? null);
        $showRegion   = strtolower($options['region'] ?? null);

        $this->lastWeek = isset($options['last-week']);
        $this->lastMonth = isset($options['last-month']);
    
        if ($showRegion !== '') {
            $this->folder = $this->dataFolder . '/dati-regioni/';
            $this->type = 'denominazione_regione';
            $this->userChoice = $showRegion;
        } elseif ($showProvince !== '') {
            $this->folder = $this->dataFolder . '/dati-province/';
            $this->type = 'denominazione_provincia';
            $this->userChoice = $showProvince;
        } else {
            die($this->usageHelp('You must select a region or a province'));
        }
    }

    protected function usageHelp(string $customMessage = null) {
        return 'COVID-19 ITALY - TRENDS' . PHP_EOL .
               ' ./trend.php [--province x|--region y] [--last-week|--last-month]' . PHP_EOL .
               ($customMessage ? ('NOTE: ' . $customMessage . PHP_EOL) : '');
    }

    protected function readCsv() {
        $files = scandir($this->folder);
        $csvFiles = array_filter($files, function ($item) { 
            return (preg_match(self::REGEX_DATA_FILES, $item) === 1);
        });

        $data = [];
        foreach($csvFiles as $csvFile) {
            $csv = Reader::createFromPath($this->folder . $csvFile, 'r');
            $csv->setHeaderOffset(0);
    
            $data = array_merge($data, iterator_to_array($csv));
        }
    
        return $data;
    }
    
    public function processData() {
        $csv = $this->readCsv();
    
        $data = [];
        foreach ($csv as $record) {
            if (strtolower($record[$this->type]) === $this->userChoice) {
                $data[$record['data']] = $record['totale_casi'];
            }
        }
    
        // TODO: MOVE THIS LOGIC IN THE CSV PROCESSING TO REDUCE THE DATA
        if ($this->lastMonth) $data = array_slice($data, max(count($data) - 31, 0), 31);
        elseif ($this->lastWeek) $data = array_slice($data, max(count($data) - 7, 0), 7);
    
        ksort($data);
    
        return $data;
    }
    
    public function printTotalCases($data) {
        printf('COVID-19 CASES IN %s' . PHP_EOL, strtoupper($this->userChoice));

        if (empty($data)) {
            printf('No cases found' . PHP_EOL);
            return;
        }

        $maxPadding = strlen(max($data));

        $sum = $previous = 0;
        foreach($data as $dataora => $casi) {
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
