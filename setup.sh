#!/bin/bash
git submodule init
git submodule update
composer install
echo "Show data http://127.0.0.1:8001/index.html"
php -S 127.0.0.1:8081 -t public
./trend.php
