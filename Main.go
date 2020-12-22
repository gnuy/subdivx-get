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
	listURL      string = "http://www.subdivx.com/index.php?accion=5&q="
	subdivxURL   string = "https://www.subdivx.com/"
	inputArgs    []string
	listPayload  []string
	subDir       []os.FileInfo
	subPosition  = flag.Int("n", -1, "número de sub en la lista")
	fileLocation = flag.String("l", ".", "ubicación de los subs en el filesystem")
	verbose      = flag.Bool("v", false, "modo verboso")
	reader       = bufio.NewReader(os.Stdin)
)

type subElement struct {
	link      string
	desc      string
	country   string // falta mappear número con país
	downloads string
	format    string
	uploader  string
	score     string
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

func selectFile(targetDir string) string {

	filePosition := 0
	subDir = ls(targetDir)
	filetable := createFileTable()
	for i, f := range subDir {
		filetable.AddRow(i, f.Name())
		if *verbose {
			fmt.Println("'" + targetDir + "/" + f.Name() + "'")
		}
		if i > 0 {
			filePosition = i
		}
	}

	if filePosition > 0 {
		filetable.Print()
		filePosition = getUserInput()
	}
	for i, f := range subDir {
		if i == filePosition {
			return "'" + targetDir + "/" + f.Name() + "'"
		}
	}

	return os.DevNull
}

func ls(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func main() {
	flag.Parse()
	inputArgs = flag.Args()
	elements := []subElement{}
	tbl := createTable()
	listPayload := strings.ReplaceAll(fmt.Sprint(inputArgs), " ", "%20")
	lines := getList(getPage(listURL + listPayload))

	if *verbose {
		fmt.Printf("\n%s\n", lines)
	}

	for i := 0; i < len(lines); i++ {
		elements = append(elements, populateElement(lines[i]))
		tbl.AddRow(i, getTitle(lines[i]), getDesc(lines[i]),
			getDownloads(lines[i]), getUploader(lines[i]), getScore(lines[i])+"⭐")
		if *verbose {
			fmt.Printf("\n%s\n", elements)
		}
	}

	if len(elements) > 0 {
		if *subPosition == -1 { // Workaround de que el debugger se tranca en el getUserInput()
			tbl.Print()
			*subPosition = getUserInput()
		}
		subPage := getPage(elements[*subPosition].link)
		downloadLink := getDownloadLink(subPage)
		downloadLinkID := getDownloadLinkID(downloadLink)
		if *verbose {
			fmt.Println("downloadLink: " + downloadLink)
			fmt.Println("downloadLinkID: " + downloadLinkID)
		}
		targetDir := *fileLocation + "/" + downloadLinkID
		subFile := getPage(subdivxURL + downloadLink) // Download sub
		os.Mkdir(downloadLinkID, 0700)
		tempFile := targetDir + "/subdivx-get.tmp"
		writefile := ioutil.WriteFile(tempFile, subFile, 0644)
		if writefile != nil {
			log.Fatal(writefile)
		}
		unzip(tempFile, targetDir)
		os.RemoveAll(tempFile)

		selectedFile := selectFile(targetDir)
		fmt.Println(selectedFile)

	} else {
		fmt.Println("No se encontraron subs.")
	}
}
