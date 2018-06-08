# GOpom

[![Go Report Card](https://goreportcard.com/badge/github.com/edibImamovic/GOpom)](https://goreportcard.com/report/github.com/edibImamovic/GOpom)

edibImamovic/GOpom

GOpom are Page object model based on [llonchj/browsersteps](https://github.com/llonchj/browsersteps)
## Installation

    go get github.com/edibImamovic/GOpom

## Usage

1. Create cucumber features in `/features` folder.
1. Use this repository `GOpom/GOpom_test.go` file as `main_test.go` in your project.
1. Execute Selenium Server `selenium-standalone start` (https://www.seleniumhq.org/download/)
1. Run `godog` or `go test`.


## Acknowledgements

* [tebeka/selenium](https://github.com/tebeka/selenium) project for Selenium client for Golang.
* [DATA-DOG/godog](http://github.com/DATA-DOG/godog) project for Cucumber implementation for Golang.
* [llonchj/browsersteps](https://github.com/llonchj/browsersteps) Browser automation for Golang using Selenium and Cucumber


