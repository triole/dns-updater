package main

import (
	"encoding/json"
	"io/ioutil"
)

func readIPDataJSON() (ipd tIPDataSet) {
	raw, err := ioutil.ReadFile(fileIPDataJSON)
	lg.LogError("Error reading ip data json", err)
	err = json.Unmarshal(raw, &ipd)
	lg.LogError("Can not unmarshal ip data json", err)
	return
}

func writeIPDataJSON(ipd tIPDataSet) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = ioutil.WriteFile(fileIPDataJSON, JSONstring, 0644)
	lg.LogError("Error writing json file", err)
}
