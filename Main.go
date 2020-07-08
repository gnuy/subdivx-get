package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	listURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	listPayload string = "mr robot s03e01"
	subdHeaders        = map[string]string{
		"User-Agent": "🦊Mozilla🦊/5.0 (X11; 🐧Linux🐧 x86_64; rv:68.0) 🦎Gecko🦎/20100101 🔥Firefox🔥/68.0"}
	star string = "⭐"
)

type subElement struct {
	link      string
	desc      string
	country   string
	downloads string
	format    string
	uploader  string
	score     string
}

func main() {

	listPayload := strings.ReplaceAll(listPayload, " ", "%20")

	fmt.Println(listURL + listPayload + "\n" + subdHeaders["User-Agent"])
	page := getPage(subdHeaders, listURL, listPayload)

	re := regexp.MustCompile("<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>")
	lines := re.FindAllString(string(page), -1)

	fmt.Println("\n\n", len(lines))

	for i := 0; i < len(lines); i++ {
		fmt.Printf("%s%+q%s\n\n", star, lines[i], star)
	}

}
