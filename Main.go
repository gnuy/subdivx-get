package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	subdURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	subdPayload string = "mr robot s03"
	subdHeaders        = map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0"}
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

func main() {

	subdPayload := strings.ReplaceAll(subdPayload, " ", "%20")

	page := getPage(subdHeaders, subdURL, subdPayload)

	titleStartIndex := strings.Index(string(page), "<Title>")
	titleEndIndex := strings.Index(string(page), "</title>")
	pageTitle := []byte(string(page)[titleStartIndex:titleEndIndex])
	fmt.Println(string(pageTitle))

}
