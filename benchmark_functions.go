package main

import (
	"projet/cle"
	"projet/experimentation"
	"projet/filebinomiale"
	"projet/tasmin"
	"slices"
	"time"
)

func benchmarkFunc(file string, action func([]cle.Cle)) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	action(cles)
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func tabAjoutIteratif(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tab := tasmin.Tableau{cles[0]}
		tab.AjoutIteratif(cles[1:])
	})
}

func tabConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tasmin.Construction(cles)
	})
}

func arbreAjoutIteratif(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		arbre := tasmin.NewArbre(cles[0])
		arbre.AjoutIteratif(cles[1:])
	})
}

func arbreConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tasmin.ConstructionArbre(cles)
	})
}

func tabUnion(file1, file2 string) int64 {
	cles1 := parseFile(file1)
	cles2 := parseFile(file2)

	tab1 := tasmin.Construction(cles1)
	tab2 := tasmin.Construction(cles2)

	tStart := time.Now().UnixMicro()
	tab1.Union(tab2)
	tEnd := time.Now().UnixMicro()

	return tEnd - tStart
}

func arbreUnion(file1, file2 string) int64 {
	cles1 := parseFile(file1)
	cles2 := parseFile(file2)

	arbre1 := tasmin.ConstructionArbre(cles1)
	arbre2 := tasmin.ConstructionArbre(cles2)

	tStart := time.Now().UnixMicro()
	arbre1.Union(&arbre2)
	tEnd := time.Now().UnixMicro()

	return tEnd - tStart
}

func fileConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		filebinomiale.Construction(cles)
	})
}

func fileUnion(file1, file2 string) int64 {
	cles1 := parseFile(file1)
	cles2 := parseFile(file2)

	fb1 := filebinomiale.Construction(cles1)
	fb2 := filebinomiale.Construction(cles2)

	tStart := time.Now().UnixMicro()
	fb1.Union(fb2)
	tEnd := time.Now().UnixMicro()

	return tEnd - tStart
}

func timeIt(f func()) int64 {
	t := time.Now().UnixMicro()
	f()
	return time.Now().UnixMicro() - t
}

func findMin(tab tasmin.Tableau) cle.Cle {
	m := tab[0]
	for i := 1; i < len(tab)-1; i++ {
		if tab[i].Inf(m) {
			m = tab[i]
		}
	}
	return m
}

func timeItNano(f func()) int64 {
	t := time.Now().UnixNano()
	f()
	return time.Now().UnixNano() - t
}

type timeS struct {
	time int64
	size int
}

func md5Construction() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()
	var timesTas []timeS
	var timesFile []timeS
	for _, currData := range data {
		timesTas = append(timesTas, timeS{time: timeIt(func() { tasmin.Construction(currData) }), size: len(currData)})
		timesFile = append(timesFile, timeS{time: timeIt(func() { filebinomiale.Construction(currData) }), size: len(currData)})
	}
	return timesTas, timesFile
}

func md5SupprMin() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()

	var allData []cle.Cle
	for _, currData := range data {
		for i := 0; i < len(currData); i++ {
			if !slices.Contains(allData, currData[i]) {
				allData = append(allData, currData[i])
			}
		}
	}

	tas := tasmin.Construction(allData)
	file := filebinomiale.Construction(data[0])

	var timesTas []timeS
	var timesFile []timeS

	for i := 0; i < 20000; i++ {
		timesTas = append(timesTas, timeS{
			time: timeItNano(func() { tas.SupprMin() }),
			size: len(data[0]) + 1 + i,
		})
	}

	for i := 0; i < 2000; i++ {
		timesFile = append(timesFile, timeS{
			time: timeItNano(func() { file.SupprMin() }),
			size: len(data[0]) + 1 + i,
		})
	}

	return timesTas, timesFile
}

func md5Ajout() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()
	tas := tasmin.Construction(data[0])
	var restData []cle.Cle
	for _, d := range data[1:] {
		restData = append(restData, d...)
	}
	file := filebinomiale.Construction(data[0])

	var timesTas []timeS
	for i := 0; i < 2000; i++ {
		timesTas = append(timesTas, timeS{time: timeItNano(func() { tas.Ajout(restData[i]) }), size: len(data[0]) + i})
	}

	var timesFile []timeS
	for i := 0; i < 2000; i++ {
		timesFile = append(timesFile, timeS{time: timeItNano(func() { file.Ajout(restData[i]) }), size: len(data[0]) + i})
	}

	return timesTas, timesFile
}

func md5Union() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()

	tas := tasmin.Construction(data[0])
	file := filebinomiale.Construction(data[0])

	var timesTas []timeS
	for i := 1; i < len(data); i++ {
		toAdd := tasmin.Construction(data[i])
		timesTas = append(timesTas, timeS{time: timeItNano(func() { tas.Union(toAdd) }), size: len(tas) + len(data[i])})
	}

	var timesFile []timeS
	for i := 1; i < len(data); i++ {
		toAdd := filebinomiale.Construction(data[i])
		timesFile = append(timesFile, timeS{time: timeItNano(func() { file.Union(toAdd) }), size: len(data[i])})
	}

	return timesTas, timesFile
}
