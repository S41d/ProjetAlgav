package main

import (
	"fmt"
	"os"
	"projet/cle"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseFile(path string) []cle.Cle {
	fileContent, err := os.ReadFile(path)
	check(err)
	strContent := string(fileContent)
	lines := strings.Split(strContent, "\n")
	var cles []cle.Cle
	for _, s := range lines {
		if s != "" {
			cles = append(cles, cle.HexToCle(s))
		}
	}
	return cles
}

func jeuxDonnes(nbJeu int) []string {
	jeu := []string{
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_1000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_5000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_10000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_20000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_50000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_80000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_120000.txt",
		"fichiers_tests/jeu_" + fmt.Sprint(nbJeu) + "_nb_cles_200000.txt",
	}
	return jeu
}

func foreachJeu(fn func(string) int64) [8]int64 {
	times := [8]int64{}
	for nbJeu := 1; nbJeu <= 5; nbJeu++ {
		jeu := jeuxDonnes(nbJeu)
		for i := 0; i < 8; i++ {
			times[i] += fn(jeu[i])
		}
	}
	for i := 0; i < 8; i++ {
		times[i] /= 5
	}
	return times
}

func main() {
	grapheConstructionTasMinTableau()
	grapheConstructionTasMinArbre()
}
