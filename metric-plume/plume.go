package main

import (
	"os"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	_ "github.com/elastic/beats/metricbeat/module"
	_ "github.com/elastic/beats/metricbeat/include"
	mb "github.com/elastic/beats/metricbeat/beater"
)

type Plume struct {
	metricbeat beat.Beater
}

type PlumeBeat beat.Beat;

func New(b *beat.Beat, rawConfig *common.Config) (beat.Beater, error) {	
	mb, err := mb.New(b, rawConfig)

	if err != nil {
		return nil, err
	}

	plume := &Plume{metricbeat: mb}

	return plume, nil
}

func (plume *Plume) Run(b *beat.Beat) error {
    return plume.metricbeat.Run(b)
}

func (plume *Plume) Stop() {
    plume.metricbeat.Stop()
}

func main() {
    if err := beat.Run("plume", "", New); err != nil {
		os.Exit(1)
	}
}
