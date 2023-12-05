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
				Size:       0,
				height:     0,
				cle:        cle.Cle{},
				parent:     nil,
				leftChild:  nil,
				rightChild: nil,
			},
			args: args{c: cle1},
		},
		{
			name: "Add key to a 1 node tree",
			arbre: Arbre{
				Size:       1,
				height:     1,
				cle:        cle1,
				parent:     nil,
				leftChild:  nil,
				rightChild: nil,
			},
			args: args{c: cle2},
		},
		{
			name: "Add key to create a left-skewed tree",
			arbre: Arbre{
				Size:       2,
				height:     2,
				cle:        cle1,
				parent:     nil,
				leftChild:  &Arbre{cle: cle2, Size: 1, height: 1, parent: nil},
				rightChild: nil,
			},
			args: args{c: cle3},
		},
		{
			name: "Add key to create a balanced tree",
			arbre: Arbre{
				Size:       3,
				height:     2,
				cle:        cle1,
				parent:     nil,
				leftChild:  &Arbre{cle: cle2, Size: 1, height: 1, parent: nil},
				rightChild: &Arbre{cle: cle3, Size: 1, height: 1, parent: nil},
			},
			args: args{c: cle4},
		},
		{
			name: "Add key to create a tree + switches",
			arbre: Arbre{
				Size:       2,
				height:     2,
				cle:        cle2,
				parent:     nil,
				leftChild:  &Arbre{cle: cle3, Size: 1, height: 1, parent: nil},
				rightChild: nil,
			},
			args: args{c: cle1},
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:       tt.arbre.Size,
				height:     tt.arbre.height,
				cle:        tt.arbre.cle,
				parent:     tt.arbre.parent,
				leftChild:  tt.arbre.leftChild,
				rightChild: tt.arbre.rightChild,
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
				expectedNode = t.cle
				expectedStringTree = "Tas{\n          cle: 2,\n          left: nil,\n          right: nil\n        }"
			case "Add key to a 1 node tree":
				expectedSize = 2
				expectedHeight = 2
				expectedNode = t.leftChild.cle
				expectedStringTree = "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: nil,\n            right: nil\n          },\n          right: nil\n        }"
			case "Add key to create a left-skewed tree":
				expectedSize = 3
				expectedHeight = 2
				expectedNode = t.rightChild.cle
				expectedStringTree = "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: nil,\n            right: nil\n          },\n          right: Tas{\n            cle: 5,\n            left: nil,\n            right: nil\n          }"
			case "Add key to create a balanced tree":
				expectedSize = 4
				expectedHeight = 3
				expectedNode = t.leftChild.leftChild.cle
				expectedStringTree = "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: Tas{\n              cle: 10,\n              left: nil,\n              right: nil\n            },\n            right: nil\n          },\n          right: Tas{\n            cle: 5,\n            left: nil,\n            right: nil\n          }"
			case "Add key to create a tree + switches":
				expectedSize = 3
				expectedHeight = 2
				expectedNode = t.cle
				expectedStringTree = "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 5,\n            left: nil,\n            right: nil\n          },\n          right: Tas{\n            cle: 6,\n            left: nil,\n            right: nil\n          }"
				// Ajouter d'autres cas de test au besoin.
			}

			// Assertion pour vérifier la taille de l'arbre
			if t.Size != expectedSize {
				t1.Errorf("Taille incorrecte de l'arbre. Got %d, want %d", t.Size, expectedSize)
			}

			// Assertion pour vérifier la hauteur de l'arbre
			if t.height != expectedHeight {
				t1.Errorf("Hauteur incorrecte de l'arbre. Got %d, want %d", t.height, expectedHeight)
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
				Size:       0,
				height:     0,
				cle:        cle.Cle{},
				parent:     nil,
				leftChild:  nil,
				rightChild: nil,
			},
			args: args{cles: []cle.Cle{cle1, cle2, cle3}},
		},
		{
			name: "Add keys iteratively to a non-empty tree",
			arbre: Arbre{
				Size:       1,
				height:     1,
				cle:        cle2,
				parent:     nil,
				leftChild:  nil,
				rightChild: nil,
			},
			args: args{cles: []cle.Cle{cle1, cle3}},
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:       tt.arbre.Size,
				height:     tt.arbre.height,
				cle:        tt.arbre.cle,
				parent:     tt.arbre.parent,
				leftChild:  tt.arbre.leftChild,
				rightChild: tt.arbre.rightChild,
			}
			t.AjoutIteratif(tt.args.cles)

			// Assertions pour vérifier la taille, la hauteur de l'arbre et la valeur de l'ajout itératif
			var expectedSize, expectedHeight int
			var expectedStringTree string

			switch tt.name {
			case "Add keys iteratively to an empty tree":
				expectedSize = len(tt.args.cles)
				expectedHeight = 2
				expectedStringTree = "Tas{\n          cle: 2,\n          left: nil,\n          right: nil\n        }"

			case "Add keys iteratively to a non-empty tree":
				expectedSize = len(tt.args.cles) + 1
				expectedHeight = 2
				expectedStringTree = "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: nil,\n            right: nil\n          },\n          right: Tas{\n            cle: 5,\n            left: nil,\n            right: nil\n          }\n        }"
			}

			if t.Size != expectedSize {
				t1.Errorf("Taille incorrecte de l'arbre. Got %d, want %d", t.Size, expectedSize)
			}

			if t.height != expectedHeight {
				t1.Errorf("Hauteur incorrecte de l'arbre. Got %d, want %d", t.height, expectedHeight)
			}

			if t.cle != cle1 {
				t1.Errorf("Valeur de racine. Got %d, want %d", t.cle, cle1)
			}

			if t.leftChild.cle != cle2 {
				t1.Errorf("Valeur de fils gauche. Got %d, want %d", t.leftChild.cle, cle2)
			}

			if t.rightChild.cle != cle3 {
				t1.Errorf("Valeur de fils droit. Got %d, want %d", t.rightChild.cle, cle3)
			}

			if expectedStringTree != t._string(0) {
				t1.Errorf("Affichage de l'arbre incorrecte. Got %v, want %v", t._string(0), expectedStringTree)
			}

		})
	}
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
			want: "Tas{\n          cle: 0,\n          left: nil,\n          right: nil\n        }",
		},
		{
			name: "String representation of a non empty tree",
			fields: fields{
				Size:       3,
				height:     2,
				cle:        cle1,
				parent:     nil,
				leftChild:  &Arbre{cle: cle2, Size: 1, height: 1, parent: nil},
				rightChild: &Arbre{cle: cle3, Size: 1, height: 1, parent: nil},
			},
			want: "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: nil,\n            right: nil\n          },\n          right: Tas{\n            cle: 5,\n            left: nil,\n            right: nil\n          }\n        }",
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:       tt.fields.Size,
				height:     tt.fields.height,
				cle:        tt.fields.cle,
				parent:     tt.fields.parent,
				leftChild:  tt.fields.leftChild,
				rightChild: tt.fields.rightChild,
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
				leftChild:  &Arbre{cle: cle2, Size: 1, height: 1, parent: nil},
				rightChild: &Arbre{cle: cle3, Size: 1, height: 1, parent: nil},
			},
			args: args{indent: 0},
			want: "Tas{\n          cle: 12,\n          left: Tas{\n            cle: 16,\n            left: nil,\n            right: nil\n          },\n          right: Tas{\n            cle: 15,\n            left: nil,\n            right: nil\n          }\n        }",
		},
		// Ajouter plus de cas de test au besoin.
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:       tt.fields.Size,
				height:     tt.fields.height,
				cle:        tt.fields.cle,
				parent:     tt.fields.parent,
				leftChild:  tt.fields.leftChild,
				rightChild: tt.fields.rightChild,
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
			want: Arbre{cle: cle.Cle{P1: 0, P2: 20}, height: 1, Size: 1},
		},
		{
			name: "Create tree with key 5",
			args: args{c: cle.Cle{P1: 0, P2: 5}},
			want: Arbre{cle: cle.Cle{P1: 0, P2: 5}, height: 1, Size: 1},
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
