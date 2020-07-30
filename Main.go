package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var (
	listURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	listPayload string = "mr robot s03e01"
	regex              = map[string]string{
		"filterList": "<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>",
		"getLink":    "<a class=\"titulo_menu_izq\" href=\"(.)*?\">"}
	subEjemplo string = "<div id=\"menu_detalle_buscador\"><div id=\"menu_titulo_buscador\"><a class=\"titulo_menu_izq\" href=\"http://www.subdivx.com/X6XNTIwNDUyX-mr-robot-s03e01.html\">Subtitulos de Mr. Robot S03E01</a></div><img src=\"img/calif2.gif\" class=\"detalle_calif\" name=\"detalle_calif\"></div><div id=\"buscador_detalle\">\n<div id=\"buscador_detalle_sub\">subidos por pabloaran y modificados para que coincidan con el capítulo de popcorn time  mr robot s03e01 eps3 0 power-saver-mode h 1080p amzn web-dl ddp5 1 h 264-ntb</div><div id=\"buscador_detalle_sub_datos\"><b>Downloads:</b> 1,046 <b>Cds:</b> 1 <b>Comentarios:</b> <a rel=\"nofollow\" href=\"popcoment.php?idsub=NTIwNDUy\" onclick=\"return hs.htmlExpand(this, { objectType: 'iframe' } )\">2</a> <b>Formato:</b> SubRip <b>Subido por:</b> <a class=\"link1\" href=\"http://www.subdivx.com/X9X1991429\">pincharata</a> <img src=\"/pais/1.gif\" width=\"16\" height=\"12\"> <b>el</b> 29/11/2017 </div></div>"
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

func extract(line string, field string) string {
	re := regexp.MustCompile(regex[field])
	strParts := strings.Split(regex[field], "(.)*?")
	raw := re.FindString(line)
	parsedValue := raw[len(strParts[0]) : len(raw)-len(strParts[1])]
	return parsedValue

}

func main() {
	// page := getPage("http://www.subdivx.com/X6XNTIwNDUyX-mr-robot-s03e01.html")
	// fmt.Printf("%s", page)
	file := getPage("http://www.subdivx.com/bajar.php?id=520452&u=8")
	err := ioutil.WriteFile("file", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
	archiverunzip("file", "/tmp")
	// fmt.Printf("%s", file)
}

func toUtf8(iso8859_1buf []byte) string {
	buf := make([]rune, len(iso8859_1buf))
	for i, b := range iso8859_1buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

// func main() {

// 	listPayload := strings.ReplaceAll(listPayload, " ", "%20")

// 	fmt.Println(listURL + listPayload + "\n" + subdHeaders["User-Agent"])
// 	page := getPage(listURL + listPayload)

// 	re := regexp.MustCompile(regex["filterList"])
// 	lines := re.FindAll(page, -1)

// 	fmt.Println("\n\n", len(lines))

// 	for i := 0; i < len(lines); i++ {
// 		fmt.Printf("%s%s%s\n\n", "⭐", extract(toUtf8(lines[i]), "getLink"), "⭐")
// 	}

// }
