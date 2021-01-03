#!/bin/sh

set -e

curl -XGET http://localhost:9200/$1/_search -H 'Content-Type: application/json' -d @$2
