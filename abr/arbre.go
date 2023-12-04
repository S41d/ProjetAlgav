package arbrerecherche

import "projet/cle"

// ArbreRecherche représente un nœud dans un arbre de recherche.
type ArbreRecherche struct {
	Elt      *cle.Cle        // Elt est l'élément stocké dans le nœud de l'arbre.
	Parent   *ArbreRecherche // Parent est un pointeur vers le nœud parent dans l'arbre.
	SaGauche *ArbreRecherche // SaGauche est un pointeur vers le sous-arbre gauche.
	SaDroit  *ArbreRecherche // SaDroit est un pointeur vers le sous-arbre droit.
}

// Contient vérifie si la clé spécifiée est présente dans l'arbre de recherche.
// La méthode utilise la comparaison définie par les méthodes Eg() et Inf() de la clé.
// Si la clé est égale à la clé du nœud actuel, la fonction retourne true.
// Si la clé est inférieure à la clé du nœud actuel, la recherche continue dans le sous-arbre gauche.
// Sinon, la recherche continue dans le sous-arbre droit.
func (a *ArbreRecherche) Contient(c cle.Cle) bool {
	if c.Eg(*a.Elt) {
		return true
	} else if c.Inf(*a.Elt) {
		return a.SaGauche.Contient(c)
	} else {
		return a.SaDroit.Contient(c)
	}
}

// Ajout insère une nouvelle clé dans l'arbre de recherche.
func (a *ArbreRecherche) Ajout(c cle.Cle) {
	// Le nœud actuel est vide, donc la clé est ajoutée à ce nœud.
	if a.Elt == nil {
		a.Elt = &c
	} else if c.Eg(*a.Elt) {
		// La clé est égale à la clé du nœud actuel, l'insertion est évitée (pas de doublon).
		return
	} else if c.Inf(*a.Elt) {
		// La clé est inférieure à la clé du nœud actuel, l'insertion continue dans le sous-arbre gauche.
		a.SaGauche.Ajout(c)
	} else {
		// La clé est supérieure à la clé du nœud actuel, l'insertion continue dans le sous-arbre droit.
		a.SaDroit.Ajout(c)
	}
}

// Suppr supprime la clé spécifiée de l'arbre de recherche.
// Si le nœud actuel est vide, la fonction ne fait rien.
// Si la clé est égale à la clé du nœud actuel, la clé du nœud actuel est supprimée.
// Si la clé est inférieure à la clé du nœud actuel, la suppression continue dans le sous-arbre gauche.
// Sinon, la suppression continue dans le sous-arbre droit.
func (a *ArbreRecherche) Suppr(c cle.Cle) {
	if a.Elt == nil {
		// Le nœud actuel est vide, la fonction ne fait rien.
		return
	} else if c.Eg(*a.Elt) {
		// La clé est égale à la clé du nœud actuel, la clé du nœud actuel est supprimée.
		a.Elt = nil
	} else if c.Inf(*a.Elt) {
		// La clé est inférieure à la clé du nœud actuel, la suppression continue dans le sous-arbre gauche.
		a.SaGauche.Suppr(c)
	} else {
		// La clé est supérieure à la clé du nœud actuel, la suppression continue dans le sous-arbre droit.
		a.SaDroit.Suppr(c)
	}
}

// EstArbreVide retourne true si l'arbre de recherche est vide, c'est-à-dire s'il n'a pas de racine.
// Sinon, retourne false.
func (a *ArbreRecherche) EstArbreVide() bool {
	// Si le nœud actuel (la racine de l'arbre) est vide, alors l'arbre est vide.
	return a.Elt == nil
}

// Racine retourne la racine de l'arbre de recherche en remontant jusqu'à la racine.
func (a *ArbreRecherche) Racine() ArbreRecherche {
	// Initialise le nœud courant à la racine de l'arbre.
	curr := a

	// Continue de remonter jusqu'à ce que le nœud courant n'ait plus de parent.
	for curr.Parent != nil {
		curr = curr.Parent
	}

	// Retourne la racine de l'arbre.
	return *curr
}

// NewArbreRecherche crée et retourne un nouvel arbre de recherche à partir d'une liste de clés.
// Les clés sont ajoutées à l'arbre dans l'ordre spécifié.
func NewArbreRecherche(cles []cle.Cle) ArbreRecherche {
	// Crée un nouvel arbre de recherche.
	a := ArbreRecherche{}

	// Ajoute chaque clé à l'arbre dans l'ordre spécifié.
	for i := 0; i < len(cles); i++ {
		a.Ajout(cles[i])
	}

	// Retourne l'arbre nouvellement créé.
	return a
}
