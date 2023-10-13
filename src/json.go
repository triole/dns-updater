package main

import (
	"encoding/json"
	"os"

	"github.com/triole/logseal"
)

func readIPDataJSON() (ipd tIPDataSet) {
	_, err := os.Stat(fileIPDataJSON)
	if os.IsNotExist(err) {
		lg.Info(
			"ip date json does not exist. consider ip as changed", nil,
		)
	} else {
		raw, err := os.ReadFile(fileIPDataJSON)
		lg.IfErrError("read", fileIPDataJSON, err, false)

		err = json.Unmarshal(raw, &ipd)
		lg.IfErrError("can not unmarshal", logseal.F{
			"data":  fileIPDataJSON,
			"error": err,
		})
	}
	return
}

func writeIPDataJSON(ipd tIPDataSet) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = os.WriteFile(fileIPDataJSON, JSONstring, 0644)
	lg.IfErrError("unable to write file",
		logseal.F{
			"data":  fileIPDataJSON,
			"error": err,
		})
}
