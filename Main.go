package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	listURL     string = "http://www.subdivx.com/index.php?accion=5&q="
	listPayload []string
	subPosition = flag.Int("n", -1, "help message for flag n")
	reader      = bufio.NewReader(os.Stdin)
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

func getUserInput() int {
	fmt.Print("\nSeleccioná el ID del sub -> ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSuffix(value, "\n")
	intvalue, err := strconv.Atoi(value)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return intvalue

}

func main() {
	flag.Parse()
	listPayload = flag.Args()
	elements := []subElement{}
	tbl := createTable()
	listPayload := strings.ReplaceAll(fmt.Sprint(listPayload), " ", "%20")

	lines := getList(getPage(listURL + listPayload))

	for i := 0; i < len(lines); i++ {
		elements = append(elements, populateElement(lines[i]))
		tbl.AddRow(i, trimString(getDesc(lines[i]), maxLengthDesc), getDownloads(lines[i]), getUploader(lines[i]), getScore(lines[i])+"⭐")
	}

	if len(elements) > 0 {
		if *subPosition == -1 {
			tbl.Print()
			*subPosition = getUserInput()
		}
		subPage := getPage(elements[*subPosition].link)
		subFile := getPage(getDownloadLink(subPage)) // Download sub

		writefile := ioutil.WriteFile("file", subFile, 0644)
		if writefile != nil {
			log.Fatal(writefile)
		}
		unzip("file", ".")
	} else {
		fmt.Println("No se encontraron subs.")
	}
}
