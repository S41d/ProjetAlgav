package main

import (
	"fmt"
	"projet/experimentation"
)

func main() {
	fmt.Println("visualisations: ")
	buildArbreExempleCharts()
	fmt.Printf(" %-10s  🮱\n", "arbre")
	buildTabExempleCharts()
	fmt.Printf(" %-10s  🮱\n", "tableau")
	buildFileExempleCharts()
	fmt.Printf(" %-10s  🮱\n", "file")

	fmt.Println("graphes: ")
	buildTasCharts()
	fmt.Printf(" %-10s  🮱\n", "tas")
	buildFileCharts()
	fmt.Printf(" %-10s  🮱\n", "file")
	buildMd5Charts()
	fmt.Printf(" %-10s  🮱\n", "md5")

	var _, words = experimentation.ParseBooksABR()
	fmt.Println("nb mots:", len(words))
}
