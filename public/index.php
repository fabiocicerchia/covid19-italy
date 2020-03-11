<pre><?php

use League\Csv\Reader;

require __DIR__ . '/../vendor/autoload.php';

$type  = $_POST['type']; // region or province
$value = $_POST['value']; // name of place

$formatter = new \Covid19\Utils\ReportFormatter;
$getopts = new \Covid19\Utils\GetoptsHandler;

$getopts->handlePlace($type, $value);
$options = $getopts->getOptions();

// TODO: ADD FACTORY
$csv = Reader::createFromPath(__DIR__ . '/..' . $options['file'], 'r');
$csv->setHeaderOffset(0);

$service = new \Covid19\Service\ReportCasesService(
    $csv->getRecords(),
    new \Covid19\ValueObject\UserChoice($options['type'], $options['userChoice']),
    $options['lastDays']
);
$service->processData();

$formatter->output($service);
