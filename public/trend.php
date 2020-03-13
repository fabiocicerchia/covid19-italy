<pre><?php

require __DIR__ . '/../vendor/autoload.php';

$type  = $_POST['type']; // region or province
$value = $_POST['value']; // name of place

$formatter = new \Covid19\Utils\ReportFormatter;
$getopts   = new \Covid19\Utils\GetoptsHandler;

$getopts->handlePlace($type, $value);

$service = \Covid19\Factory\ReportFactory::create($getopts->getOptions());
$service->processData();

$formatter->output($service);
