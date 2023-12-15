package main

import (
	"fmt"
	"github.com/sajari/regression"
	"os"
	bf "projet/benchmark_funcs"
	"projet/cle"
	"projet/filebinomiale"
	"projet/tasmin"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var jeuTailles = [8]int{
	1000, 5000, 10000, 20000,
	50000, 80000, 120000, 200000,
}

func buildTasCharts() {
	// tableau construction
	ajoutVsConsScatterBuilder(
		"graphes/tas_min/construction_tableau.html",
		"Tas Min (Tableau)",
		foreachJeu(bf.TabAjoutIteratif),
		foreachJeu(bf.TabConstruction),
	)

	//
	ajoutVsConsScatterBuilder(
		"graphes/tas_min/construction_arbre.html",
		"Tas Min (Arbre)",
		foreachJeu(bf.ArbreAjoutIteratif),
		foreachJeu(bf.ArbreConstruction),
	)

	tabVsArbreScatterBuilder(
		"graphes/tas_min/suppr_min.html",
		"Tas Min SupprMin",
		foreachJeu(bf.TabSupprMin),
		foreachJeu(bf.ArbreSupprMin),
	)

	tabVsArbreScatterBuilder(
		"graphes/tas_min/ajout.html",
		"Tas Min Ajout",
		foreachJeu(bf.TabAjout),
		foreachJeu(bf.ArbreAjout),
	)

	tabVsArbreScatterBuilder(
		"graphes/tas_min/union.html",
		"Tas Min Union",
		foreachJeuPair(bf.TabUnion),
		foreachJeuPair(bf.ArbreUnion),
	)
}

func buildFileCharts() {
	var minSize = func(times []bf.Benchmark) int {
		smallest := 0x1111111111111111
		for i := 0; i < len(times); i++ {
			if times[i].Size < smallest {
				smallest = times[i].Size
			}
		}
		return smallest
	}
	var builder = func(path string, title string, times []bf.Benchmark) {
		reg := createRegressionInstance(times)
		regPoints := regressionPoints(reg, times)

		scatterChartBuilder(path, title, minSize(times), []dataSeries{
			{name: "Regression", typ: "line", data: regPoints},
			{name: "Tableau", typ: "scatter", data: timesToScatterData(times)},
		})
	}

	builder(
		"graphes/file_binomiale/construction.html",
		"File Binomiale Construction",
		foreachJeu(bf.FileConstruction),
	)

	builder(
		"graphes/file_binomiale/ajout.html",
		"File Binomiale Ajout",
		foreachJeu(bf.FileAjout),
	)

	builder(
		"graphes/file_binomiale/suppr_min.html",
		"File Binomiale SupprMin",
		foreachJeu(bf.FileSupprMin),
	)

	builder(
		"graphes/file_binomiale/union.html",
		"File Binomiale Union",
		foreachJeuPair(bf.FileUnion),
	)
}

func buildMd5Charts() {
	fileVsTasScatterBuilder("graphes/md5/ajout.html", "MD5 Ajout", bf.Md5Ajout)
	fileVsTasScatterBuilder("graphes/md5/union.html", "MD5 Union", bf.Md5Union)
	fileVsTasScatterBuilder("graphes/md5/suppr_min.html", "MD5 SupprMin", bf.Md5SupprMin)
	fileVsTasScatterBuilder("graphes/md5/construction.html", "MD5 Construction", bf.Md5Construction)
}

func arbreToTreeData(a tasmin.Arbre) opts.TreeData {
	var children []*opts.TreeData
	if a.EnfGauche != nil {
		enfG := arbreToTreeData(*a.EnfGauche)
		enfG.LineStyle = &opts.LineStyle{Type: "dashed"}
		children = append(children, &enfG)
	}
	if a.EnfDroit != nil {
		enfD := arbreToTreeData(*a.EnfDroit)
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
		Name:       a.Cle.DecimalString(),
		Children:   children,
		SymbolSize: 35,
		Symbol:     "roundRect",
		ItemStyle:  &style,
	}
	return node
}

func tabminToTreeData(index int, tab tasmin.Tableau) opts.TreeData {
	var children []*opts.TreeData
	if tab.EnfGauche(index) < len(tab) {
		enfG := tabminToTreeData(tab.EnfGauche(index), tab)
		enfG.LineStyle = &opts.LineStyle{Type: "dashed"}
		children = append(children, &enfG)
	}
	if tab.EnfDroit(index) < len(tab) {
		enfD := tabminToTreeData(tab.EnfDroit(index), tab)
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

func fileToTreeData(file filebinomiale.FileBinomiale) []*opts.TreeData {
	var children []*opts.TreeData
	if !file.EstVide() {
		for i := 0; i < len(file); i++ {
			children2 := fileToTreeData(file[i].Enfants)
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
	treeChart.AddSeries("tas", []opts.TreeData{tabminToTreeData(0, tabmin)}).SetSeriesOptions(
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

func createRegressionInstance(times []bf.Benchmark) *regression.Regression {
	reg := new(regression.Regression)
	for i := 0; i < len(times); i++ {
		reg.Train(
			regression.DataPoint(float64(times[i].Time), []float64{float64(times[i].Size)}),
		)
	}
	check(reg.Run())
	return reg
}

func regressionPoints(reg *regression.Regression, times []bf.Benchmark) [][]interface{} {
	var points [][]interface{}
	for _, t := range times {
		predictionTas, _ := reg.Predict([]float64{float64(t.Size)})
		points = append(points, []interface{}{t.Size, max(predictionTas, 0)})
	}
	return points
}

func minSizeTimeS(times1, times2 []bf.Benchmark) int {
	minTime := 0x1111111111111111
	for i := 0; i < len(times1); i++ {
		if times1[i].Size < minTime {
			minTime = times1[i].Size
		}
		if times2[i].Size < minTime {
			minTime = times2[i].Size
		}
	}
	return minTime
}

func treeChartBuilder(data opts.TreeData) *charts.Tree {
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

type dataSeries struct {
	name string
	typ  string
	data interface{}
}

var cles = []cle.Cle{
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

func grapheExempleSupprMinArbre() {
	arbre := tasmin.ConstructionArbre(cles)
	var charters []components.Charter

	for i := 0; i < 10; i++ {
		aTreeData := arbreToTreeData(arbre)
		charters = append(charters, treeChartBuilder(aTreeData))
		arbre.SupprMin()
	}

	page := components.NewPage()
	page.AddCharts(charters...)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/arbre_suppr_min_visualization.html")
	check(page.Render(f))
}

func grapheExempleSupprMinTab() {
	tab := tasmin.Construction(cles)
	var charters []components.Charter

	for i := 0; i < 10; i++ {
		tabTreeData := tabminToTreeData(0, tab)
		charters = append(charters, treeChartBuilder(tabTreeData))
		tab.SupprMin()
	}

	page := components.NewPage()
	page.AddCharts(charters...)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/tableau_suppr_min_visualization.html")
	check(page.Render(f))
}

func grapheExempleSupprMinFile() {
	fb := filebinomiale.Construction(cles)
	var charters []components.Charter

	for i := 0; i < 10; i++ {
		charters = append(charters, treeChartBuilder(opts.TreeData{Children: fileToTreeData(fb)}))
		fb.SupprMin()
	}

	page := components.NewPage()
	page.AddCharts(charters...)

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/file_suppr_min_visualization.html")
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
	chart1 := treeChartBuilder(opts.TreeData{Children: fileToTreeData(fb1)})
	fb2 := filebinomiale.Construction(cles2)
	chart2 := treeChartBuilder(opts.TreeData{Children: fileToTreeData(fb2)})

	union := fb1.Union(fb2)
	chart3 := treeChartBuilder(opts.TreeData{Children: fileToTreeData(union)})

	page := components.NewPage()
	page.AddCharts(chart1, chart2, chart3)
	page.PageTitle = "Exemple file union"

	page.SetLayout(components.PageCenterLayout)
	f, _ := os.Create("graphes/file_exemple_union_visualization.html")
	check(page.Render(f))
}

func tabVsArbreScatterBuilder(
	path string,
	title string,
	timesTab []bf.Benchmark,
	timesArbre []bf.Benchmark,
) {
	m := minSizeTimeS(timesTab, timesArbre)

	regTab := createRegressionInstance(timesTab)
	regArbre := createRegressionInstance(timesArbre)
	regTabPoints := regressionPoints(regArbre, timesTab)
	regArbrePoints := regressionPoints(regTab, timesArbre)

	scatterChartBuilder(path, title, m, []dataSeries{
		{name: "Regression arbre", typ: "line", data: regArbrePoints},
		{name: "Arbre", typ: "scatter", data: timesToScatterData(timesArbre)},
		{name: "Regression Tab", typ: "line", data: regTabPoints},
		{name: "Tableau", typ: "scatter", data: timesToScatterData(timesTab)},
	})
}

func ajoutVsConsScatterBuilder(
	path string,
	title string,
	timesAjoutIteratif []bf.Benchmark,
	timesConstruction []bf.Benchmark,
) {
	m := minSizeTimeS(timesAjoutIteratif, timesConstruction)

	regAIter := createRegressionInstance(timesAjoutIteratif)
	regCons := createRegressionInstance(timesConstruction)
	regAIterPoints := regressionPoints(regAIter, timesAjoutIteratif)
	regConsPoints := regressionPoints(regCons, timesConstruction)

	scatterChartBuilder(path, title, m, []dataSeries{
		{name: "Regression Ajout", typ: "line", data: regAIterPoints},
		{name: "Ajouts Itératifs", typ: "scatter", data: timesToScatterData(timesAjoutIteratif)},
		{name: "Regression Cons", typ: "line", data: regConsPoints},
		{name: "Construction", typ: "scatter", data: timesToScatterData(timesConstruction)},
	})
}

func fileVsTasScatterBuilder(path string, title string, fn func() ([]bf.Benchmark, []bf.Benchmark)) {
	timesTas, timesFile := fn()
	m := minSizeTimeS(timesFile, timesTas)

	regTas := createRegressionInstance(timesTas)
	regFile := createRegressionInstance(timesFile)
	regTasPoints := regressionPoints(regTas, timesTas)
	regFilePoints := regressionPoints(regFile, timesFile)

	scatterChartBuilder(path, title, m, []dataSeries{
		{name: "Regression Tas", typ: "line", data: regTasPoints},
		{name: "Tas (Tableau)", typ: "scatter", data: timesToScatterData(timesTas)},
		{name: "Regression File", typ: "line", data: regFilePoints},
		{name: "File Binomiale", typ: "scatter", data: timesToScatterData(timesFile)},
	})
}

func scatterChartBuilder(filepath string, title string, minX int, data []dataSeries) {
	scatterChart := charts.NewScatter()
	scatterChart.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithXAxisOpts(opts.XAxis{Name: "taille des tas", Min: minX}),
		charts.WithYAxisOpts(opts.YAxis{Name: "temps(en µs)"}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "inside"}),
	)

	for i := 0; i < len(data); i++ {
		scatterChart.AddSeries(
			data[i].name,
			[]opts.ScatterData{},
			func(s *charts.SingleSeries) {
				s.Type = data[i].typ
				s.Animation = true
				s.Data = data[i].data
			},
		)
	}
	page := components.NewPage()
	page.AddCharts(
		scatterChart,
	)
	page.PageTitle = title

	page.SetLayout(components.PageCenterLayout)
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	check(page.Render(f))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
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

func foreachJeu(fn func(string) int64) []bf.Benchmark {
	times := make([]bf.Benchmark, 8)
	for nbJeu := 1; nbJeu <= 5; nbJeu++ {
		jeu := jeuxDonnes(nbJeu)
		for i := 0; i < 8; i++ {
			times[i] = bf.Benchmark{Time: times[i].Time + fn(jeu[i]), Size: jeuTailles[i]}
		}
	}
	for i := 0; i < 8; i++ {
		times[i].Time /= 5
	}
	return times
}

func foreachJeuPair(fn func(string, string) int64) []bf.Benchmark {
	times := make([]bf.Benchmark, 8)
	for nbJeu := 1; nbJeu < 5; nbJeu++ {
		jeu1 := jeuxDonnes(nbJeu)
		jeu2 := jeuxDonnes(nbJeu + 1)

		for i := 0; i < 8; i++ {
			times[i] = bf.Benchmark{Time: times[i].Time + fn(jeu1[i], jeu2[i]), Size: jeuTailles[i]}
		}
	}
	for i := 0; i < 8; i++ {
		times[i].Time /= 5
	}
	return times
}

func timesToScatterData(times []bf.Benchmark) []opts.ScatterData {
	items := make([]opts.ScatterData, len(times))
	for i := 0; i < len(times); i++ {
		items[i] = opts.ScatterData{Value: []interface{}{times[i].Size, times[i].Time}}
	}
	return items
}
