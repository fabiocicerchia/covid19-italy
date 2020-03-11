<?php

namespace Covid19\Service;

use Covid19\ValueObject\UserChoice;

class ReportCasesService
{
    protected $userChoice;
    protected $file;
    protected $lastDays;

    public function __construct(\Iterator $csvDocument, UserChoice $userChoice, int $lastDays = null)
    {
        $this->file       = $csvDocument;
        $this->userChoice = $userChoice;
        $this->lastDays   = $lastDays;
    }

    public function processData()
    {
        // the file is sorted by date ascending, we need it descending
        // TODO: move it somewhere else
        $csvData = array_reverse(iterator_to_array($this->file));
    
        $typeKey = $this->getTypeKey();
        $data = [];

        $countDays = 0;
        foreach ($csvData as $record) {
            if (strtolower($record[$typeKey]) === $this->userChoice->value()) {
                $data[$record['data']] = $record['totale_casi'];

                $countDays++;
                if ($countDays === $this->lastDays) break;
            }
        }

        ksort($data);
    
        $this->data = $data;
    }

    private function getTypeKey()
    {
        return $this->userChoice->type() === 'regione'
               ? 'denominazione_regione'
               : (
                   $this->userChoice->type() === 'provincia'
                  ? 'denominazione_provincia'
                  : null
                 );
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
