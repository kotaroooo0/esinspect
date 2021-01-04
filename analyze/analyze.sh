#!/bin/sh

set -e

curl  -XGET "http://localhost:9200/$1/_analyze" -H 'Content-Type: application/json' -d"{  \"text\": [\"$2\"],  \"analyzer\": \"my_analyzer\"}"
