package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
var (
	eggplant    string = "🍆"
	watermelon  string = "🍉"
	subdURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	subdPayload string = "Mr.Robot"
	subdHeader  string = "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0"
)
*/
func main() {
	// url := "https://reqres.in/api/users"
	var url string = "http://www.subdivx.com/index.php?q=mr%20robot%20s03&accion=5&masdesc=&subtitulos=1&realiza_b=1"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	/*	fmt.Println("Hello Go! 🇺🇾 " + eggplant + " " + watermelon)
		// http.NewRequest("GET", subdURL, reader)

		req, _ := http.NewRequest("GET", subdURL, nil)
		// req.Header.Add("content-type", "application/x-www-form-urlencoded")
		// req.Header.Add("cache-control", "no-cache")
		req.Header.Add("User-Agent", subdHeader)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
			println(err)
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))

	*/

}
