package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"phpmdsonarqube/configuration"
	"phpmdsonarqube/reportreader"
	"phpmdsonarqube/reportwriter"
)

func main() {
	// _ ignores the error message

	config, _, err := parseArgs(os.Args[0], os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	issues := reportreader.ParseJson(config)
	reportwriter.WriteJson(config, issues)
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
