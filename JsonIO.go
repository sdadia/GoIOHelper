package GoIOHelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func writeJsonToFile(data interface{}, fileName string) error {
	dataAsBytes, _ := json.Marshal(data)
	err := ioutil.WriteFile(fileName, dataAsBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
