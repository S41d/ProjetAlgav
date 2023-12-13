package filebinomiale

import (
	"math"
	"projet/cle"
	"slices"
)

type FileBinomiale []TournoiBinomial

// EstVide retourne true si la file binomiale est vide, c'est-à-dire si elle ne contient aucun nœud de tournoi binomial.
// Sinon, elle retourne false.
func (fb *FileBinomiale) EstVide() bool {
	// Vérifie si la longueur de la file binomiale est égale à zéro.
	return len(*fb) == 0
}

// MinDeg renvoie le nœud de tournoi binomial avec le degré minimum dans la file binomiale.
func (fb *FileBinomiale) MinDeg() TournoiBinomial {
	// La file binomiale est supposée être non vide, sinon cette méthode pourrait entraîner un comportement indéfini.
	// Dans un scénario de production, on devrait probablement ajouter des vérifications supplémentaires.

	// Renvoie le premier nœud de la file binomiale, qui a le degré minimum.
	return (*fb)[0]
}

// Reste renvoie une nouvelle file binomiale contenant tous les éléments à l'exception du premier.
func (fb *FileBinomiale) Reste() FileBinomiale {
	// La file binomiale est supposée être non vide, sinon cette méthode pourrait entraîner un comportement indéfini.
	// Dans un scénario de production, on devrait probablement ajouter des vérifications supplémentaires.

	// Renvoie une nouvelle file binomiale contenant tous les éléments à partir du deuxième.
	return (*fb)[1:]
}

func (fb *FileBinomiale) Size() uint {
	var sum uint
	for i := 0; i < len(*fb); i++ {
		sum += uint(math.Pow(2, float64((*fb)[i].Degre)))
	}
	return sum
}

// AjoutMin ajoute un nœud de tournoi binomial avec un degré inférieur au degré du premier nœud de la file binomiale.
// Elle renvoie une nouvelle file binomiale résultant de cet ajout.
func (fb *FileBinomiale) AjoutMin(tb TournoiBinomial) FileBinomiale {
	// Vérifie que le degré du tournoi est inférieur au degré du premier nœud de la file.
	assert(tb.Degre < fb.MinDeg().Degre, "le degre de tournoi n'est pas inférieur à MinDeg de la file")

	// Crée une nouvelle file binomiale avec le tournoi ajouté en tête.
	newFb := FileBinomiale{tb}

	// Ajoute tous les éléments de la file binomiale originale après le nouveau nœud.
	newFb = append(newFb, *fb...)

	// Renvoie la nouvelle file binomiale résultant de cet ajout.
	return newFb
}

// Union réalise l'opération d'union entre deux files binomiales.
// Elle prend une autre file binomiale en argument et renvoie une nouvelle file binomiale résultant de l'union.
func (fb *FileBinomiale) Union(o FileBinomiale) FileBinomiale {
	// Utilise une fonction auxiliaire uFret pour effectuer l'union.
	// La TournoiBinomial{} est utilisée comme nœud sentinel pour l'union initiale.
	return uFret(*fb, o, TournoiBinomial{})
}

// Ajout ajoute une nouvelle clé à la file binomiale.
// Elle crée une nouvelle file binomiale contenant un tournoi binomial avec la clé à ajouter,
// puis elle effectue l'union de cette nouvelle file avec la file binomiale existante.
func (fb *FileBinomiale) Ajout(c cle.Cle) FileBinomiale {
	// Crée une nouvelle file binomiale contenant un tournoi binomial avec la clé à ajouter.
	toAdd := FileBinomiale{TournoiBinomial{Cle: &c}}

	// Effectue l'union de la file binomiale existante avec la nouvelle file contenant la clé ajoutée.
	return fb.Union(toAdd)
}

// SupprMin supprime le nœud avec la clé minimale de la file binomiale.
// Elle retourne une nouvelle file binomiale résultant de cette opération.
func (fb *FileBinomiale) SupprMin() FileBinomiale {
	// Assure que la file binomiale n'est pas vide avant de supprimer le minimum.
	assert(!fb.EstVide(), "SupprMin sur file vide")

	// Initialise le nœud smallest avec le premier élément de la file.
	smallest := (*fb)[0]
	smallestI := 0

	// Parcours les éléments de la file pour trouver le nœud avec la clé minimale.
	for i := 1; i < len(*fb); i++ {
		curr := &(*fb)[i]
		if curr.Cle.Inf(*smallest.Cle) {
			smallest = *curr
			smallestI = i
		}
	}

	*fb = slices.Delete(*fb, smallestI, smallestI+1)

	// Décapite le nœud smallest de la file.
	decapite := smallest.Decapite()

	// Effectue l'union de la file binomiale d'origine avec la file résultante après la décapitation.
	*fb = fb.Union(decapite)
	return *fb
}

// Construction crée une file binomiale à partir d'une liste de clés.
// Elle ajoute chaque clé à la file binomiale et renvoie la file résultante.
func Construction(cles []cle.Cle) FileBinomiale {
	// Initialise une file binomiale vide.
	fb := FileBinomiale{}

	// Ajoute chaque clé à la file binomiale.
	for i := 0; i < len(cles); i++ {
		fb = fb.Ajout(cles[i])
	}

	// Renvoie la file binomiale résultante.
	return fb
}

// uFret réalise l'union de deux files binomiales avec un nœud de tournoi binomial supplémentaire.
// Elle prend deux files binomiales (f1 et f2) ainsi qu'un tournoi binomial (t) en argument,
// et renvoie une nouvelle file binomiale résultant de l'union.
func uFret(f1, f2 FileBinomiale, t TournoiBinomial) FileBinomiale {
	// Cas de base : Si le tournoi est vide.
	if t.EstVide() {
		// Si l'une des deux files est vide, renvoie l'autre file.
		if f1.EstVide() {
			return f2
		}
		if f2.EstVide() {
			return f1
		}

		// Obtient les nœuds de degré minimum dans les deux files.
		t1 := f1.MinDeg()
		t2 := f2.MinDeg()

		// Compare les degrés et effectue l'union en conséquence.
		if t1.Degre < t2.Degre {
			union := f2.Union(f1.Reste())
			return union.AjoutMin(t1)
		}
		if t2.Degre < t1.Degre {
			union := f1.Union(f2.Reste())
			return union.AjoutMin(t2)
		}
		if t1.Degre == t2.Degre {
			return uFret(f1.Reste(), f2.Reste(), t1.Union(t2))
		}
	} else {
		// Cas : Le tournoi n'est pas vide.
		if f1.EstVide() {
			return f2.Union(t.File())
		}
		if f2.EstVide() {
			return f1.Union(t.File())
		}

		// Obtient les nœuds de degré minimum dans les deux files.
		t1 := f1.MinDeg()
		t2 := f2.MinDeg()

		// Compare les degrés et effectue l'union en conséquence.
		if t.Degre < t1.Degre && t.Degre < t2.Degre {
			union := f1.Union(f2)
			return union.AjoutMin(t)
		}
		if t.Degre == t1.Degre && t1.Degre == t2.Degre {
			tmp := uFret(f1.Reste(), f2.Reste(), t1.Union(t2))
			return tmp.AjoutMin(t)
		}
		if t.Degre == t1.Degre && t.Degre < t2.Degre {
			return uFret(f1.Reste(), f2, t1.Union(t))
		}
		if t.Degre == t2.Degre && t.Degre < t1.Degre {
			return uFret(f2.Reste(), f1, t2.Union(t))
		}
	}
	// Si aucun des cas ci-dessus n'est satisfait, déclenche une panique (erreur).
	panic("Fin de uFret, ne devrait pas être atteint")
}
