package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getPage(headers map[string]string, url string, payload string) []byte {

	req, _ := http.NewRequest("GET", url+payload, nil)

	req.Header.Add("User-Agent", headers["User-Agent"])

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body
}
