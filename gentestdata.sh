#!/bin/bash
# Generate test data from GitHub status API data

set -e
set -o pipefail

download_json() {
    curl -Ls "$1" | python -mjson.tool > "$2"
}

rm -rf testdata
mkdir -p testdata/api

download_json https://status.github.com/api.json testdata/api.json

grep json testdata/api.json | cut -d\" -f4 | while read url; do
    echo "$url"
    download_json "$url" testdata/api/`basename $url`
done
