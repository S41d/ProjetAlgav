package main

import (
	"projet/cle"
	"projet/tasmin"
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
