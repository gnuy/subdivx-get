package main

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	maxLengthTitle = 30
	maxLengthDesc  = 80
	regex          = map[string]string{
		"filterList": "<div id=\"menu_detalle_buscador\">(.|\n)*?</div></div>",

		"getLink":          "<a class=\"titulo_menu_izq\" href=\"(.|\n)*?\">",
		"getDesc":          "<div id=\"buscador_detalle_sub\">(.|\n)*?</div>",
		"getCountry":       "src=\"/pais/(.|\n)*?.gif",
		"getScore":         "src=\"img/calif(.|\n)*?.gif",
		"getFormat":        "<b>Formato:</b> (.|\n)*? <b>",
		"getDate":          "<b>el</b> (.|\n)*? </div>",
		"getDownloads":     "<b>Downloads:</b> (.|\n)*? <b>",
		"getUploaderStep1": "<b>Subido por:</b> <a class=(.|\n)*?\">(.|\n)*?</a>",
		"getUploaderStep2": "\">(.|\n)*?</",
		"getTitle":         ">Subtitulos de (.|\n)*?</a>",

		"getDownloadLink":   "<a class=\"link1\" href=\"(.|\n)*?\">Bajar",
		"getDownloadLinkId": "id=(.|\n)*?&",
	}
)

func getTitle(line []byte) string {
	return trimString(extract(toUtf8(line), "getTitle"), maxLengthTitle)
}

func getDownloadLinkID(line string) string {
	return extract(line, "getDownloadLinkId")
}

func getDownloadLink(line []byte) string {
	return extract(toUtf8(line), "getDownloadLink")
}

func getLink(line []byte) string {
	return extract(toUtf8(line), "getLink")
}

func getDesc(line []byte) string {
	return trimString(extract(toUtf8(line), "getDesc"), maxLengthDesc)
}

func getCountry(line []byte) string {
	return extract(toUtf8(line), "getCountry")
}

func getScore(line []byte) string {
	return extract(toUtf8(line), "getScore")
}

func getFormat(line []byte) string {
	return extract(toUtf8(line), "getFormat")
}

func getDate(line []byte) string {
	return extract(toUtf8(line), "getDate")
}

func getDownloads(line []byte) string {
	return extract(toUtf8(line), "getDownloads")
}

func getUploader(line []byte) string {
	return extract(extract(toUtf8(line), "getUploaderStep1"), "getUploaderStep2")
}

func extract(line string, field string) string {
	re := regexp.MustCompile(regex[field])
	strParts := strings.Split(regex[field], "(.|\n)*?")
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

func createFileTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Archivo")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func createTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Título", "Descripción", "Usuario", "Calif.")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func trimString(value string, length int) string {
	if len(value) > length {
		value = value[:length]
	}
	return value
}
