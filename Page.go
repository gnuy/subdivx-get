package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	subdHeaders = map[string]string{
		"User-Agent":     "🦊Mozilla🦊/5.0 (X11; 🐧Linux🐧 x86_64; rv:68.0) 🦎Gecko🦎/20100101 🔥Firefox🔥/68.0",
		"Accept-Charset": "utf-8",
		"Cookie":         "__cfduid=dea8419e3bf838c5ec1b8624c00ba126e1599785667; con_impr=5; cant_down=16; bajo_una_vez=0; bajo_una_vez_diario=0; contd=3; cs15=566391; cs14=215575; cs13=277494; __cf_bm=edba632a5a68f4ad6890f4f58ff571044dc84d1a-1599793485-1800-Ac0IxmDEnWMbTjrtkhEluRQMTH6hnt2KhSJGCa7KPLxY",
		"Content-Type":   "application/x-www-form-urlencoded",
	}
)

func getRedirectPage(url string) (string, map[string][]string) {

	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range subdHeaders {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	return res.Request.URL.Scheme + "://" + res.Request.URL.Host + res.Request.URL.Path, res.Header
}

func getPage(url string) ([]byte, map[string][]string) {

	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range subdHeaders {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if *verbose {
		fmt.Printf("%s\n%T\n", res.Header, res.Header)
	}

	return body, res.Header
}

func postPage(url string, reqbody []byte) []byte {

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqbody))

	for k, v := range subdHeaders {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if *verbose {
		fmt.Printf("%s\n%T\n", res.Header, res.Header)
	}

	return body
}
