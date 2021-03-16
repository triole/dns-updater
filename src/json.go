package main

import (
	"encoding/json"
	"io/ioutil"
)

func readIPDataJSON() (ipd IPData) {
	raw, err := ioutil.ReadFile(fileIPDataJSON)
	lg.LogIfErr("Error reading ip data json: %q", err)
	err = json.Unmarshal(raw, &ipd)
	lg.LogIfErr("Can not unmarshal ip data json: %q", err)
	return
}

func writeIPDataJSON(ipd IPData) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = ioutil.WriteFile(fileIPDataJSON, JSONstring, 0644)
	lg.LogIfErr("Error writing json file: %q", err)

}
