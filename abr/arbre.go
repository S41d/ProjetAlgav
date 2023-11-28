package arbrerecherche

import "projet/cle"

type ArbreRecherche struct {
	Elt      *cle.Cle
	Parent   *ArbreRecherche
	SaGauche *ArbreRecherche
	SaDroit  *ArbreRecherche
}

func (a *ArbreRecherche) Contient(c cle.Cle) bool {
	if c.Eg(*a.Elt) {
		return true
	} else if c.Inf(*a.Elt) {
		return a.SaGauche.Contient(c)
	} else {
		return a.SaDroit.Contient(c)
	}
}

func (a *ArbreRecherche) Ajout(c cle.Cle) {
	if a.Elt == nil {
		a.Elt = &c
	} else if c.Eg(*a.Elt) {
		return
	} else if c.Inf(*a.Elt) {
		a.SaGauche.Ajout(c)
	} else {
		a.SaDroit.Ajout(c)
	}
}

func (a *ArbreRecherche) Suppr(c cle.Cle) {
	if a.Elt == nil {
		return
	} else if c.Eg(*a.Elt) {
		a.Elt = nil
	} else if c.Inf(*a.Elt) {
		a.SaGauche.Suppr(c)
	} else {
		a.SaDroit.Suppr(c)
	}
}

func (a *ArbreRecherche) EstArbreVide() bool {
	return a.Elt == nil
}

func (a *ArbreRecherche) Racine() ArbreRecherche {
	curr := a
	for curr.Parent != nil {
		curr = curr.Parent
	}
	return *curr
}

func NewArbreRecherche(cles []cle.Cle) ArbreRecherche {
	a := ArbreRecherche{}
	for i := 0; i < len(cles); i++ {
		a.Ajout(cles[i])
	}
	return a
}
