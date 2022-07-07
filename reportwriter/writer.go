package reportwriter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"phpmdsonarqube/configuration"
	"phpmdsonarqube/sonar"
)

func WriteJson(config *configuration.Config, issues *sonar.Sonar) {
	json, err := json.Marshal(issues)

	if err != nil {
		log.Fatal("Could not marshal json", err)
	}

	err = ioutil.WriteFile(config.Output, json, 0644)

	if err != nil {
		log.Fatal("Could not write file", config.Output, err)
	}
}
