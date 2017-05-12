# pyroclast-plumes

Plumes send data to Pyroclast.

## Plumes

- file-plume
- metric-plume

## Build from source usage

1. Clone this repository
2. `cd` to the directory of the Plume you want to use
3. Run `go get github.com/elastic/beats`
4. Run `go build plume.go`
5. Link the Plume binary, perhaps with `ln -s /usr/bin/file-plume file-plume`
6. Download a Plume Configuration file from Pyroclast
7. Start sending your data: `file-plume -e -c downloaded-configuration.yml -c file-plume.yml`

## File Plume Docker usage

1. Clone this repository
2. `cd` to the file-plume directory
3. Build the container with `docker build -t pyroclast-file-plume:0.1.0 .`
4. Run the container with: `docker run -e TOPIC_FILE=<topic-file> -v $PWD:/opt/pyroclast-plumes/topics -v <data-dir>:/data pyroclast-file-plume:0.1.0`

In (4), `topic-file` is the configuration file you download from the Topic page. `<data-dir>` is the directory on your host file system with files to watch and transmit.

## Credit

Plumes are thin wrappers around the excellent ElasticSearch [Beats](https://www.elastic.co/products/beats) library.
