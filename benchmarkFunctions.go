package main

import (
	"projet/tasMin"
	"time"
)

func tabAjoutIteratif(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	tab := tasMin.Tableau{cles[0]}
	tab.AjoutIteratif(cles[1:])
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func tabConstruction(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	tasMin.Construction(cles)
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func arbreAjoutIteratif(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	arbre := tasMin.NewArbre(cles[0])
	arbre.AjoutIteratif(cles[1:])
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}
