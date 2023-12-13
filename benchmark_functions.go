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

func timeIt(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
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

type timeS struct {
	time int64
	size int
}

func md5Construction() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()
	timesTas := make([]timeS, len(data))
	timesFile := make([]timeS, len(data))
	for i, currData := range data {
		timesTas[i] = timeS{
			time: timeIt(func() { tasmin.Construction(currData) }).Microseconds(),
			size: len(currData),
		}
		timesFile[i] = timeS{
			time: timeIt(func() { filebinomiale.Construction(currData) }).Microseconds(),
			size: len(currData),
		}
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
	file := filebinomiale.Construction(allData)

	timesTas := make([]timeS, 200)
	timesFile := make([]timeS, 200)

	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesTas[j] = timeS{
				time: timeIt(func() { tas.SupprMin() }).Nanoseconds(),
				size: len(tas),
			}
			j++
		} else {
			tas.SupprMin()
		}
	}

	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesFile[j] = timeS{
				time: timeIt(func() { file = file.SupprMin() }).Nanoseconds(),
				size: int(file.Size()),
			}
			j++
		} else {
			file.SupprMin()
		}
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

	timesTas := make([]timeS, 200)
	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesTas[j] = timeS{
				time: timeIt(func() { tas.Ajout(restData[i]) }).Nanoseconds(),
				size: len(data[0]) + i,
			}
			j++
		} else {
			tas.Ajout(restData[i])
		}
	}

	timesFile := make([]timeS, 200)
	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesFile[j] = timeS{
				time: timeIt(func() { file.Ajout(restData[i]) }).Nanoseconds(),
				size: len(data[0]) + i,
			}
			j++
		} else {
			file.Ajout(restData[i])
		}
	}

	return timesTas, timesFile
}

func md5Union() ([]timeS, []timeS) {
	data := experimentation.ParseBooks()

	tas := tasmin.Construction(data[0])
	file := filebinomiale.Construction(data[0])

	timesTas := make([]timeS, len(data)-1)
	for i := 1; i < len(data); i++ {
		toAdd := tasmin.Construction(data[i])
		timesTas[i-1] = timeS{
			time: timeIt(func() { tas.Union(toAdd) }).Nanoseconds(),
			size: len(tas) + len(data[i]),
		}
	}

	timesFile := make([]timeS, len(data)-1)
	for i := 1; i < len(data); i++ {
		toAdd := filebinomiale.Construction(data[i])
		size := file.Size()
		timesFile[i-1] = timeS{
			time: timeIt(func() { file = file.Union(toAdd) }).Nanoseconds(),
			size: int(size) + len(data[i]),
		}
	}

	return timesTas, timesFile
}
