package files_binomiales

import "projet/cle"

type TournoiBinomial struct {
	Cle     *cle.Cle
	Parent  *TournoiBinomial
	Enfants FileBinomiale
	Degre   int
}

func (tb TournoiBinomial) EstVide() bool {
	return tb.Cle == nil
}

func (tb TournoiBinomial) Union(o TournoiBinomial) TournoiBinomial {
	assert(tb.Degre == o.Degre, "degre inégal")
	tb.Enfants = append(tb.Enfants, o)
	return tb
}

func (tb TournoiBinomial) Decapite() FileBinomiale {
	assert(!tb.EstVide(), "tb vide")
	return tb.Enfants
}

func (tb TournoiBinomial) File() FileBinomiale {
	assert(!tb.EstVide(), "tb vide")
	return FileBinomiale{tb}
}

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}
