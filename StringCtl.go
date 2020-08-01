package main

import (
	"regexp"
	"strings"
)

var (
	regex = map[string]string{
		//Extract list of subs from html
		"filterList": "<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>",
		//Extract each field
		"getLink":          "<a class=\"titulo_menu_izq\" href=\"(.)*?\">",
		"getDesc":          "<div id=\"buscador_detalle_sub\">(.)*?</div>",
		"getCountry":       "src=\"/pais/(.)*?.gif",
		"getScore":         "src=\"img/calif(.)*?.gif",
		"getFormat":        "<b>Formato:</b> (.)*? <b>",
		"getDate":          "<b>el</b> (.)*? </div>",
		"getDownloads":     "<b>Downloads:</b> (.)*? <b>",
		"getUploader":      "<b>Subido por:</b> <a class=(.)*?\">(.)*?</a>",
		"getUploaderStep1": "<b>Subido por:</b> <a class=(.)*?\">(.)*?</a>",
		"getUploaderStep2": "\">(.)*?</"}
)

func getUploader(line []byte) string {
	return extract(extract(toUtf8(line), "getUploaderStep1"), "getUploaderStep2")
}

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
