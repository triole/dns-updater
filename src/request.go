package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeUpdateRequest(url string) {
	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		lg.Fatal("Error initializing the request", err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(hostname, token))
	response, err := client.Do(req)
	if err != nil {
		lg.Fatal("An error occured during the update request ", err)
	}
	if response.StatusCode == 200 {
		lg.Log("Update successful. Response code was 200.")
	} else {
		lg.Log("Update failed. Response code was " + strconv.Itoa(response.StatusCode))
	}
}
