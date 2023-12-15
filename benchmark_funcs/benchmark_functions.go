package benchmark_funcs

import (
	"os"
	"projet/cle"
	"projet/experimentation"
	"projet/filebinomiale"
	"projet/tasmin"
	"slices"
	"strings"
	"time"
)

type Benchmark struct {
	Time int64
	Size int
}

func parseFile(path string) []cle.Cle {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
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

func benchmarkFunc(file string, action func([]cle.Cle)) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	action(cles)
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func TabAjoutIteratif(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tab := tasmin.Tableau{cles[0]}
		tab.AjoutIteratif(cles[1:])
	})
}

func TabConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tasmin.Construction(cles)
	})
}

func TabSupprMin(file string) int64 {
	cles := parseFile(file)
	tab := tasmin.Construction(cles)
	return timeIt(func() {
		tab.SupprMin()
	}).Microseconds()
}

func TabAjout(file string) int64 {
	cles := parseFile(file)
	tab := tasmin.Construction(cles[:len(cles)-2])
	return timeIt(func() {
		tab.Ajout(cles[len(cles)-1])
	}).Microseconds()
}

func ArbreAjoutIteratif(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		arbre := tasmin.NewArbre(cles[0])
		arbre.AjoutIteratif(cles[1:])
	})
}

func ArbreConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		tasmin.ConstructionArbre(cles)
	})
}

func ArbreSupprMin(file string) int64 {
	arbre := tasmin.ConstructionArbre(parseFile(file))
	return timeIt(func() {
		arbre.SupprMin()
	}).Microseconds()
}

func ArbreAjout(file string) int64 {
	var cles = parseFile(file)
	arbre := tasmin.ConstructionArbre(cles[:len(cles)-2])
	return timeIt(func() {
		arbre.Ajout(cles[len(cles)-1])
	}).Microseconds()
}

func TabUnion(file1, file2 string) int64 {
	cles1 := parseFile(file1)
	cles2 := parseFile(file2)

	tab1 := tasmin.Construction(cles1)
	tab2 := tasmin.Construction(cles2)

	tStart := time.Now().UnixMicro()
	tab1.Union(tab2)
	tEnd := time.Now().UnixMicro()

	return tEnd - tStart
}

func ArbreUnion(file1, file2 string) int64 {
	cles1 := parseFile(file1)
	cles2 := parseFile(file2)

	arbre1 := tasmin.ConstructionArbre(cles1)
	arbre2 := tasmin.ConstructionArbre(cles2)

	tStart := time.Now().UnixMicro()
	arbre1.Union(&arbre2)
	tEnd := time.Now().UnixMicro()

	return tEnd - tStart
}

func FileConstruction(file string) int64 {
	return benchmarkFunc(file, func(cles []cle.Cle) {
		filebinomiale.Construction(cles)
	})
}

func FileAjout(file string) int64 {
	var cles = parseFile(file)
	var fb = filebinomiale.Construction(cles[:len(cles)-2])
	return timeIt(func() {
		fb.Ajout(cles[len(cles)-1])
	}).Microseconds()
}

func FileSupprMin(file string) int64 {
	var cles = parseFile(file)
	var fb = filebinomiale.Construction(cles)
	var moyenne = int64(0)
	for i := 0; i < 500; i++ {
		moyenne += timeIt(func() { fb.SupprMin() }).Nanoseconds()
	}
	moyenne /= 500
	return moyenne
}

func FileUnion(file1, file2 string) int64 {
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

func Md5Construction() ([]Benchmark, []Benchmark) {
	data := experimentation.ParseBooks()
	timesTas := make([]Benchmark, len(data))
	timesFile := make([]Benchmark, len(data))
	for i, currData := range data {
		timesTas[i] = Benchmark{
			Time: timeIt(func() { tasmin.Construction(currData) }).Microseconds(),
			Size: len(currData),
		}
		timesFile[i] = Benchmark{
			Time: timeIt(func() { filebinomiale.Construction(currData) }).Microseconds(),
			Size: len(currData),
		}
	}
	return timesTas, timesFile
}

func Md5SupprMin() ([]Benchmark, []Benchmark) {
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

	timesTas := make([]Benchmark, 200)
	timesFile := make([]Benchmark, 200)

	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			t := time.Now().UnixNano()
			tas.SupprMin()
			timesTas[j] = Benchmark{
				Time: time.Now().UnixNano() - t,
				Size: len(tas),
			}
			j++
		} else {
			tas.SupprMin()
		}
	}

	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesFile[j] = Benchmark{
				Time: timeIt(func() { file.SupprMin() }).Nanoseconds(),
				Size: int(file.Size()),
			}
			j++
		} else {
			file.SupprMin()
		}
	}

	return timesTas, timesFile
}

func Md5Ajout() ([]Benchmark, []Benchmark) {
	data := experimentation.ParseBooks()
	tas := tasmin.Construction(data[0])
	var restData []cle.Cle
	for _, d := range data[1:] {
		restData = append(restData, d...)
	}
	file := filebinomiale.Construction(data[0])

	timesTas := make([]Benchmark, 200)
	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			t := time.Now()
			tas.Ajout(restData[i])
			timesTas[j] = Benchmark{
				Time: time.Since(t).Nanoseconds(),
				Size: len(data[0]) + i,
			}
			j++
		} else {
			tas.Ajout(restData[i])
		}
	}

	timesFile := make([]Benchmark, 200)
	for i, j := 0, 0; i < 20000; i++ {
		if i%100 == 0 {
			timesFile[j] = Benchmark{
				Time: timeIt(func() { file.Ajout(restData[i]) }).Nanoseconds(),
				Size: len(data[0]) + i,
			}
			j++
		} else {
			file.Ajout(restData[i])
		}
	}

	return timesTas, timesFile
}

func Md5Union() ([]Benchmark, []Benchmark) {
	data := experimentation.ParseBooks()

	tas := tasmin.Construction(data[0])
	file := filebinomiale.Construction(data[0])

	timesTas := make([]Benchmark, len(data)-1)
	for i := 1; i < len(data); i++ {
		toAdd := tasmin.Construction(data[i])
		timesTas[i-1] = Benchmark{
			Time: timeIt(func() { tas.Union(toAdd) }).Nanoseconds(),
			Size: len(tas) + len(data[i]),
		}
	}

	timesFile := make([]Benchmark, len(data)-1)
	for i := 1; i < len(data); i++ {
		toAdd := filebinomiale.Construction(data[i])
		size := file.Size()
		timesFile[i-1] = Benchmark{
			Time: timeIt(func() { file = file.Union(toAdd) }).Nanoseconds(),
			Size: int(size) + len(data[i]),
		}
	}

	return timesTas, timesFile
}
