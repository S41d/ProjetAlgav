package main

import (
	"fmt"
	"os"
	"projet/cle"
	"projet/tasmin"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
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
			times[i] += fn(jeu[i])
		}
	}
	for i := 0; i < 8; i++ {
		times[i] /= 5
	}
	return times
}

func foreachJeuPair(fn func(string, string) int64) [8]int64 {
	times := [8]int64{}
	for nbJeu := 1; nbJeu < 5; nbJeu++ {
		jeu1 := jeuxDonnes(nbJeu)
		jeu2 := jeuxDonnes(nbJeu + 1)

		for i := 0; i < 8; i++ {
			times[i] += fn(jeu1[i], jeu2[i])
		}
	}
	for i := 0; i < 8; i++ {
		times[i] /= 5
	}
	return times
}

func grapheConstructionTasMinTableau() {
	timesAjoutIteratif := foreachJeu(tabAjoutIteratif)
	timesConstruction := foreachJeu(tabConstruction)
	for i := 0; i < 9; i++ {
		iter := foreachJeu(tabAjoutIteratif)
		cons := foreachJeu(tabConstruction)
		for i := 0; i < 8; i++ {
			timesAjoutIteratif[i] += iter[i]
			timesConstruction[i] += cons[i]
		}
	}

	for i := 0; i < 8; i++ {
		timesAjoutIteratif[i] /= 10
		timesConstruction[i] /= 10
	}

	lineChart := charts.NewLine()
	lineChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Tas Min (Tableau)",
			Subtitle: "Construction vs AjoutsIteratif",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "# itérations"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	lineChart.
		AddSeries(
			"Construction",
			timesToLineData(timesConstruction),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"AjoutIteratif",
			timesToLineData(timesAjoutIteratif),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: true, ShowSymbol: true, SymbolSize: 15, Symbol: "circle"},
		))

	page := components.NewPage()
	page.AddCharts(
		lineChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/construction_tas_min(tableau).html")
	check(page.Render(f))
}

func grapheConstructionTasMinArbre() {
	timesAjoutIteratif := foreachJeu(arbreAjoutIteratif)
	timesConstruction := foreachJeu(arbreConstruction)

	for i := 0; i < 9; i++ {
		iter := foreachJeu(arbreAjoutIteratif)
		cons := foreachJeu(arbreConstruction)
		for i := 0; i < 8; i++ {
			timesAjoutIteratif[i] += iter[i]
			timesConstruction[i] += cons[i]
		}
	}

	for i := 0; i < 8; i++ {
		timesAjoutIteratif[i] /= 10
		timesConstruction[i] /= 10
	}

	lineChart := charts.NewLine()
	lineChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Tas Min (Arbre)",
			Subtitle: "Construction vs AjoutsIteratif",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "# itérations"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	//lineChart.SetXAxis(jeuTailles).
	lineChart.
		AddSeries(
			"Construction",
			timesToLineData(timesConstruction),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"AjoutIteratif",
			timesToLineData(timesAjoutIteratif),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: true, ShowSymbol: true, SymbolSize: 15, Symbol: "circle"},
		))

	page := components.NewPage()
	page.AddCharts(
		lineChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/construction_tas_min(arbre).html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

var jeuTailles = [8]string{
	"1000", "5000", "10000", "20000",
	"50000", "80000", "120000", "200000",
}

func timesToLineData(times [8]int64) []opts.LineData {
	items := make([]opts.LineData, 8)
	for i := 0; i < 8; i++ {
		items[i] = opts.LineData{Value: []interface{}{jeuTailles[i], fmt.Sprint(times[i])}}
	}
	return items
}

func timesToLineDataSlice(times []timeS) []opts.LineData {
	items := make([]opts.LineData, len(times))
	for i := 0; i < len(times); i++ {
		items[i] = opts.LineData{Value: []interface{}{times[i].size, times[i].time}}
	}
	return items
}

func timesToScatterDataSlice(times []timeS) []opts.ScatterData {
	items := make([]opts.ScatterData, len(times))
	for i := 0; i < len(times); i++ {
		items[i] = opts.ScatterData{Value: []interface{}{times[i].size, times[i].time}}
	}
	return items
}

func tabminToGraphNode(index int, tab tasmin.Tableau) opts.TreeData {
	var children []*opts.TreeData
	if tab.EnfGauche(index) < len(tab) {
		enfG := tabminToGraphNode(tab.EnfGauche(index), tab)
		enfG.LineStyle = &opts.LineStyle{Type: "dashed"}
		children = append(children, &enfG)
	}
	if tab.EnfDroit(index) < len(tab) {
		enfD := tabminToGraphNode(tab.EnfDroit(index), tab)
		enfD.LineStyle = &opts.LineStyle{Type: "solid"}
		children = append(children, &enfD)
	}
	style := opts.ItemStyle{
		Color: "lightgray",
	}
	if len(children) == 0 {
		style.Color = "white"
		style.BorderColor = "lightgray"
	}

	node := opts.TreeData{
		Name:       tab[index].DecimalString(),
		Children:   children,
		SymbolSize: 35,
		Symbol:     "roundRect",
		ItemStyle:  &style,
	}
	return node
}

func grapheExempleTasMinTableau() {
	cles := []cle.Cle{
		cle.HexToCle("1f"),
		cle.HexToCle("1e"),
		cle.HexToCle("1d"),
		cle.HexToCle("1c"),
		cle.HexToCle("1b"),
		cle.HexToCle("1a"),
		cle.HexToCle("19"),
		cle.HexToCle("18"),
		cle.HexToCle("17"),
		cle.HexToCle("16"),
		cle.HexToCle("15"),
		cle.HexToCle("14"),
		cle.HexToCle("13"),
		cle.HexToCle("12"),
		cle.HexToCle("11"),
		cle.HexToCle("10"),
		cle.HexToCle("f"),
		cle.HexToCle("e"),
		cle.HexToCle("d"),
		cle.HexToCle("c"),
		cle.HexToCle("b"),
		cle.HexToCle("a"),
		cle.HexToCle("9"),
		cle.HexToCle("8"),
		cle.HexToCle("7"),
		cle.HexToCle("6"),
		cle.HexToCle("5"),
		cle.HexToCle("4"),
		cle.HexToCle("3"),
		cle.HexToCle("2"),
		cle.HexToCle("1"),
	}

	tabmin := tasmin.Construction(cles)

	treeChart := charts.NewTree()
	treeChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Tas Min Tableau",
			Left:  "center",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
	)
	treeChart.AddSeries("tas", []opts.TreeData{tabminToGraphNode(0, tabmin)}).SetSeriesOptions(
		charts.WithTreeOpts(
			opts.TreeChart{
				Layout:           "orthogonal",
				Orient:           "TB",
				InitialTreeDepth: -1,
			},
		),
		charts.WithSeriesAnimation(true),
	)

	page := components.NewPage()
	page.AddCharts(
		treeChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/tas_min_visualization.html")
	check(page.Render(f))
}

func grapheUnionTasMin() {
	timesTab := foreachJeuPair(tabUnion)
	timesArbre := foreachJeuPair(arbreUnion)

	for i := 0; i < 9; i++ {
		iter := foreachJeu(arbreAjoutIteratif)
		cons := foreachJeu(arbreConstruction)
		for j := 0; j < 8; j++ {
			timesTab[j] += iter[j]
			timesArbre[j] += cons[j]
		}
	}

	for i := 0; i < 8; i++ {
		timesTab[i] /= 10
		timesArbre[i] /= 10
	}

	lineChart := charts.NewLine()
	lineChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Union de 2 tas",
			Left:  "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	lineChart.
		AddSeries(
			"Tableau",
			timesToLineData(timesTab),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"Arbre",
			timesToLineData(timesArbre),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: true, ShowSymbol: true, SymbolSize: 15, Symbol: "circle"},
		))

	page := components.NewPage()
	page.AddCharts(
		lineChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/union_tas_min.html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func grapheMd5Construction() {
	timesTas, timesFile := md5Construction()
	m := minTimeS(timesFile, timesTas)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Experimentation",
			Subtitle: "Construction",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	scatterChart.
		AddSeries(
			"Tas (tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"File Binomiale",
			timesToScatterDataSlice(timesFile),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithScatterChartOpts(
			opts.ScatterChart{},
		))

	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/md5_construction.html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func grapheMd5SupprMin() {
	timesTas, timesFile := md5SupprMin()
	m := minTimeS(timesFile, timesTas)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Experimentation",
			Subtitle: "SupprMin",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	scatterChart.
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"File Binomiale",
			timesToScatterDataSlice(timesFile),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithScatterChartOpts(
			opts.ScatterChart{},
		))

	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/md5_suppr_min.html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func grapheMd5Ajout() {
	timesTas, timesFile := md5Ajout()
	m := minTimeS(timesFile, timesTas)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Experimentation",
			Subtitle: "Ajout",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	scatterChart.
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"File Binomiale",
			timesToScatterDataSlice(timesFile),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithScatterChartOpts(
			opts.ScatterChart{},
		))

	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/md5_ajout.html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func grapheMd5Union() {
	timesTas, timesFile := md5Union()
	m := minTimeS(timesFile, timesTas)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Experimentation",
			Subtitle: "Union",
			Left:     "center",
		}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "right"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
	)

	scatterChart.
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"File Binomiale",
			timesToScatterDataSlice(timesFile),
			charts.WithSeriesAnimation(true),
		).
		SetSeriesOptions(charts.WithScatterChartOpts(
			opts.ScatterChart{},
		))

	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create("graphes/md5_union.html")
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func minTimeS(times1, times2 []timeS) int {
	minTime := 0x1111111111111111
	for i := 0; i < len(times1); i++ {
		if times1[i].size < minTime {
			minTime = times1[i].size
		}
		if times2[i].size < minTime {
			minTime = times2[i].size
		}
	}
	return minTime
}
