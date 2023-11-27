package files_binomiales

import "projet/cle"

type FileBinomiale []TournoiBinomial

func (fb FileBinomiale) EstVide() bool {
	return len(fb) == 0
}

func (fb FileBinomiale) MinDeg() TournoiBinomial {
	return fb[0]
}

func (fb FileBinomiale) Reste() FileBinomiale {
	return fb[1:]
}

func (fb FileBinomiale) AjoutMin(tb TournoiBinomial) FileBinomiale {
	assert(tb.Degre < fb.MinDeg().Degre,
		"le degre de tournoi n'est pas inf à MinDeg de la file")
	newFb := FileBinomiale{tb}
	newFb = append(newFb, fb...)
	return newFb
}

func (fb FileBinomiale) Union(o FileBinomiale) FileBinomiale {
	return uFret(fb, o, TournoiBinomial{})
}

func (fb FileBinomiale) Ajout(c cle.Cle) FileBinomiale {
	toAdd := FileBinomiale{TournoiBinomial{Cle: &c}}
	return fb.Union(toAdd)
}

func (fb FileBinomiale) SupprMin() FileBinomiale {
	assert(!fb.EstVide(), "SupprMin sur file vide")
	minimal := &fb[0]
	for i := 1; i < len(fb); i++ {
		curr := &fb[i]
		if minimal.Cle.Inf(*curr.Cle) {
			minimal = curr
		}
	}
	decapite := minimal.Decapite()
	return fb.Union(decapite)
}

func Construction(cles [] cle.Cle) FileBinomiale {
	fb := FileBinomiale{}
	for i := 0; i < len(cles); i++ {
		fb.Ajout(cles[i])	
	}
	return fb
}

func uFret(f1, f2 FileBinomiale, t TournoiBinomial) FileBinomiale {
	if t.EstVide() {
		if f1.EstVide() {
			return f2
		}
		if f2.EstVide() {
			return f1
		}

		t1 := f1.MinDeg()
		t2 := f2.MinDeg()
		if t1.Degre < t2.Degre {
			return f2.Union(f1.Reste()).AjoutMin(t1)
		}
		if t2.Degre < t1.Degre {
			return f1.Union(f2.Reste()).AjoutMin(t2)
		}
		if t1.Degre == t2.Degre {
			return uFret(f1.Reste(), f2.Reste(), t1.Union(t2))
		}
	} else {
		if f1.EstVide() {
			return f2.Union(t.File())
		}
		if f2.EstVide() {
			return f1.Union(t.File())
		}

		t1 := f1.MinDeg()
		t2 := f2.MinDeg()
		if t.Degre < t1.Degre && t.Degre < t2.Degre {
			return f1.Union(f2).AjoutMin(t)
		}
		if t.Degre == t1.Degre && t.Degre == t2.Degre {
			return uFret(f1.Reste(), f2.Reste(), t1.Union(t2)).AjoutMin(t)
		}
		if t.Degre == t1.Degre && t.Degre < t2.Degre {
			return uFret(f1.Reste(), f2, t1.Union(t))
		}
		if t.Degre == t2.Degre && t.Degre < t1.Degre {
			return uFret(f2.Reste(), f1, t2.Union(t))
		}
	}
	panic("end of uFret, should not reach")
}
