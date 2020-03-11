<pre><?php

require __DIR__ . '/../vendor/autoload.php';

$type  = $_POST['type']; // region or province
$value = $_POST['value']; // name of place

$formatter = new \Covid19\Utils\ReportFormatter;
$getopts = new \Covid19\Utils\GetoptsHandler;

$getopts->handlePlace($type, $value);
$service = new \Covid19\Service\ReportCasesService(
    __DIR__ . '/../data',
    $getopts->getOptions()
);
$service->processData();
$formatter->output($service);
