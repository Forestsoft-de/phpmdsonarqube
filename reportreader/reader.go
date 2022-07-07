package reportreader

import (
	"encoding/json"
	"log"
	"os"
	"phpmdsonarqube/configuration"
	"phpmdsonarqube/sonar"
)

func ParseJson(config *configuration.Config) (issues *sonar.Sonar) {
	content, err := os.ReadFile(config.Input)

	if err != nil {
		log.Fatal("Could not read input file: '", config.Input, "'", err)
	}

	data := sonar.Phpmd{}
	unmarshallErr := json.Unmarshal([]byte(content), &data)

	if unmarshallErr != nil {
		log.Fatal("Could not unmarshall json", unmarshallErr)
	}

	issues = &sonar.Sonar{}
	collection := make([]sonar.Issue, 0)

	for i := 0; i < len(data.Files); i++ {
		for violationCnt := 0; violationCnt < len(data.Files[i].Violations); violationCnt++ {
			issue := sonar.Issue{}
			issue.EngineId = "phpmd"
			issue.RuleId = data.Files[i].Violations[violationCnt].Rule
			issue.Typ = "CODE_SMELL"
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
	case 2:
		return "CRITICAL"
	case 3:
		return "MAJOR"
	case 4:
		return "MINOR"
	}
	return "INFO"
}
