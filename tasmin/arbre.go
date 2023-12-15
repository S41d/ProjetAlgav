package tasmin

import (
	"math"
	"projet/cle"
	"strings"
)

// Arbre est une structure de données représentant un nœud dans un arbre.
type Arbre struct {
	// Size indique le nombre total de nœuds dans le sous-arbre enraciné à ce nœud.
	Size int
	// Height représente la hauteur du sous-arbre enraciné à ce nœud.
	Height int
	// Cle est la valeur de la clé stockée dans ce nœud.
	Cle cle.Cle
	// Parent pointe vers le nœud Parent dans l'arbre.
	Parent *Arbre
	// LeftChild pointe vers le nœud enfant gauche dans l'arbre.
	EnfGauche *Arbre
	// RightChild pointe vers le nœud enfant droit dans l'arbre.
	EnfDroit *Arbre
}

// Ajout ajoute un élément à l'arbre.
func (t *Arbre) Ajout(c cle.Cle) {
	if t.Size == 0 {
		*t = NewArbre(c)
		return
	}
	// Initialisation du pointeur courant à la racine de l'arbre
	curr := t

	// Vérifier si l'arbre est plein (2^h - 1 nœuds), où h est la hauteur actuelle de l'arbre
	if curr.Size == pow(2, curr.Height)-1 {
		// Si l'arbre est plein, descendre à gauche jusqu'au dernier niveau
		for curr.EnfGauche != nil {
			curr.Size += 1
			curr.Height += 1
			curr = curr.EnfGauche
		}

		// Mettre à jour la taille et la hauteur du nœud actuel
		curr.Size += 1
		curr.Height += 1

		// Ajouter un nouveau nœud à gauche avec la clé donnée
		curr.EnfGauche = &Arbre{Cle: c, Size: 1, Height: 1, Parent: curr}
		curr = curr.EnfGauche
	} else {
		// Si l'arbre n'est pas plein, parcourir l'arbre pour trouver l'emplacement d'insertion
		for {
			curr.Size += 1

			// Si le nœud courant n'a pas d'enfant gauche, ajouter le nouveau nœud à gauche
			if curr.EnfGauche == nil {
				curr.Height += 1
				curr.EnfGauche = &Arbre{Cle: c, Size: 1, Height: 1, Parent: curr}

				// Mettre à jour la hauteur des parents jusqu'à la racine
				currHeight := curr
				for currHeight.Parent != nil {
					if currHeight.Parent.Height == currHeight.Height+1 {
						break
					}
					currHeight.Parent.Height += 1
					currHeight = currHeight.Parent
				}

				// Aller au nouveau nœud ajouté
				curr = curr.EnfGauche
				break
			} else if curr.EnfDroit == nil {
				// Si le nœud courant n'a pas d'enfant droit, ajouter le nouveau nœud à droite
				curr.EnfDroit = &Arbre{Cle: c, Size: 1, Height: 1, Parent: curr}
				curr = curr.EnfDroit
				break
			}

			// Choisir le chemin en fonction de la taille des sous-arbres
			if curr.EnfGauche.Size <= curr.EnfDroit.Size || curr.EnfGauche.Size != pow(2, curr.EnfGauche.Height)-1 {
				curr = curr.EnfGauche
			} else {
				curr = curr.EnfDroit
			}
		} // for
	} // else

	// Rétablir la propriété de l'arbre binaire équilibré en effectuant des échanges
	for curr.Parent != nil {
		if !curr.Parent.Cle.Inf(curr.Cle) {
			// Échanger les clés si la clé du Parent est supérieure à la clé actuelle
			temp := *curr
			curr.Cle = curr.Parent.Cle
			curr.Parent.Cle = temp.Cle
			curr = curr.Parent
		} else {
			break
		}
	} // for
}

func ConstructionArbre(cles []cle.Cle) Arbre {
	if len(cles) == 1 {
		return NewArbre(cles[0])
	}
	hauteur := int(math.Floor(math.Log2(float64(len(cles))))) + 1
	numNodes := 0
	if len(cles) == pow(2, hauteur)-1 {
		numNodes = pow(2, hauteur-1)
	} else {
		numNodes = len(cles) - (pow(2, hauteur-1) - 1)
	}
	bound := len(cles) - numNodes
	nodesClesEnf := cles[bound:]
	nodesEnf := listCleToListArbre(nodesClesEnf)

	for i := hauteur - 1; i > 0; i-- {
		cles = cles[:len(cles)-numNodes]
		numNodes = pow(2, i-1)
		nodesCles := cles[len(cles)-numNodes:]
		nodes := listCleToListArbre(nodesCles)
		for j := 0; j < len(nodes); j++ {
			k := j * 2
			if k >= len(nodesEnf) {
				break
			}
			if k != len(nodesEnf)-1 {
				nodes[j].EnfGauche = &nodesEnf[k]
				nodes[j].EnfDroit = &nodesEnf[k+1]

				nodes[j].EnfGauche.Parent = &nodes[j]
				nodes[j].EnfDroit.Parent = &nodes[j]

				nodes[j].Size = 1 + nodes[j].EnfGauche.Size + nodes[j].EnfDroit.Size
				nodes[j].Height = 1 + max(nodes[j].EnfGauche.Height, nodes[j].EnfDroit.Height)
			} else {
				nodes[j].EnfGauche = &nodesEnf[k]
				nodes[j].EnfGauche.Parent = &nodes[j]

				nodes[j].Size = 1 + nodesEnf[k].Size
				nodes[j].Height = 1 + nodesEnf[k].Height
			}
			nodes[j].reOrganiser()
		}
		nodesEnf = nodes
		nodesClesEnf = nodesCles
	}
	return nodesEnf[0]
}

func (t *Arbre) SupprMin() cle.Cle {
	if t.Size == 1 {
		return t.Cle
	}
	deleted := t.Cle

	adjustHeight := func(curr *Arbre) {
		for curr.Parent != nil {
			p := curr.Parent
			if p.EnfGauche == curr {
				p.Height -= 1
				curr = p
			} else {
				break
			}
		}
	}

	curr := t
	for curr.Size != 1 {
		if curr.EnfDroit != nil {
			if curr.EnfGauche.Height == curr.EnfDroit.Height {
				curr = curr.EnfDroit
			} else {
				// hauteur de EnfGauche >= hauteur de EnfDroit
				curr = curr.EnfGauche
			}
		} else if curr.EnfGauche != nil {
			curr = curr.EnfGauche
		}
	}
	m := curr.Cle
	if curr.Parent.EnfGauche == curr {
		toDelete := curr
		adjustHeight(curr)
		// toDelete.Parent.Size--
		toDelete.Parent.EnfGauche = nil
	} else {
		curr.Parent.EnfDroit = nil
	}
	for curr.Parent != nil {
		curr.Parent.Size--
		curr = curr.Parent
	}
	curr.Cle = m
	*t = *curr
	t.reOrganiser()
	return deleted
}

func (t *Arbre) Union(o *Arbre) Arbre {
	cles := append(ClesOfArbre(t), ClesOfArbre(o)...)
	return ConstructionArbre(cles)
}

func ClesOfArbre(a *Arbre) []cle.Cle {
	cles := make([]cle.Cle, a.Size)

	var aux func(curr *Arbre, idx int)
	aux = func(curr *Arbre, idx int) {
		cles[idx] = curr.Cle
		if curr.EnfGauche != nil {
			aux(curr.EnfGauche, (2*idx)+1)
		}
		if curr.EnfDroit != nil {
			aux(curr.EnfDroit, (2*idx)+2)
		}
	}
	aux(a, 0)
	return cles
}

func (t *Arbre) reOrganiser() {
	smallest := &Arbre{}
	if t.EnfGauche != nil && t.EnfDroit != nil {
		if t.EnfGauche.Cle.Inf(t.EnfDroit.Cle) {
			smallest = t.EnfGauche
		} else {
			smallest = t.EnfDroit
		}
	} else if t.EnfGauche != nil {
		smallest = t.EnfGauche
	} else {
		return
	}
	if smallest.Cle.Inf(t.Cle) {
		t.Cle, smallest.Cle = smallest.Cle, t.Cle
		smallest.reOrganiser()
	}
}

// AjoutIteratif ajoute de manière itérative une liste de clés à l'arbre.
func (t *Arbre) AjoutIteratif(cles []cle.Cle) {
	// Parcourir toutes les clés dans la liste
	for index := 0; index < len(cles); index++ {
		// Appeler la fonction Ajout pour ajouter la clé actuelle à l'arbre
		t.Ajout(cles[index])
	}
}

// _string génère une représentation en chaîne de caractères de l'arbre, avec un décalage pour l'indentation.
func (t *Arbre) _string(indent int) string {
	// Initialiser la chaîne avec l'en-tête de l'arbre
	str := "Tas{\n"
	indent += 1 // Augmenter l'indentation pour les sous-arbres

	// Ajouter la représentation en chaîne de la clé du nœud courant
	str += strings.Repeat("  ", indent) + "Cle: " + t.Cle.DecimalString() + ",\n"

	// Ajouter d'autres propriétés du nœud (commentées pour le moment)
	// str += strings.Repeat("  ", indent) + "size: " + fmt.Sprint(t.size) + ",\n"
	// str += strings.Repeat("  ", indent) + "Height: " + fmt.Sprint(t.Height) + ",\n"

	// Ajouter la représentation en chaîne du sous-arbre gauche
	if t.EnfGauche != nil {
		str += strings.Repeat("  ", indent) + "left: " + t.EnfGauche._string(indent) + ",\n"
	} else {
		str += strings.Repeat("  ", indent) + "left: nil,\n"
	}

	// Ajouter la représentation en chaîne du sous-arbre droit
	if t.EnfDroit != nil {
		str += strings.Repeat("  ", indent) + "right: " + t.EnfDroit._string(indent) + "\n"
	} else {
		str += strings.Repeat("  ", indent) + "right: nil\n"
	}

	// Ajouter la fermeture de l'arbre avec le bon décalage
	str += strings.Repeat("  ", indent-1) + "}"

	return str
}

// NewArbre crée un nouvel arbre avec la clé spécifiée, initialisant la hauteur à 1 et la taille à 1.
func NewArbre(c cle.Cle) Arbre {
	return Arbre{Cle: c, Height: 1, Size: 1}
}

// pow retourne x élevé à la puissance y.
func pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// String retourne une représentation en chaîne de caractères de l'arbre.
func (t *Arbre) String() string {
	return t._string(0) + "\n"
}

func listCleToListArbre(cles []cle.Cle) []Arbre {
	var arbres []Arbre
	for i := 0; i < len(cles); i++ {
		arbres = append(arbres, NewArbre(cles[i]))
	}
	return arbres
}
