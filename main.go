package main

import (
	"fmt"
	"os"
	"projet/cle"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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
			times[i] = fn(jeu[i])
		}
	}
	return times
}

func timesToBarData(times [8]int64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 8; i++ {
		items = append(items, opts.BarData{
			Value:   times[i],
			Tooltip: &opts.Tooltip{Show: true},
		})
	}
	return items
}

func main() {
	timesAjoutIteratif := foreachJeu(tabAjoutIteratif)
	timesConstruction := foreachJeu(tabConstruction)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Tableau",
			Subtitle: "Construction vs AjoutsIteratif",
			Left:     "center",
		}),
		//charts.WithToolboxOpts(opts.Toolbox{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "# itérations"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
	)
	jeuTailles := []string{
		"1000", "5000", "10000", "20000", "50000",
		"80000", "120000", "200000",
	}
	bar.SetXAxis(jeuTailles).
		AddSeries("Construction", timesToBarData(timesConstruction)).
		AddSeries("AjoutIteratif", timesToBarData(timesAjoutIteratif))
	f, _ := os.Create("tableau_construction.html")
	check(bar.Render(f))
}
