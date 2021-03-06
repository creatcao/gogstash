package cmd

import (
	"runtime"
	"time"

	"github.com/tsaikd/gogstash/config"

	// module loader
	_ "github.com/tsaikd/gogstash/modloader"
)

func gogstash(confpath string) (err error) {
	logger := config.Logger

	if runtime.GOMAXPROCS(0) == 1 && runtime.NumCPU() > 1 {
		logger.Warnf("set GOMAXPROCS = %d to get better performance", runtime.NumCPU())
	}

	conf, err := config.LoadFromFile(confpath)
	if err != nil {
		return
	}

	if err = conf.RunInputs(); err != nil {
		return
	}

	if err = conf.RunFilters(); err != nil {
		return
	}

	if err = conf.RunOutputs(); err != nil {
		return
	}

	for {
		// all event run in routine, go into infinite sleep
		time.Sleep(1 * time.Hour)
	}
}
