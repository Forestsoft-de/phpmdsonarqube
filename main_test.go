package main

/*
  https://docs.sonarqube.org/latest/analysis/generic-issue/
*/
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"phpmdsonarqube/configuration"
	"phpmdsonarqube/reportreader"
	"phpmdsonarqube/reportwriter"
	"phpmdsonarqube/sonar"
	"reflect"
	"testing"
)

func TestArgs(t *testing.T) {
	t.Run("Args should parsed", func(t *testing.T) {
		expectedConf := configuration.Config{Input: "datasets/sonar.json", Output: "output.json"}

		actual, buffer, err := parseArgs("phpmdsonaqube", []string{"-input", "datasets/sonar.json", "-output", "output.json"})

		if actual == nil {
			t.Errorf("Actual is nil, buffer: '%s', error: '%s'", buffer, err)
		}

		if actual.Input != expectedConf.Input {
			t.Errorf("Test failed for input file, expected: '%s', got:  '%s'", expectedConf.Input, actual.Input)
		}

		if actual.Output != expectedConf.Output {
			t.Errorf("Test failed for output file, expected: '%s', got:  '%s'", expectedConf.Output, actual.Input)
		}
	})

}

func TestParseJson(t *testing.T) {
	t.Run("parseJson should return files with violations", func(t *testing.T) {
		config := &configuration.Config{Input: "datasets/phpmd.json", Output: "output.json"}
		issues := reportreader.ParseJson(config)

		if len(issues.Issues) != 4 {
			t.Errorf("parseJson should return 4 issues, got:  '%d'", len(issues.Issues))
		}
	})
}

func TestWriteJson(t *testing.T) {
	t.Run("writeJson should create correct json file", func(t *testing.T) {
		config := &configuration.Config{Input: "datasets/phpmd.json", Output: "/tmp/output.json"}

		sonarConfig := &sonar.Sonar{}

		issues := make([]sonar.Issue, 0)

		textRange := sonar.TextRange{}
		textRange.StartLine = 30
		textRange.EndLine = 30
		textRange.StartColumn = 9
		textRange.EndColumn = 14

		location := sonar.PrimaryLocation{}
		location.FilePath = "sources/A.java"
		location.Message = "fully-fleshed issue"
		location.TextRange = textRange

		issue1 := sonar.Issue{
			EngineId: "phpmd",
			RuleId:   "S1234",
			Typ:      "CODE_SMELL",
			Severity: "BLOCKER",
		}

		//issue1 := Issue{}
		issue1.PrimaryLocation = location

		issues = append(issues, issue1)
		sonarConfig.Issues = issues

		reportwriter.WriteJson(config, sonarConfig)

		expected, err1 := ioutil.ReadFile("datasets/sonar.json")

		if err1 != nil {
			log.Fatal(err1)
		}

		actual, err2 := ioutil.ReadFile("/tmp/output.json")

		if err2 != nil {
			log.Fatal(err2)
		}

		var err error
		expectedJson := sonar.Sonar{}
		actualJson := sonar.Sonar{}

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
