package main

import (
	"encoding/json"
	"os"
)

func readIPDataJSON() (ipd tIPDataSet) {
	_, err := os.Stat(fileIPDataJSON)
	if os.IsNotExist(err) {
		lg.LogInfo(
			"ip date json does not exist. consider ip as changed", nil,
		)
	} else {
		raw, err := os.ReadFile(fileIPDataJSON)
		lg.LogIfFileError("read", fileIPDataJSON, err, false)

		err = json.Unmarshal(raw, &ipd)
		lg.LogIfFileError("unmarshal", fileIPDataJSON, err, false)
	}
	return
}

func writeIPDataJSON(ipd tIPDataSet) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = os.WriteFile(fileIPDataJSON, JSONstring, 0644)
	lg.LogIfFileError("write", fileIPDataJSON, err, false)
}
