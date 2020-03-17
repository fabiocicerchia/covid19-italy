<pre><?php

$type  = preg_replace('/[^a-z]/i', '', $_POST['type']); // region or province
$value = preg_replace('/[^a-z\']/i', '', $_POST['value']); // region or province

$command = 'go run ../main.go '
if ($type === 'region') {
    $command .= '--region ' . $value;
} elseif ($type === 'province') {
    $command .= '--province ' . $value;
}

exec($command);
