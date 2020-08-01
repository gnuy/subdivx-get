package main

import (
	"regexp"
	"strings"
)

var (
	regex = map[string]string{
		"filterList": "<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>",
		"getLink":    "<a class=\"titulo_menu_izq\" href=\"(.)*?\">",
		"getDesc":    "<div id=\"buscador_detalle_sub\">(.)*?</div>",
		"getCountry": "src=\"/pais/(.)*?.gif"}
)

func extract(line string, field string) string {
	re := regexp.MustCompile(regex[field])
	strParts := strings.Split(regex[field], "(.)*?")
	raw := re.FindString(line)
	parsedValue := raw[len(strParts[0]) : len(raw)-len(strParts[1])]
	return parsedValue
}

func getList(page []byte) [][]byte {
	re := regexp.MustCompile(regex["filterList"])
	lines := re.FindAll(page, -1)
	return lines
}

func toUtf8(iso8859_1buf []byte) string {
	buf := make([]rune, len(iso8859_1buf))
	for i, b := range iso8859_1buf {
		buf[i] = rune(b)
	}
	return string(buf)
}
