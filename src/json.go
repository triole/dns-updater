package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/triole/logseal"
)

func (conf *tConf) readIPDataJSON() (ipd tIPDataSet) {
	lg.Debug("read data json", logseal.F{
		"path": conf.DataJSONFile,
	})
	_, err := os.Stat(conf.DataJSONFile)
	if os.IsNotExist(err) {
		lg.Info(
			"ip date json does not exist. consider ip as changed", nil,
		)
	} else {
		raw, err := os.ReadFile(conf.DataJSONFile)
		lg.IfErrError("read", conf.DataJSONFile, err, false)
		if err == nil {
			err = json.Unmarshal(raw, &ipd)
			lg.IfErrError("can not unmarshal", logseal.F{
				"path":  conf.DataJSONFile,
				"error": err,
			})
			if err == nil {
				lg.Debug("data json info", logseal.F{
					"path":    conf.DataJSONFile,
					"content": fmt.Sprintf("%+v", ipd),
				})
			}
		}
	}
	return
}

func (conf *tConf) writeIPDataJSON(ipd tIPDataSet) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = os.WriteFile(conf.DataJSONFile, JSONstring, 0644)
	lg.IfErrError("unable to write file",
		logseal.F{
			"data":  conf.DataJSONFile,
			"error": err,
		})
}
