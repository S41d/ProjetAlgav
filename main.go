package main

import (
	"fmt"
	"projet/experimentation"
)

func main() {
	buildArbreExempleCharts()
	buildTabExempleCharts()
	buildFileExempleCharts()

	buildTasCharts()
	buildFileCharts()
	buildMd5Charts()

	var _, words = experimentation.ParseBooksABR()
	fmt.Println("nb mots:", len(words))
}
