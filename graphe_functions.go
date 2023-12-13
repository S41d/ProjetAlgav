package main

import (
	"fmt"
	"github.com/sajari/regression"
	"os"
	"projet/cle"
	"projet/filebinomiale"
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
	page.PageTitle = "Tas Min (Tableau) - Construction"

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

	page.PageTitle = "Tas Min (Arbre) - Construction"
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

func fileToTreeNode(file filebinomiale.FileBinomiale) []*opts.TreeData {
	var children []*opts.TreeData
	if !file.EstVide() {
		for i := 0; i < len(file); i++ {
			children2 := fileToTreeNode(file[i].Enfants)
			children = append(children, &opts.TreeData{
				Name:       file[i].Cle.DecimalString(),
				Children:   children2,
				Symbol:     "roundRect",
				SymbolSize: 25,
			})
		}
	}
	return children
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
	page.PageTitle = "Visu Tas min"

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

	regTas := createRegressionInstance(timesTas)
	regFile := createRegressionInstance(timesFile)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: "Md5 Construction"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside"}),
	)

	regTasPoints := regressionPoints(regTas, timesTas)
	regFilePoints := regressionPoints(regFile, timesFile)

	scatterChart.
		AddSeries(
			"Regression Tas",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regTasPoints
			},
		).
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"Regression File",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regFilePoints
			},
		).
		AddSeries(
			"File Binomiale",
			timesToScatterDataSlice(timesFile),
			charts.WithSeriesAnimation(true),
		)

	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)
	page.PageTitle = "MD5 Construction"

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

	regTas := createRegressionInstance(timesTas)
	regFile := createRegressionInstance(timesFile)
	regTasPoints := regressionPoints(regTas, timesTas)
	regFilePoints := regressionPoints(regFile, timesFile)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: "MD5 SupprMin"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside"}),
	)

	scatterChart.
		AddSeries(
			"Regression Tas",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regTasPoints
			},
		).
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"Regression File",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regFilePoints
			},
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
	page.PageTitle = "MD5 SupprMin"

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

	regTas := createRegressionInstance(timesTas)
	regFile := createRegressionInstance(timesFile)
	regTasPoints := regressionPoints(regTas, timesTas)
	regFilePoints := regressionPoints(regFile, timesFile)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: "MD5 Ajout"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside"}),
	)

	scatterChart.
		AddSeries(
			"Regression Tas",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				// s.Animation = true
				s.Data = regTasPoints
			},
		).
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"Regression File",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				// s.Animation = true
				s.Data = regFilePoints
			},
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
	page.PageTitle = "MD5 Ajout"

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

	regTas := createRegressionInstance(timesTas)
	regFile := createRegressionInstance(timesFile)
	regTasPoints := regressionPoints(regTas, timesTas)
	regFilePoints := regressionPoints(regFile, timesFile)

	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: "MD5 Union"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: m}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside"}),
	)

	scatterChart.
		AddSeries(
			"Regression Tas",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regTasPoints
			},
		).
		AddSeries(
			"Tas (Tableau)",
			timesToScatterDataSlice(timesTas),
			charts.WithSeriesAnimation(true),
		).
		AddSeries(
			"Regression File",
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = "line"
				s.Animation = true
				s.Data = regFilePoints
			},
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
	page.PageTitle = "MD5 Union"

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

func createRegressionInstance(times []timeS) *regression.Regression {
	reg := new(regression.Regression)
	for i := 0; i < len(times); i++ {
		reg.Train(
			regression.DataPoint(float64(times[i].time), []float64{float64(times[i].size)}),
		)
	}
	check(reg.Run())
	return reg
}

func regressionPoints(reg *regression.Regression, times []timeS) [][]interface{} {
	var points [][]interface{}
	for _, t := range times {
		predictionTas, _ := reg.Predict([]float64{float64(t.size)})
		points = append(points, []interface{}{t.size, predictionTas})
	}
	return points
}

func createFileTreeChart(file filebinomiale.FileBinomiale) *charts.Tree {
	var data opts.TreeData
	data.Name = ""
	data.Children = fileToTreeNode(file)

	treeChart := charts.NewTree()
	treeChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Height: "800px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "File binomiale",
			Left:  "center",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
	)
	treeChart.AddSeries("tas", []opts.TreeData{data}).SetSeriesOptions(
		charts.WithTreeOpts(
			opts.TreeChart{
				Layout:           "orthogonal",
				Orient:           "TB",
				InitialTreeDepth: -1,
			},
		),
		charts.WithSeriesAnimation(true),
	)
	return treeChart
}

func grapheExempleSupprMinFile() {
	cles := []cle.Cle{
		// cle.HexToCle("1f"),
		// cle.HexToCle("1e"),
		// cle.HexToCle("1d"),
		// cle.HexToCle("1c"),
		// cle.HexToCle("1b"),
		// cle.HexToCle("1a"),
		// cle.HexToCle("19"),
		// cle.HexToCle("18"),
		// cle.HexToCle("17"),
		// cle.HexToCle("16"),
		// cle.HexToCle("15"),
		// cle.HexToCle("14"),
		// cle.HexToCle("13"),
		// cle.HexToCle("12"),
		// cle.HexToCle("11"),
		// cle.HexToCle("10"),
		// cle.HexToCle("f"),
		// cle.HexToCle("e"),
		// cle.HexToCle("d"),
		// cle.HexToCle("c"),
		// cle.HexToCle("b"),
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
	fb := filebinomiale.Construction(cles)
	var charters []components.Charter

	for i := 0; i < 10; i++ {
		charters = append(charters, createFileTreeChart(fb))
		fb = fb.SupprMin()
	}

	page := components.NewPage()
	page.AddCharts(charters...)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/file_visualization.html")
	check(page.Render(f))
}

func grapheExempleFileUnion() {
	cles1 := []cle.Cle{
		cle.HexToCle("5"),
		cle.HexToCle("4"),
		cle.HexToCle("3"),
		cle.HexToCle("2"),
		cle.HexToCle("1"),
	}

	cles2 := []cle.Cle{
		cle.HexToCle("a"),
		cle.HexToCle("9"),
		cle.HexToCle("8"),
		cle.HexToCle("7"),
		cle.HexToCle("6"),
	}

	fb1 := filebinomiale.Construction(cles1)
	chart1 := createFileTreeChart(fb1)
	fb2 := filebinomiale.Construction(cles2)
	chart2 := createFileTreeChart(fb2)

	union := fb1.Union(fb2)
	chart3 := createFileTreeChart(union)

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.PageTitle = "Exemple file union"

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/file_exemple_union_visualization.html")
	check(page.Render(f))
}
