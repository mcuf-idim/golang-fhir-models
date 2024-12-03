#!/usr/bin/env bash

ln -s ../../../fhir-models-gen/fhir/root.go fhir/root.go
wget -O definitions.zip https://www.hl7.org/fhir/r4b/definitions.json.zip
unzip definitions.zip profiles-resources.json profiles-types.json valuesets.json -d fhir
rm definitions.zip

go generate ./fhir
