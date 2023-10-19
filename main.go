package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

func seedMap(rows [][]string) map[string]string {
	m := make(map[string]string)

	for _, row := range rows {
		fields := strings.Split(row[0], ";")
		if len(fields) >= 2 {
			m[fields[0]] = fields[1]
		}
	}

	return m
}

func enterToClose() {
	fmt.Println("Aperte 'enter' para sair")
	fmt.Scanln()
	os.Exit(1)
}

func openFile(fileName string) (file *os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", fileName)
		enterToClose()
	}

	return file
}

func main() {
	file := openFile("base.csv")
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler arquivo base")
		enterToClose()
	}

	m := seedMap(rows)

	var xmlFileName string
	fmt.Println("Digite o nome do arquivo xml: ")
	fmt.Scanln(&xmlFileName)

	if !strings.HasSuffix(xmlFileName, ".xml") {
		xmlFileName += ".xml"
	}

	f := openFile(xmlFileName)
	defer f.Close()

	doc, err := xmlquery.Parse(f)
	if err != nil {
		fmt.Println("Erro ao ler arquivo xml")
		enterToClose()
	}

	nodes := xmlquery.Find(doc, "//cProd")
	nNF := xmlquery.FindOne(doc, "//nNF")

	for _, node := range nodes {
		code := node.InnerText()

		if value, ok := m[code]; ok {
			node.FirstChild.Data = value
		}
	}

	updatedFileName := nNF.InnerText() + ".xml"

	xmlFile, err := os.Create(updatedFileName)
	if err != nil {
		fmt.Println("Erro ao criar arquivo xml")
		enterToClose()
	}
	defer xmlFile.Close()

	_, err = xmlFile.WriteString(doc.OutputXML(true))
	if err != nil {
		fmt.Println("Erro ao escrever arquivo xml")
		enterToClose()
	}

	fmt.Println(doc.OutputXML(true))
}
