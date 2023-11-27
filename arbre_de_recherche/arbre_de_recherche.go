package arbre_de_recherche

import "projet/cle"

type ArbreRecherche struct {
	cle   *cle.Cle
	left  *ArbreRecherche
	right *ArbreRecherche
}

func (a *ArbreRecherche) contient(c cle.Cle) bool {
	if c.Eg(*a.cle) {
		return true
	} else if c.Inf(*a.cle) {
		return a.left.contient(c)
	} else {
		return a.right.contient(c)
	}
}

func (a *ArbreRecherche) insert(c cle.Cle) {
	if a.cle == nil {
		a.cle = &c
	} else if c.Inf(*a.cle) {
		a.left.insert(c)
	} else {
		a.right.insert(c)
	}
}

func NewArbreRecherche(cles []cle.Cle) ArbreRecherche {
	a := ArbreRecherche{}
	for i := 0; i < len(cles); i++ {
		a.insert(cles[i])
	}
	return a
}
