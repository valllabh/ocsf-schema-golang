#!/bin/bash

set -euxo pipefail

verions=("v1.0" "v1.1" "v1.2" "v1.3")
for i in "${verions[@]}"
do
    config_file="/configs/config-${i}.yaml"
    echo "${i}: Getting event list..."
    rm -f ./schema/tmp.json
    ocsf-tool schema-class-list --config "${config_file}" --output "./schema/tmp.json"

    jq . "./schema/tmp.json" > "./schema/${i}-schema-class-list.json"
    rm -f ./schema/tmp.json

    echo "${i}: Creating proto files..."
    jq -r "[.[] | to_entries[] | .value] | join(\" \")" "./schema/${i}-schema-class-list.json" | xargs ocsf-tool generate-proto --config "${config_file}"  --golang-root-package / --proto-output ./proto

    jq . schema.json > "schema/${i}-schema.json"
    rm -f schema.json
    echo "${i}: done"
done

