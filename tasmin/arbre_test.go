package tasmin

import (
	"projet/cle"
	"reflect"
	"testing"
)

func TestArbre_Ajout(t1 *testing.T) {
	// Créer des clés de test
	cle1 := cle.Cle{P1: 0, P2: 2}
	cle2 := cle.Cle{P1: 0, P2: 6}
	cle3 := cle.Cle{P1: 0, P2: 5}
	cle4 := cle.Cle{P1: 0, P2: 10}

	type args struct {
		c cle.Cle
	}

	tests := []struct {
		name  string
		arbre Arbre
		args  args
	}{
		{
			name: "Add key to an empty tree",
			arbre: Arbre{
				Size:      0,
				Height:    0,
				Cle:       cle.Cle{},
				Parent:    nil,
				EnfGauche: nil,
				EnfDroit:  nil,
			},
			args: args{c: cle1},
		},
		{
			name: "Add key to a 1 node tree",
			arbre: Arbre{
				Size:      1,
				Height:    1,
				Cle:       cle1,
				Parent:    nil,
				EnfGauche: nil,
				EnfDroit:  nil,
			},
			args: args{c: cle2},
		},
		{
			name: "Add key to create a left-skewed tree",
			arbre: Arbre{
				Size:      2,
				Height:    2,
				Cle:       cle1,
				Parent:    nil,
				EnfGauche: &Arbre{Cle: cle2, Size: 1, Height: 1, Parent: nil},
				EnfDroit:  nil,
			},
			args: args{c: cle3},
		},
		{
			name: "Add key to create a balanced tree",
			arbre: Arbre{
				Size:      3,
				Height:    2,
				Cle:       cle1,
				Parent:    nil,
				EnfGauche: &Arbre{Cle: cle2, Size: 1, Height: 1, Parent: nil},
				EnfDroit:  &Arbre{Cle: cle3, Size: 1, Height: 1, Parent: nil},
			},
			args: args{c: cle4},
		},
		{
			name: "Add key to create a tree + switches",
			arbre: Arbre{
				Size:      2,
				Height:    2,
				Cle:       cle2,
				Parent:    nil,
				EnfGauche: &Arbre{Cle: cle3, Size: 1, Height: 1, Parent: nil},
				EnfDroit:  nil,
			},
			args: args{c: cle1},
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:      tt.arbre.Size,
				Height:    tt.arbre.Height,
				Cle:       tt.arbre.Cle,
				Parent:    tt.arbre.Parent,
				EnfGauche: tt.arbre.EnfGauche,
				EnfDroit:  tt.arbre.EnfDroit,
			}
			t.Ajout(tt.args.c)

			// Assertions pour vérifier la taille, la hauteur de l'arbre et la valeur de l'ajout
			var expectedSize, expectedHeight int
			var expectedNode cle.Cle
			var expectedStringTree string

			switch tt.name {
			case "Add key to an empty tree":
				expectedSize = 1
				expectedHeight = 1
				expectedNode = t.Cle
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: nil,\n  right: nil\n}"
			case "Add key to a 1 node tree":
				expectedSize = 2
				expectedHeight = 2
				expectedNode = t.EnfGauche.Cle
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  },\n  right: nil\n}"
			case "Add key to create a left-skewed tree":
				expectedSize = 3
				expectedHeight = 2
				expectedNode = t.EnfDroit.Cle
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  }\n}"
			case "Add key to create a balanced tree":
				expectedSize = 4
				expectedHeight = 3
				expectedNode = t.EnfGauche.EnfGauche.Cle
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: Tas{\n      Cle: 10,\n      left: nil,\n      right: nil\n    },\n    right: nil\n  },\n  right: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  }\n}"
			case "Add key to create a tree + switches":
				expectedSize = 3
				expectedHeight = 2
				expectedNode = t.Cle
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  }\n}"
				// Ajouter d'autres cas de test au besoin.
			}

			// Assertion pour vérifier la taille de l'arbre
			if t.Size != expectedSize {
				t1.Errorf("Taille incorrecte de l'arbre. Got %d, want %d", t.Size, expectedSize)
			}

			// Assertion pour vérifier la hauteur de l'arbre
			if t.Height != expectedHeight {
				t1.Errorf("Hauteur incorrecte de l'arbre. Got %d, want %d", t.Height, expectedHeight)
			}

			// Assertion pour vérifier la structure de l'arbre
			if expectedNode != tt.args.c {
				t1.Errorf("Valeur de l'insertion incorrecte. Got %v, want %v", expectedNode, tt.args.c)
			}

			if expectedStringTree != t._string(0) {
				t1.Errorf("Affichage de l'arbre incorrecte. Got %v, want %v", t._string(0), expectedStringTree)
			}
			// Ajouter d'autres assertions en fonction de vos critères de succès.
		})
	}
}

func TestArbre_AjoutIteratif(t1 *testing.T) {
	// Créer des clés de test
	cle1 := cle.Cle{P1: 0, P2: 2}
	cle2 := cle.Cle{P1: 0, P2: 6}
	cle3 := cle.Cle{P1: 0, P2: 5}

	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name  string
		arbre Arbre
		args  args
	}{
		{
			name: "Add keys iteratively to an empty tree",
			arbre: Arbre{
				Size:      0,
				Height:    0,
				Cle:       cle.Cle{},
				Parent:    nil,
				EnfGauche: nil,
				EnfDroit:  nil,
			},
			args: args{cles: []cle.Cle{cle1, cle2, cle3}},
		},
		{
			name: "Add keys iteratively to a non-empty tree",
			arbre: Arbre{
				Size:      1,
				Height:    1,
				Cle:       cle2,
				Parent:    nil,
				EnfGauche: nil,
				EnfDroit:  nil,
			},
			args: args{cles: []cle.Cle{cle1, cle3}},
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:      tt.arbre.Size,
				Height:    tt.arbre.Height,
				Cle:       tt.arbre.Cle,
				Parent:    tt.arbre.Parent,
				EnfGauche: tt.arbre.EnfGauche,
				EnfDroit:  tt.arbre.EnfDroit,
			}
			t.AjoutIteratif(tt.args.cles)

			// Assertions pour vérifier la taille, la hauteur de l'arbre et la valeur de l'ajout itératif
			var expectedSize, expectedHeight int
			var expectedStringTree string

			switch tt.name {
			case "Add keys iteratively to an empty tree":
				expectedSize = len(tt.args.cles)
				expectedHeight = 2
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  }\n}"

			case "Add keys iteratively to a non-empty tree":
				expectedSize = len(tt.args.cles) + 1
				expectedHeight = 2
				expectedStringTree = "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  }\n}"
			}

			if t.Size != expectedSize {
				t1.Errorf("Taille incorrecte de l'arbre. Got %d, want %d", t.Size, expectedSize)
			}

			if t.Height != expectedHeight {
				t1.Errorf("Hauteur incorrecte de l'arbre. Got %d, want %d", t.Height, expectedHeight)
			}

			if t.Cle != cle1 {
				t1.Errorf("Valeur de racine. Got %d, want %d", t.Cle, cle1)
			}

			if t.EnfGauche.Cle != cle2 {
				t1.Errorf("Valeur de fils gauche. Got %d, want %d", t.EnfGauche.Cle, cle2)
			}

			if t.EnfDroit.Cle != cle3 {
				t1.Errorf("Valeur de fils droit. Got %d, want %d", t.EnfDroit.Cle, cle3)
			}

			if expectedStringTree != t._string(0) {
				t1.Errorf("Affichage de l'arbre incorrecte. Got %v, want %v", t._string(0), expectedStringTree)
			}

		})
	}
}

func TestArbre_Construction(t *testing.T) {
	clesFull := []cle.Cle{
		{P1: 0, P2: 6},
		{P1: 0, P2: 2},
		{P1: 0, P2: 5},
	}
	clesNonFull := []cle.Cle{
		{P1: 0, P2: 2},
		{P1: 0, P2: 6},
		{P1: 0, P2: 5},
		{P1: 0, P2: 4},
	}

	tests := []struct {
		name string
		args []cle.Cle
		want Arbre
	}{
		{
			name: "Create a full tree",
			args: clesFull,
			want: Arbre{
				Size:   3,
				Height: 2,
				Cle:    cle.Cle{P1: 0, P2: 2},
				EnfGauche: &Arbre{
					Size:   1,
					Height: 1,
					Cle:    cle.Cle{P1: 0, P2: 6},
				},
				EnfDroit: &Arbre{
					Size:   1,
					Height: 1,
					Cle:    cle.Cle{P1: 0, P2: 5},
				},
			},
		},
		{
			name: "Create a non full tree",
			args: clesNonFull,
			want: Arbre{
				Size:   4,
				Height: 3,
				Cle:    cle.Cle{P1: 0, P2: 2},
				EnfGauche: &Arbre{
					Size:   2,
					Height: 2,
					Cle:    cle.Cle{P1: 0, P2: 4},
					EnfGauche: &Arbre{
						Size:   1,
						Height: 1,
						Cle:    cle.Cle{P1: 0, P2: 6},
					},
				},
				EnfDroit: &Arbre{
					Size:   1,
					Height: 1,
					Cle:    cle.Cle{P1: 0, P2: 5},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {
			arbre := ConstructionArbre(test.args)
			if arbre.Cle != test.want.Cle ||
				arbre.Size != test.want.Size ||
				arbre.Height != test.want.Height ||
				arbre.EnfGauche.Cle != test.want.EnfGauche.Cle ||
				arbre.EnfGauche.Size != test.want.EnfGauche.Size ||
				arbre.EnfGauche.Height != test.want.EnfGauche.Height ||
				arbre.EnfDroit.Cle != test.want.EnfDroit.Cle ||
				arbre.EnfDroit.Size != test.want.EnfDroit.Size ||
				arbre.EnfDroit.Height != test.want.EnfDroit.Height {
				t1.Errorf("Constuction() = %v, want %v", arbre.String(), test.want.String())
			}
		})
	}
}

func TestArbre_SupprMin(t *testing.T) {
	arbreFull := Arbre{
		Size:   3,
		Height: 2,
		Cle:    cle.Cle{P1: 0, P2: 2},
		EnfGauche: &Arbre{
			Size:   1,
			Height: 1,
			Cle:    cle.Cle{P1: 0, P2: 6},
		},
		EnfDroit: &Arbre{
			Size:   1,
			Height: 1,
			Cle:    cle.Cle{P1: 0, P2: 5},
		},
	}
	arbreFull.EnfGauche.Parent = &arbreFull
	arbreFull.EnfDroit.Parent = &arbreFull

	tests := []struct {
		name  string
		arbre Arbre
		state Arbre
		want  cle.Cle
	}{
		{
			name:  "Extract Min from a full tree",
			arbre: arbreFull,
			state: Arbre{
				Size:   2,
				Height: 2,
				Cle:    cle.Cle{P1: 0, P2: 5},
				EnfGauche: &Arbre{
					Size:   1,
					Height: 1,
					Cle:    cle.Cle{P1: 0, P2: 6},
				},
			},
			want: cle.Cle{P1: 0, P2: 2},
		},
	}
	t.Run(tests[0].name, func(t *testing.T) {
		arbre := tests[0].arbre
		deleted := arbre.SupprMin()
		if deleted != tests[0].want {
			t.Errorf("SupprMin() = %v, got %v", deleted, tests[0].want)
		}
	})
}

func TestArbre_String(t1 *testing.T) {
	type fields struct {
		Size       int
		height     int
		cle        cle.Cle
		parent     *Arbre
		leftChild  *Arbre
		rightChild *Arbre
	}
	cle1 := cle.Cle{P1: 0, P2: 2}
	cle2 := cle.Cle{P1: 0, P2: 6}
	cle3 := cle.Cle{P1: 0, P2: 5}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String representation of an empty tree",
			fields: fields{
				Size:       0,
				height:     0,
				cle:        cle.Cle{},
				parent:     nil,
				leftChild:  nil,
				rightChild: nil,
			},
			want: "Tas{\n  Cle: 0,\n  left: nil,\n  right: nil\n}\n",
		},
		{
			name: "String representation of a non empty tree",
			fields: fields{
				Size:       3,
				height:     2,
				cle:        cle1,
				parent:     nil,
				leftChild:  &Arbre{Cle: cle2, Size: 1, Height: 1, Parent: nil},
				rightChild: &Arbre{Cle: cle3, Size: 1, Height: 1, Parent: nil},
			},
			want: "Tas{\n  Cle: 2,\n  left: Tas{\n    Cle: 6,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 5,\n    left: nil,\n    right: nil\n  }\n}\n",
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:      tt.fields.Size,
				Height:    tt.fields.height,
				Cle:       tt.fields.cle,
				Parent:    tt.fields.parent,
				EnfGauche: tt.fields.leftChild,
				EnfDroit:  tt.fields.rightChild,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArbre__string(t1 *testing.T) {
	// Créer des clés de test
	cle1 := cle.Cle{P1: 0, P2: 12}
	cle2 := cle.Cle{P1: 0, P2: 16}
	cle3 := cle.Cle{P1: 0, P2: 15}

	type fields struct {
		Size       int
		height     int
		cle        cle.Cle
		parent     *Arbre
		leftChild  *Arbre
		rightChild *Arbre
	}
	type args struct {
		indent int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test _string with both children",
			fields: fields{
				Size:       2,
				height:     2,
				cle:        cle1,
				parent:     nil,
				leftChild:  &Arbre{Cle: cle2, Size: 1, Height: 1, Parent: nil},
				rightChild: &Arbre{Cle: cle3, Size: 1, Height: 1, Parent: nil},
			},
			args: args{indent: 0},
			want: "Tas{\n  Cle: 12,\n  left: Tas{\n    Cle: 16,\n    left: nil,\n    right: nil\n  },\n  right: Tas{\n    Cle: 15,\n    left: nil,\n    right: nil\n  }\n}",
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:      tt.fields.Size,
				Height:    tt.fields.height,
				Cle:       tt.fields.cle,
				Parent:    tt.fields.parent,
				EnfGauche: tt.fields.leftChild,
				EnfDroit:  tt.fields.rightChild,
			}
			if got := t._string(tt.args.indent); got != tt.want {
				t1.Errorf("_string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArbre(t *testing.T) {
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name string
		args args
		want Arbre
	}{
		{
			name: "Create tree with key 20",
			args: args{c: cle.Cle{P1: 0, P2: 20}},
			want: Arbre{Cle: cle.Cle{P1: 0, P2: 20}, Height: 1, Size: 1},
		},
		{
			name: "Create tree with key 5",
			args: args{c: cle.Cle{P1: 0, P2: 5}},
			want: Arbre{Cle: cle.Cle{P1: 0, P2: 5}, Height: 1, Size: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArbre(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArbre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Exponentiation of 2^3",
			args: args{x: 2, y: 3},
			want: 8,
		},
		{
			name: "Exponentiation of 5^2",
			args: args{x: 5, y: 2},
			want: 25,
		},
		{
			name: "Exponentiation of 0^3",
			args: args{x: 0, y: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
