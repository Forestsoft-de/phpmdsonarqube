package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"phpmdsonarqube/configuration"
	"phpmdsonarqube/reportreader"
	"phpmdsonarqube/sonar"
)

func main() {
	// _ ignores the error message

	config, _, err := parseArgs(os.Args[0], os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	issues := reportreader.ParseJson(config)
	writeJson(config, issues)
}

func writeJson(config *configuration.Config, issues *sonar.Sonar) {
	json, err := json.Marshal(issues)

	if err != nil {
		log.Fatal("Could not marshal json", err)
	}

	err = ioutil.WriteFile(config.Output, json, 0644)

	if err != nil {
		log.Fatal("Could not write file", config.Output, err)
	}
}
func parseArgs(progname string, args []string) (config *configuration.Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf configuration.Config
	flags.StringVar(&conf.Input, "input", "", "set input json file")
	flags.StringVar(&conf.Output, "output", "", "set output json file")
	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	return &conf, "", nil
}
