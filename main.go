package main

import (
	"fmt"
	"flag"
	"upm/logger"
	"upm/config"
	"upm/pkg"
)

/*
	As init() function is responsible for CLI arguments parsing,
	we have to store jobs to run in main after applying config
*/
const UNPACK = 0

var cfg config.UPMConfig

type Job struct {
	jobType int
	data interface{}
}

type unpackJob struct {
	from string
}

var jobs []Job

func init(){
	/*
		Init function

		Parses CLI arguments and init config file
	*/

	var logLvl int
	var configPath string
	var unpackPath string

	flag.IntVar(&logLvl, "log", 1, "Verbose level: 0-NONE, 1-INFO, 2-DEBUG")
	flag.StringVar(&configPath, "config", "./config/default.conf", "Provide custom path to config file")
	flag.StringVar(&unpackPath, "unpack", "", "Unpack selected package file")

	flag.Parse()

	log := logger.UPMLogger{
		Lvl: logLvl,
	}
	// Logger injection
	config.Log = log
	pkg.Log = log

	var err error
	cfg, err = config.ReadConfig(configPath)

	if err != nil {
		log.Error("Failed to load config: %s", err)
	}
	jobs = append(jobs, Job{
		UNPACK,
		unpackJob{unpackPath},
	})
}

func main() {
	fmt.Println(cfg)
	out := make(chan interface{})

	for _, job := range jobs {
		go func(job Job) {
			switch(job.jobType) {
				case UNPACK:
					out <- pkg.Unpack(job.data.(unpackJob).from)
				default: break;
			}
		}(job)
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Println(<-out)
	}
}

