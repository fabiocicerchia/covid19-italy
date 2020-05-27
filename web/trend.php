<?php

$type   = preg_replace('/[^a-z]/i', '', $_GET['type']); // region or province
$value  = preg_replace('/[^a-z .\']/i', '', $_GET['value']); // region or province
$format = $_GET['format'] === 'json' ? 'json' : 'text';

$command = 'cd '.__DIR__.'/../ && ./bin/covid19-trend --format ' . $format;
if ($type === 'region') {
    $command .= ' --region "' . $value . '"';
} elseif ($type === 'province') {
    $command .= ' --province "' . $value . '"';
}
$command .= ' 2>&1';

system($command);
