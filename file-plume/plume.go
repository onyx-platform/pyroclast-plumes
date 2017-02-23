package main

import (
	"os"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	fb "github.com/elastic/beats/filebeat/beater"
)

type Plume struct {
	filebeat beat.Beater
}

type PlumeBeat beat.Beat;

func New(b *beat.Beat, rawConfig *common.Config) (beat.Beater, error) {	
	fb, err := fb.New(b, rawConfig)

	if err != nil {
		return nil, err
	}

	plume := &Plume{filebeat: fb}

	return plume, nil
}

func (plume *Plume) Run(b *beat.Beat) error {
    return plume.filebeat.Run(b)
}

func (plume *Plume) Stop() {
    plume.filebeat.Stop()
}

func main() {
    if err := beat.Run("plume", "", New); err != nil {
		os.Exit(1)
	}
}
