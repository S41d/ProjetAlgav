package tasMin

import (
	"errors"
	"projet/cle"
	"strings"
)

type Tableau []cle.Cle

func (*Tableau) Parent(i int) int {
	return (i - 1) / 2
}

func (*Tableau) EnfGauche(i int) int {
	return 2*i + 1
}

func (*Tableau) EnfDroit(i int) int {
	return 2*i + 2
}

func (t *Tableau) SupprMin() (cle.Cle, error) {
	if len(*t) == 0 {
		return cle.Cle{}, errors.New("SupprMin sur tas vide")
	}
	cleSupprime := (*t)[0]
	(*t)[0] = (*t)[len(*t)-1]
	t.trier(0)
	return cleSupprime, nil
}

func (t *Tableau) Ajout(c cle.Cle) {
	*t = append(*t, c)
	currI := len(*t) - 1
	parentI := t.Parent(currI)
	currInfParent := (*t)[currI].Inf((*t)[parentI])

	for currI > 0 && currInfParent {
		tmp := (*t)[parentI]
		(*t)[parentI] = (*t)[currI]
		(*t)[currI] = tmp

		currI = parentI
		parentI = t.Parent(currI)
		currInfParent = (*t)[currI].Inf((*t)[parentI])
	}
}

func (t *Tableau) AjoutIteratif(cles []cle.Cle) {
	for i := 0; i < len(cles); i++ {
		t.Ajout(cles[i])
	}
}

func Construction(cles []cle.Cle) Tableau {
	var t Tableau = cles
	// derniers noeuds avec enfants
	start := (len(cles) / 2) - 1

	for i := start; i >= 0; i-- {
		t.trier(i)
	}
	return t
}

func Union(t1 Tableau, t2 Tableau) Tableau {
	t := append(t1, t2...)
	return Construction(t)
}

func (t *Tableau) trier(index int) {
	if len(*t) <= 1 {
		return
	}

	gauche := t.EnfGauche(index)
	droite := t.EnfDroit(index)

	minimum := index
	if gauche < len(*t) && (*t)[gauche].Inf((*t)[index]) {
		minimum = gauche
	}

	if droite < len(*t) && (*t)[droite].Inf((*t)[minimum]) {
		minimum = droite
	}

	if minimum != index {
		tmp := (*t)[index]
		(*t)[index] = (*t)[minimum]
		(*t)[minimum] = tmp
		t.trier(minimum)
	}
}

func (t *Tableau) _string(i int, indent int) string {
	str := "Tas{\n"
	indent += 1
	// str += strings.Repeat("│ ", indent) + "cle: " + (*t)[i].DecimalString() + ",\n"
	str += strings.Repeat("  ", indent) + "cle: " + (*t)[i].DecimalString() + ",\n"
	leftI := t.EnfGauche(i)
	rightI := t.EnfDroit(i)
	if leftI < len(*t) {
		str += strings.Repeat("  ", indent) + "left: " + t._string(leftI, indent) + ",\n"
	} else {
		str += strings.Repeat("  ", indent) + "left: nil,\n"
	}
	if rightI < len(*t) {
		str += strings.Repeat("  ", indent) + "right: " + t._string(rightI, indent) + "\n"
	} else {
		str += strings.Repeat("  ", indent) + "right: nil\n"
	}
	str += strings.Repeat("  ", indent-1) + "}"
	return str
}

func (t *Tableau) String() string {
	return t._string(0, 0) + "\n"
}
