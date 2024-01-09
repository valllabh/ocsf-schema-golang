#!/bin/bash


while IFS= read -r -d '' enumValueMapJson; do
    echo "Processing $enumValueMapJson"

    # replace proto from $file
    newEnumValueMapJson=${enumValueMapJson/proto\//}
    newEnumValueMapJson=${newEnumValueMapJson/enum-value-map.json/enum-value-map/enum-value-map.json}

    # create directory if not exist
    mkdir -p $(dirname $newEnumValueMapJson)

    # copy file
    echo "Creating to $newEnumValueMapJson"
    cp $enumValueMapJson $newEnumValueMapJson

    # replace enum-value-map.json to enum-value-map.go.template from $newFile
    enumValueMapGoFile=${newEnumValueMapJson/enum-value-map.json/map.go}
    
     # create directory if not exist
    mkdir -p $(dirname $enumValueMapGoFile)

    # copy file
    echo "Creating to $enumValueMapGoFile"
    cp "./enum-value/enum-value-map.go.template" $enumValueMapGoFile

    echo ""

done < <(find ./proto -name "enum-value-map.json" -print0)

