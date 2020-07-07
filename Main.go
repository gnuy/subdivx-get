package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	subdURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	subdPayload string = "mr robot s03e01"
	subdHeaders        = map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0"}
)

func main() {

	subdPayload := strings.ReplaceAll(subdPayload, " ", "%20")

	page := getPage(subdHeaders, subdURL, subdPayload)

	re := regexp.MustCompile("<div id=\"menu_detalle_buscador\">(.|\n)*?<div id=\"menu_detalle_buscador\">")
	lines := re.FindAllString(string(page), -1)

	titleStartIndex := strings.Index(string(page), "subidos")
	titleEndIndex := strings.Index(string(page), "pabloaran")
	pageTitle := []byte(string(page)[titleStartIndex:titleEndIndex])

	fmt.Println(string(pageTitle))

	fmt.Printf("%+q", lines)

}
