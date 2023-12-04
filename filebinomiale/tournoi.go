package filebinomiale

import "projet/cle"

// TournoiBinomial représente un nœud dans un arbre de tournois binomiaux.
type TournoiBinomial struct {
	// Cle est la clé associée au nœud.
	Cle *cle.Cle

	// Parent est un pointeur vers le parent du nœud dans l'arbre de tournois.
	Parent *TournoiBinomial

	// Enfants est une file de tournois binomiaux représentant les fils du nœud.
	Enfants FileBinomiale

	// Degre est le degré du nœud, c'est-à-dire le nombre d'enfants directs.
	Degre int
}

// EstVide retourne true si le nœud du tournoi est vide, c'est-à-dire s'il ne contient pas de clé.
// Sinon, elle retourne false.
func (tb TournoiBinomial) EstVide() bool {
	// Vérifie si la clé du nœud est nulle.
	return tb.Cle == nil
}

// Union réalise l'opération d'union entre deux nœuds de tournoi binomial.
// Elle prend un autre nœud de tournoi binomial en argument, effectue l'union des deux nœuds
// en fonction de leurs degrés, et renvoie le nœud résultant.
func (tb TournoiBinomial) Union(o TournoiBinomial) TournoiBinomial {
	// Assure que les deux nœuds ont le même degré avant l'union.
	assert(tb.Degre == o.Degre, "degre inégal")

	// Ajoute le nœud en argument à la liste des enfants du nœud actuel.
	tb.Enfants = append(tb.Enfants, o)

	// Retourne le nœud résultant après l'union.
	return tb
}

// Decapite réalise l'opération de décapitation sur le nœud du tournoi binomial.
// Elle renvoie la liste des enfants du nœud décapité.
func (tb TournoiBinomial) Decapite() FileBinomiale {
	// Assure que le nœud du tournoi n'est pas vide avant la décapitation.
	assert(!tb.EstVide(), "tb vide")

	// Retourne la liste des enfants du nœud décapité.
	return tb.Enfants
}

// File crée une file binomiale à partir du nœud du tournoi binomial actuel.
// Elle renvoie une nouvelle file binomiale contenant le nœud actuel.
func (tb TournoiBinomial) File() FileBinomiale {
	// Assure que le nœud du tournoi n'est pas vide avant de créer la file.
	assert(!tb.EstVide(), "tb vide")

	// Retourne une nouvelle file binomiale contenant le nœud actuel.
	return FileBinomiale{tb}
}

// assert vérifie la condition spécifiée. Si la condition est fausse, la fonction déclenche une panique
// avec le message d'erreur spécifié.
func assert(condition bool, msg string) {
	// Vérifie si la condition est fausse.
	if !condition {
		// Déclenche une panique avec le message d'erreur spécifié.
		panic(msg)
	}
}
