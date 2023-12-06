package tasmin

import (
	"errors"
	"projet/cle"
	"strings"
)

type Tableau []cle.Cle

func (*Tableau) Parent(i int) int {
	return (i - 1) / 2
}

// Parent retourne l'indice du parent de l'élément situé à l'indice donné dans le tableau.
func (*Tableau) EnfGauche(i int) int {
	return 2*i + 1
}

// EnfDroit retourne l'indice du fils droit de l'élément situé à l'indice donné dans le tableau.
func (*Tableau) EnfDroit(i int) int {
	return 2*i + 2
}

// SupprMin supprime et retourne la clé minimale du tableau, qui doit être un tas.
func (t *Tableau) SupprMin() (cle.Cle, error) {
	// Vérifier si le tableau est vide.
	if len(*t) == 0 {
		// Si le tableau est vide, retourner une clé nulle et une erreur.
		return cle.Cle{}, errors.New("SupprMin sur tas vide")
	}

	// Sauvegarder la clé à supprimer (la clé minimale dans un tas).
	cleSupprime := (*t)[0]

	// Remplacer la clé à l'indice 0 par la dernière clé dans le tableau.
	(*t)[0] = (*t)[len(*t)-1]
	// Supprimer le dernier élément
	*t = (*t)[1:]

	// Appeler la méthode de tri pour maintenir la propriété du tas.
	t.trier(0)

	// Retourner la clé supprimée avec succès et aucune erreur.
	return cleSupprime, nil
}

// Ajout ajoute une nouvelle clé au tableau, qui doit être un tas, et réorganise le tableau
// pour maintenir la propriété du tas.
func (t *Tableau) Ajout(c cle.Cle) {
	// Ajouter la nouvelle clé à la fin du tableau.
	*t = append(*t, c)

	// Récupérer l'indice actuel de la nouvelle clé.
	currI := len(*t) - 1

	// Calculer l'indice du parent dans le tableau.
	parentI := t.Parent(currI)

	// Vérifier si la nouvelle clé est inférieure à son parent.
	currInfParent := (*t)[currI].Inf((*t)[parentI])

	// Réorganiser le tableau tant que la propriété du tas est violée.
	for currI > 0 && currInfParent {
		// Échanger la nouvelle clé avec son parent.
		(*t)[parentI], (*t)[currI] = (*t)[currI], (*t)[parentI]

		// Mettre à jour les indices pour le prochain tour de boucle.
		currI = parentI
		parentI = t.Parent(currI)
		currInfParent = (*t)[currI].Inf((*t)[parentI])
	}
}

// AjoutIteratif ajoute itérativement toutes les clés du tableau donné au tableau actuel,
// qui doit être un tas. Cette méthode utilise la méthode Ajout pour chaque clé ajoutée.
func (t *Tableau) AjoutIteratif(cles []cle.Cle) {
	// Parcourir toutes les clés du tableau donné.
	for i := 0; i < len(cles); i++ {
		// Appeler la méthode Ajout pour ajouter la clé au tableau actuel.
		t.Ajout(cles[i])
	}
}

// Construction construit un tas binaire à partir d'un tableau de clés donné.
// Elle renvoie un nouveau tableau qui représente un tas binaire.
func Construction(cles []cle.Cle) Tableau {
	// Créer un nouveau tableau initialisé avec les clés fournies.
	var t Tableau = cles

	// Trouver l'indice du dernier nœud ayant des enfants dans le tableau.
	// Cela permet de commencer la construction du tas à partir du bas.
	start := (len(cles) / 2) - 1

	// Itérer à partir du dernier nœud avec des enfants jusqu'au premier nœud du tableau.
	for i := start; i >= 0; i-- {
		// Appeler la méthode de tri pour réorganiser le tableau et maintenir la propriété du tas.
		t.trier(i)
	}

	// Retourner le nouveau tableau qui représente un tas binaire.
	return t
}

// Union combine le tableau actuel avec un autre tableau donné, puis reconstruit
// le tableau résultant pour maintenir la propriété du tas.
// La méthode modifie le tableau actuel et renvoie le tableau résultant.
func (t *Tableau) Union(o Tableau) Tableau {
	// Ajouter tous les éléments de l'autre tableau à la fin du tableau actuel.
	*t = append(*t, o...)

	// Reconstruire le tableau pour maintenir la propriété du tas.
	return Construction(*t)
}

// trier réorganise le tableau de manière récursive pour maintenir la propriété du tas
// à partir de l'indice donné. Elle compare l'élément actuel avec ses enfants
// et effectue des échanges si nécessaire pour maintenir la propriété du tas.
func (t *Tableau) trier(index int) {
	// Vérifier si le tableau est vide ou contient un seul élément.
	if len(*t) <= 1 {
		return
	}

	// Calculer les indices des fils gauche et droit de l'élément à l'indice donné.
	gauche := t.EnfGauche(index)
	droite := t.EnfDroit(index)

	// Trouver l'indice du minimum parmi l'élément actuel, le fils gauche et le fils droit.
	minimum := index
	if gauche < len(*t) && (*t)[gauche].Inf((*t)[index]) {
		minimum = gauche
	}

	if droite < len(*t) && (*t)[droite].Inf((*t)[minimum]) {
		minimum = droite
	}

	// Si le minimum n'est pas l'indice actuel, échanger les éléments et continuer la récursion.
	if minimum != index {
		(*t)[index], (*t)[minimum] = (*t)[minimum], (*t)[index]

		// Appeler récursivement la méthode trier sur le nouveau indice du minimum.
		t.trier(minimum)
	}
}

// _string génère une représentation en chaîne de caractères récursive du tas
// à partir de l'indice donné. Cette méthode est utilisée à des fins de débogage.
func (t *Tableau) _string(i int, indent int) string {
	// Initialise la chaîne avec le début de la représentation du tas.
	str := "Tas{\n"
	indent += 1
	// str += strings.Repeat("│ ", indent) + "cle: " + (*t)[i].DecimalString() + ",\n"

	// Ajoute la représentation en chaîne de l'élément actuel avec son indentation.
	str += strings.Repeat("  ", indent) + "cle: " + (*t)[i].DecimalString() + ",\n"

	// Calcul des indices des fils gauche et droit de l'élément actuel.
	leftI := t.EnfGauche(i)
	rightI := t.EnfDroit(i)

	// Ajoute la représentation en chaîne du fils gauche s'il existe, sinon "nil".
	if leftI < len(*t) {
		str += strings.Repeat("  ", indent) + "left: " + t._string(leftI, indent) + ",\n"
	} else {
		str += strings.Repeat("  ", indent) + "left: nil,\n"
	}

	// Ajoute la représentation en chaîne du fils droit s'il existe, sinon "nil".
	if rightI < len(*t) {
		str += strings.Repeat("  ", indent) + "right: " + t._string(rightI, indent) + "\n"
	} else {
		str += strings.Repeat("  ", indent) + "right: nil\n"
	}

	// Ajoute la dernière ligne de fermeture du tas avec l'indentation appropriée.
	str += strings.Repeat("  ", indent-1) + "}"

	return str
}

// String génère une représentation en chaîne de caractères du tas.
func (t *Tableau) String() string {
	// Appelle la méthode _string pour générer la représentation en chaîne à partir de la racine.
	return t._string(0, 0) + "\n"
}
