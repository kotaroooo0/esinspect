#!/bin/sh

set -e

until curl -s localhost:9200 > /dev/null; do
  >&2 echo "Elasticsearch is unavailable - sleeping"
  sleep 3
done
>&2 echo "Elasticsearch is up - executing command"

sh $@

