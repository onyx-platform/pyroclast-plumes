# pyroclast-plumes

Plumes send data to Pyroclast.

## Plumes

- file-plume
- metric-plume

## Usage

1. Clone this repository
2. `cd` to the directory of the Plume you want to use
3. Run `go get github.com/elastic/beats`
4. Run `go build plume.go`
5. Link the Plume binary, perhaps with `ln -s /usr/bin/file-plume file-plume`
6. Download a Plume Configuration file from Pyroclast
7. Start sending your data: `file-plume -e -c downloaded-configuration.yml -c file-plume.yml`

## Credit

Plumes are thin wrappers around the excellent ElasticSearch [Beats](https://www.elastic.co/products/beats) library.
