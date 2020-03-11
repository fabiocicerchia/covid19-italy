<pre><?php

require __DIR__ . '/../vendor/autoload.php';

$service = new \Covid19\Service\ReportCasesService(__DIR__ . '/../data');

$type  = $_POST['type']; // region or province
$value = $_POST['value']; // name of place

$service->handlePlace($type, $value);
$data = $service->processData();
$service->printTotalCases($data);
