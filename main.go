package main

import (
	"fmt"
	"flag"
	"os"

	"upm/logger"
	"upm/config"
	"upm/pkg/manager"
)

/*
	As init() function is responsible for CLI arguments parsing,
	we have to store jobs to run in main after applying config
*/
const UNPACK = 0

type Job struct {
	jobType int
	data interface{}
}

type unpackJob struct {
	from string
}

var jobs []Job

type unpackFlags []string

func (i *unpackFlags) String() string {
	return ""
}

func (i *unpackFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func init(){
	/*
		Init function

		Parses CLI arguments and init config file
	*/

	var logLvl int
	var configPath string
	var unpackList unpackFlags

	flag.IntVar(&logLvl, "log", 1, "Verbose level: 0-NONE, 1-INFO, 2-DEBUG")
	flag.StringVar(&configPath, "config", "./config/default.conf", "Provide custom path to config file")
	flag.Var(&unpackList, "unpack", "Unpack selected packages")

	flag.Parse()

	config.Config.ReadConfig(configPath)
	config.Config.Log.Level = logLvl
	logger.Log.Init(config.Config.Log.Level)

	if len(unpackList) != 0 {
		for _, path := range unpackList {
			jobs = append(jobs, Job{
				UNPACK,
				unpackJob{path},
			})
		}
	}
}

func main() {
	// Creating cache directory if doesn't exist
	Log := logger.Log
	Log.SetPrefix("main: ")

	err := os.MkdirAll(config.Config.Cache.Dir, 0755)
	if err != nil {
		Log.Fatal("%s", err)
	}

	out := make(chan interface{})

	for _, job := range jobs {
		go func(job Job) {
			switch(job.jobType) {
				case UNPACK:
					out <- manager.Unpack(job.data.(unpackJob).from)
				default: break;
			}
		}(job)
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Println(<-out)
	}
}

