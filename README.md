# Golang FHIR Models

This repository contains FHIR® models for Go. The models consist of Go structs for each resource and data type in the `fhir-models` module. The structs are suitable for JSON encoding/decoding. The repository also contains a generator for creating Go models from FHIR Definitions.

## Features

* resources implement the [Marshaler][1] interface
* Unmarshal functions are provided for every resource
* enums are provided for every ValueSet used in a [required binding][2], has a computer friendly name and refers only to one CodeSystem
* enums implement `Code()`, `Display()` and `Definition()` methods

## Development

The generator in the `fhir-models-gen` module is self-hosting, i.e. it depends data types that it has generated itself. To update to newer FHIR versions, make sure to first install the current version of the generator. Then regenerate the reduced set of definitions in `fhir-models-gen/fhir` that the generator itself needs. Lastly, rebuild the generator. It can then be used to generate current models for the core FHIR definitions in `fhir-models` or other FHIR profiles.

[1]: <https://golang.org/pkg/encoding/json/#Marshaler>
[2]: <https://www.hl7.org/fhir/terminologies.html#strength>
