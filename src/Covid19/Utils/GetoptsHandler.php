<?php

namespace Covid19\Utils;

class GetoptsHandler
{
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

        $this->options['lastWeek']  = isset($options['last-week']);
        $this->options['lastMonth'] = isset($options['last-month']);
    
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
        $this->options['folder']     = '/dati-regioni/';
        $this->options['type']       = 'denominazione_regione';
        $this->options['userChoice'] = $name;
    }

    protected function showProvince(string $name)
    {
        $this->options['folder']     = '/dati-province/';
        $this->options['type']       = 'denominazione_provincia';
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
