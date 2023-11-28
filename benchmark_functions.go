package main

import (
	"projet/tasmin"
	"time"
)

func tabAjoutIteratif(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	tab := tasmin.Tableau{cles[0]}
	tab.AjoutIteratif(cles[1:])
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func tabConstruction(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	tasmin.Construction(cles)
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}

func arbreAjoutIteratif(file string) int64 {
	cles := parseFile(file)
	tStart := time.Now().UnixMicro()
	arbre := tasmin.NewArbre(cles[0])
	arbre.AjoutIteratif(cles[1:])
	tEnd := time.Now().UnixMicro()
	return tEnd - tStart
}
