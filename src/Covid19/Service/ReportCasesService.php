<?php

namespace Covid19\Service;

use League\Csv\Reader;

class ReportCasesService
{
    const REGEX_DATA_FILES = '/^dpc-covid19-ita-(province|regioni)-\d{8}\.csv$/';

    protected $userChoice;
    protected $type;
    protected $folder;
    protected $lastWeek = false;
    protected $lastMonth = false;

    public function __construct(string $folder, array $options = [])
    {
        $this->folder = $folder . $options['folder'];

        if (isset($options['type'])) {
            $this->type = (string) $options['type'];
        }
        if (isset($options['userChoice'])) {
            $this->userChoice = (string) $options['userChoice'];
        }
        if (isset($options['lastWeek'])) {
            $this->lastWeek = (bool) $options['lastWeek'];
        }
        if (isset($options['lastMonth'])) {
            $this->lastMonth = (bool) $options['lastMonth'];
        }
    }

    protected function readCsv()
    {
        $files = scandir($this->folder);
        $csvFiles = array_filter($files, function ($item) {
            return (preg_match(self::REGEX_DATA_FILES, $item) === 1);
        });

        $data = [];
        foreach ($csvFiles as $csvFile) {
            $csv = Reader::createFromPath($this->folder . $csvFile, 'r');
            $csv->setHeaderOffset(0);
    
            $data = array_merge($data, iterator_to_array($csv));
        }
    
        return $data;
    }
    
    public function processData()
    {
        $csv = $this->readCsv();
    
        $data = [];
        foreach ($csv as $record) {
            if (strtolower($record[$this->type]) === $this->userChoice) {
                $data[$record['data']] = $record['totale_casi'];
            }
        }
    
        // TODO: MOVE THIS LOGIC IN THE CSV PROCESSING TO REDUCE THE DATA
        if ($this->lastMonth) {
            $data = array_slice($data, max(count($data) - 31, 0), 31);
        } elseif ($this->lastWeek) {
            $data = array_slice($data, max(count($data) - 7, 0), 7);
        }
    
        ksort($data);
    
        $this->data = $data;
    }

    public function getData()
    {
        return $this->data;
    }

    public function getUserChoice()
    {
        return $this->userChoice;
    }
}
