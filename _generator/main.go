package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/benweidig/runes/_generator/models"
)

func main() {

	fmt.Print("Reading Blocks... ")
	blocksFile, err := os.Open("data/Blocks.txt")
	if err != nil {
		panic(err)
	}
	defer blocksFile.Close()

	var blocks []models.Block
	scanner := bufio.NewScanner(blocksFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		block := models.NewBlock(line)
		blocks = append(blocks, block)
	}
	fmt.Printf("%d found!\n", len(blocks))

	fmt.Print("Reading Unicode Data... ")

	dataFile, err := os.Open("data/UnicodeData.txt")
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	pointsByBlock := make(map[models.Block][]models.CodePoint, len(blocks))
	reader := csv.NewReader(dataFile)
	reader.Comma = ';'

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		cp := models.NewCodePoint(line)
		// Detect Block
		for _, block := range blocks {
			if block.Range().IsIn(cp.Point()) {
				pointsByBlock[block] = append(pointsByBlock[block], cp)
				continue
			}
		}
	}
	fmt.Println("Done!")

	// 3. Write packages
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"raw": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("template.tml"))

	for block, codepoints := range pointsByBlock {
		fmt.Printf("Writing Block '%s'... ", block.Name())

		packageName := block.PackageName()
		path := fmt.Sprintf("../%s/", packageName)
		packagepath, _ := filepath.Abs(fmt.Sprintf("%s/%s.go", path, packageName))

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
		f, err := os.Create(packagepath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		tmpl.ExecuteTemplate(f, "template.tml", models.NewTemplateData(block, codepoints))

		generatedMarker, _ := filepath.Abs(fmt.Sprintf("%s/.generated", path))
		os.OpenFile(generatedMarker, os.O_RDONLY|os.O_CREATE, 0666)
		fmt.Println("Done!")
	}

}
