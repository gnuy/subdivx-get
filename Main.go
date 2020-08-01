package main

import (
	"fmt"
	"strings"
)

var (
	listURL     string = "http://www.subdivx.com/index.php?accion=5&masdesc=&subtitulos=1&realiza_b=1&q="
	listPayload string = "mr robot s03e01" //deshardcdear, pasar por parámetro
	//sacar, ésto es uno de los elementos de la lista en getList(getPage(listURL + listPayload))
	subEjemplo string = "<div id=\"menu_detalle_buscador\"><div id=\"menu_titulo_buscador\"><a class=\"titulo_menu_izq\" href=\"http://www.subdivx.com/X6XNTIwNDUyX-mr-robot-s03e01.html\">Subtitulos de Mr. Robot S03E01</a></div><img src=\"img/calif2.gif\" class=\"detalle_calif\" name=\"detalle_calif\"></div><div id=\"buscador_detalle\">\n<div id=\"buscador_detalle_sub\">subidos por pabloaran y modificados para que coincidan con el capítulo de popcorn time  mr robot s03e01 eps3 0 power-saver-mode h 1080p amzn web-dl ddp5 1 h 264-ntb</div><div id=\"buscador_detalle_sub_datos\"><b>Downloads:</b> 1,046 <b>Cds:</b> 1 <b>Comentarios:</b> <a rel=\"nofollow\" href=\"popcoment.php?idsub=NTIwNDUy\" onclick=\"return hs.htmlExpand(this, { objectType: 'iframe' } )\">2</a> <b>Formato:</b> SubRip <b>Subido por:</b> <a class=\"link1\" href=\"http://www.subdivx.com/X9X1991429\">pincharata</a> <img src=\"/pais/1.gif\" width=\"16\" height=\"12\"> <b>el</b> 29/11/2017 </div></div>"
)

type subElement struct {
	link      string //√
	desc      string //√
	country   string //√ falta mappear número con país
	downloads string //√
	format    string //√
	uploader  string
	score     string //√
	date      string //√
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

func main() {

	listPayload := strings.ReplaceAll(listPayload, " ", "%20")

	fmt.Println(listURL + listPayload + "\n" + subdHeaders["User-Agent"])
	lines := getList(getPage(listURL + listPayload))

	fmt.Println("\n\n", len(lines))

	// popular un []subElement con todos los campos correspondientes`
	for i := 0; i < len(lines); i++ {
		//		fmt.Printf("%s%s%s\n\n", "⭐", extract(toUtf8(lines[i]), "getLink"), "⭐")
		fmt.Printf("%s%s%s\n\n", "⭐", extract(toUtf8(lines[i]), "getDownloads"), "⭐")

		// fmt.Printf("%s%s%s\n\n", "⭐", toUtf8(lines[i]), "⭐")
	}

}
