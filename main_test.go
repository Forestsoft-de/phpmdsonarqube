package main

/*
  https://docs.sonarqube.org/latest/analysis/generic-issue/
*/
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestArgs(t *testing.T) {
	t.Run("Args should parsed", func(t *testing.T) {
		expectedConf := Config{"datasets/sonar.json", "output.json"}

		actual, buffer, err := parseArgs("phpmdsonaqube", []string{"-input", "datasets/sonar.json", "-output", "output.json"})

		if actual == nil {
			t.Errorf("Actual is nil, buffer: '%s', error: '%s'", buffer, err)
		}

		if actual.input != expectedConf.input {
			t.Errorf("Test failed for input file, expected: '%s', got:  '%s'", expectedConf.input, actual.input)
		}

		if actual.output != expectedConf.output {
			t.Errorf("Test failed for output file, expected: '%s', got:  '%s'", expectedConf.output, actual.input)
		}
	})

}

func TestParseJson(t *testing.T) {
	t.Run("parseJson should return files with violations", func(t *testing.T) {
		config := &Config{input: "datasets/phpmd.json", output: "output.json"}
		issues := parseJson(config)

		if len(issues.Issues) != 4 {
			t.Errorf("parseJson should return 4 issues, got:  '%d'", len(issues.Issues))
		}
	})
}

func TestWriteJson(t *testing.T) {
	t.Run("writeJson should create correct json file", func(t *testing.T) {
		config := &Config{input: "datasets/phpmd.json", output: "/tmp/output.json"}

		sonarConfig := &sonar{}

		issues := make([]Issue, 0)

		textRange := TextRange{}
		textRange.StartLine = 30
		textRange.EndLine = 30
		textRange.StartColumn = 9
		textRange.EndColumn = 14

		location := PrimaryLocation{}
		location.FilePath = "sources/A.java"
		location.Message = "fully-fleshed issue"
		location.TextRange = textRange

		issue1 := Issue{
			EngineId: "phpmd",
			RuleId:   "S1234",
			Typ:      "CODE_SMELL",
			Severity: "BLOCKER",
		}

		//issue1 := Issue{}
		issue1.PrimaryLocation = location

		issues = append(issues, issue1)
		sonarConfig.Issues = issues

		writeJson(config, sonarConfig)

		expected, err1 := ioutil.ReadFile("datasets/sonar.json")

		if err1 != nil {
			log.Fatal(err1)
		}

		actual, err2 := ioutil.ReadFile("/tmp/output.json")

		if err2 != nil {
			log.Fatal(err2)
		}

		var err error
		expectedJson := sonar{}
		actualJson := sonar{}

		err = json.Unmarshal([]byte(expected), &expectedJson)
		if err != nil {
			t.Errorf("Error mashalling string 1 :: %s", err.Error())
		}

		err = json.Unmarshal([]byte(actual), &actualJson)
		if err != nil {
			t.Errorf("Error mashalling string 2 :: %s", err.Error())
		}

		equal := reflect.DeepEqual(expectedJson, actualJson)

		if !equal {
			t.Errorf("Files are not equal: \n+++\n '%v'\n---\n'%v'", actualJson, expectedJson)
		}
	})
}
