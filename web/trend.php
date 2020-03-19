<pre><?php

$type  = preg_replace('/[^a-z]/i', '', $_POST['type']); // region or province
$value = preg_replace('/[^a-z\']/i', '', $_POST['value']); // region or province

$command = 'cd '.__DIR__.'/../ && ./bin/covid19-trend ';
if ($type === 'region') {
    $command .= '--region ' . $value;
} elseif ($type === 'province') {
    $command .= '--province ' . $value;
}
$command .= ' 2>&1';

echo system($command);
