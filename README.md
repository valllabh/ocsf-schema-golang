# Go Implementation of Open Cybersecurity Schema Framework (OCSF)
## Overview

This repository, `ocsf-schema-golang`, offers a Go implementation of the [Open Cybersecurity Schema Framework (OCSF)](https://schema.ocsf.io/). The OCSF is an initiative aimed at standardizing the schema for cybersecurity data. This implementation in Go is tailored for developers who need an effective way to handle OCSF-compliant data within their Go-based applications.

## Features 
- **OCSF Version** : OCSF V1.0.0 
- **Compliance with OCSF Schema** : Fully aligned with the OCSF standards. 
- **Performance-Optimized** : Engineered for high-speed and efficient processing of cybersecurity data. 
- **Seamless Integration** : Crafted for easy incorporation into Go applications.
## Getting Started
### Prerequisites
- Go version 1.21 or later
### Installation

Install the package using the Go command:

```bash
go get github.com/valllabh/ocsf-schema-golang
```


### Usage

Basic usage example:

```go
package main

import "github.com/valllabh/ocsf-schema-golang/ocsf/v1_0_0/events/findings"

func main() {

	securityFinding := findings.SecurityFinding{
		Message: "This is a sample message",
	}

}
```


## Internal Dependencies

This project leverages `ocsf-tool`, available at [OCSF Tool](https://github.com/valllabh/ocsf-tool). The OCSF-Tool is a command-line utility designed for managing OCSF schemas and supports the generation of Proto files for OCSF Schemas. These Proto files are then utilized for generating Go files within this project.

## License

This project is licensed under the [Apache License 2.0](LICENSE) .

