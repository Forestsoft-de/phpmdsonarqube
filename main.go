package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// _ ignores the error message

	config, _, err := parseArgs(os.Args[0], os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	issues := parseJson(config)
	writeJson(config, issues)
}

type Config struct {
	input  string
	output string
}

type sonar struct {
	Issues []Issue
}

type Issue struct {
	EngineId        string          `json:"engineId"`
	RuleId          string          `json:"ruleId"`
	Typ             string          `json:"type"`
	Severity        string          `json:"severity"`
	PrimaryLocation PrimaryLocation `json:"primaryLocation"`
	effortMinutes   int
}

type PrimaryLocation struct {
	Message   string    `json:"message"`
	FilePath  string    `json:"filePath"`
	TextRange TextRange `json:"textRange"`
}

type TextRange struct {
	StartLine   int `json:"startLine"`
	EndLine     int `json:"endLine"`
	StartColumn int `json:"startColumn"`
	EndColumn   int `json:"endColumn"`
}
type phpmd struct {
	Version     string `json:"version"`
	PackageName string `json:"package"`
	Files       []file `json:"files"`
}
type file struct {
	Name       string      `json:"file"`
	Violations []violation `json:"violations"`
}
type violation struct {
	BeginLine       int    `json:"beginLine"`
	EndLine         int    `json:"endLine"`
	PackageName     string `json:"package"`
	Function        string `json:"function"`
	Class           string `json:"class"`
	Method          string `json:"method"`
	Description     string `json:"description"`
	Rule            string `json:"rule"`
	RuleSet         string `json:"ruleSet"`
	ExternalInfoUrl string `json:"externalInfoUrl"`
	Priority        int    `json:"priority"`
}

func writeJson(config *Config, issues *sonar) {
	json, err := json.Marshal(issues)

	if err != nil {
		log.Fatal("Could not marshal json", err)
	}

	err = ioutil.WriteFile(config.output, json, 0644)

	if err != nil {
		log.Fatal("Could not write file", config.output, err)
	}
}

func parseJson(config *Config) (issues *sonar) {
	content, err := os.ReadFile(config.input)

	if err != nil {
		log.Fatal("Could not read input file: '", config.input, "'", err)
	}

	data := phpmd{}
	unmarshallErr := json.Unmarshal([]byte(content), &data)

	if unmarshallErr != nil {
		log.Fatal("Could not unmarshall json", unmarshallErr)
	}

	issues = &sonar{}
	collection := make([]Issue, 0)

	for i := 0; i < len(data.Files); i++ {
		for violationCnt := 0; violationCnt < len(data.Files[i].Violations); violationCnt++ {
			issue := Issue{}
			issue.EngineId = "phpmd"
			issue.RuleId = data.Files[i].Violations[violationCnt].Rule
			issue.Typ = "VULNERABILITY"
			issue.Severity = getServerityByPriority(data.Files[i].Violations[violationCnt].Priority)
			issue.PrimaryLocation.Message = data.Files[i].Violations[violationCnt].Description
			issue.PrimaryLocation.FilePath = data.Files[i].Name
			issue.PrimaryLocation.TextRange.StartLine = data.Files[i].Violations[violationCnt].BeginLine
			issue.PrimaryLocation.TextRange.EndLine = data.Files[i].Violations[violationCnt].EndLine

			collection = append(collection, issue)
		}
	}
	issues.Issues = collection

	return issues
}

func getServerityByPriority(priority int) string {
	switch priority {
	case 1:
		return "BLOCKER"
	}
	return "INFO"
}

func parseArgs(progname string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf Config
	flags.StringVar(&conf.input, "input", "", "set input json file")
	flags.StringVar(&conf.output, "output", "", "set output json file")
	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	return &conf, "", nil
}
