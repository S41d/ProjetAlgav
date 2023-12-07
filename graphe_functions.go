package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

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

	lineChart.SetXAxis(jeuTailles).
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

	lineChart.SetXAxis(jeuTailles).
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
	f, _ := os.Create("graphes/construction_tas_min(arbre).html")
	check(page.Render(f))
}

var jeuTailles = [8]string{
	"1000", "5000", "10000", "20000", "50000",
	"80000", "120000", "200000",
}

func timesToLineData(times [8]int64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 8; i++ {
		items = append(items, opts.LineData{Value: times[i]})
	}
	return items
}
