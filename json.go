package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readIPDataJSON() (ipd IPData) {
	raw, err := ioutil.ReadFile(fileIPDataJSON)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(raw, &ipd)
	return
}

func writeIPDataJSON(ipd IPData) {
	JSONstring, _ := json.Marshal(ipd)
	err = ioutil.WriteFile(fileIPDataJSON, JSONstring, 0644)
}
