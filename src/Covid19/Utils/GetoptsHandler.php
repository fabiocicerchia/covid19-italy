<?php

namespace Covid19\Utils;

class GetoptsHandler
{
    const DATA_FILE_PROVINCE = '/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv';
    const DATA_FILE_REGIONI  = '/data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv';

    protected $options = [];

    public function handleOptions()
    {
        $options = getopt(
            '',
            ['province:', 'region:', 'help', 'last-week', 'last-month']
        );
    
        if (isset($options['help'])) {
            die($this->usageHelp());
        }
    
        $province = $options['province'] ?? null;
        $region   = $options['region'] ?? null;

        $type  = $province ? 'province' : ($region ? 'region' : null);
        $value = $province ?? $region ?? null;

        $this->options['lastDays'] = isset($options['last-week']) ? 7 : null;
        $this->options['lastDays'] = isset($options['last-month']) ? 31 : $this->options['lastDays'];
    
        $this->handlePlace($type, $value);
    }

    public function handlePlace($type, $value)
    {
        if ($type === 'region') {
            $this->showRegion(strtolower($value));
        } elseif ($type === 'province') {
            $this->showProvince(strtolower($value));
        } else {
            die($this->usageHelp('You must select a region or a province'));
        }
    }

    protected function showRegion(string $name)
    {
        $this->options['file']       = self::DATA_FILE_REGIONI;
        $this->options['type']       = 'regione';
        $this->options['userChoice'] = $name;
    }

    protected function showProvince(string $name)
    {
        $this->options['file']       = self::DATA_FILE_PROVINCE;
        $this->options['type']       = 'provincia';
        $this->options['userChoice'] = $name;
    }

    protected function usageHelp(string $customMessage = null)
    {
        return 'COVID-19 ITALY - TRENDS' . PHP_EOL .
               ' ./trend.php [--province x|--region y] [--last-week|--last-month]' . PHP_EOL .
               ($customMessage ? ('NOTE: ' . $customMessage . PHP_EOL) : '');
    }

    public function getOptions()
    {
        return $this->options;
    }
}
