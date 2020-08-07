package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	maxLengthDesc        = 100
	listURL       string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	listPayload   string = "mr robot s03e01" //deshardcdear, pasar por parámetro
	//sacar, ésto es uno de los elementos de la lista en getList(getPage(listURL + listPayload))
)

type subElement struct {
	link      string
	desc      string
	country   string // falta mappear número con país
	downloads string
	format    string
	uploader  string
	score     string
	date      string
}

// func main() {
// 	// page := getPage("http://www.subdivx.com/X6XNTIwNDUyX-mr-robot-s03e01.html")
// 	// fmt.Printf("%s", page)
// 	file := getPage("http://www.subdivx.com/bajar.php?id=520452&u=8")
// 	err := ioutil.WriteFile("file", file, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	unzip("file", "/tmp")
// 	// fmt.Printf("%s", file)
// }

func populateElement(line []byte) subElement {
	return subElement{
		link:      getLink(line),
		desc:      getDesc(line),
		country:   getCountry(line),
		downloads: getDownloads(line),
		format:    getFormat(line),
		uploader:  getUploader(line),
		score:     getScore(line),
		date:      getDate(line),
	}
}

func createTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	// tbl := table.New("ID", "Description", "Country", "Downloads", "Format", "Uploader", "Score", "Date")
	tbl := table.New("ID", "Description", "Country", "Downloads", "Uploader", "Score")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func trimString(value string, length int) string {
	if len(value) > length {
		value = value[:length]
	}
	return value
}

func main() {
	elements := []subElement{}
	tbl := createTable()
	listPayload := strings.ReplaceAll(listPayload, " ", "%20")

	fmt.Println(listURL + listPayload + "\n" + subdHeaders["User-Agent"])
	lines := getList(getPage(listURL + listPayload))

	for i := 0; i < len(lines); i++ {
		elements = append(elements, populateElement(lines[i]))
		// tbl.AddRow(i, getDesc(lines[i]), getCountry(lines[i]), getDownloads(lines[i]), getFormat(lines[i]), getUploader(lines[i]), getScore(lines[i])+"⭐", getDate(lines[i]))
		tbl.AddRow(i, trimString(getDesc(lines[i]), maxLengthDesc), getCountry(lines[i]), getDownloads(lines[i]), getUploader(lines[i]), getScore(lines[i])+"⭐")
	}
	tbl.Print()

	subPage := getPage(elements[3].link)         // Hay que mostrar lista y dar a elegir el nro de elemento
	subFile := getPage(getDownloadLink(subPage)) // Download sub

	err := ioutil.WriteFile("file", subFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
	unzip("file", ".")

}
