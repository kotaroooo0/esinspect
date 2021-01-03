#!/bin/sh

set -e

curl -XPUT "http://localhost:9200/$1" -H 'Content-Type: application/json' -d@$2
