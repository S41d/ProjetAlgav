package tasmin

import (
	_ "errors"
	"projet/cle"
	"reflect"
	_ "strings"
	"testing"
)

func TestConstruction(t *testing.T) {
	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name string
		args args
		want Tableau
	}{
		{
			name: "Cas normal",
			args: args{cles: []cle.Cle{{P1: 0, P2: 13}, {P1: 0, P2: 14}, {P1: 0, P2: 15}, {P1: 0, P2: 12}, {P1: 0, P2: 2}, {P1: 0, P2: 4}, {P1: 0, P2: 8}, {P1: 0, P2: 7}, {P1: 0, P2: 6}, {P1: 0, P2: 10}, {P1: 0, P2: 5}}},
			want: Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 4}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Construction(tt.args.cles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Construction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_Ajout(t *testing.T) {
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name string
		t    Tableau
		args args
		want Tableau
	}{
		{
			name: "Cas normal",
			t:    Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}},
			args: args{c: cle.Cle{P1: 0, P2: 4}},
			want: Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 4}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 0, P2: 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.Ajout(tt.args.c)

			if !reflect.DeepEqual(tt.t, tt.want) {
				t.Errorf("Ajout() got = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTableau_AjoutIteratif(t *testing.T) {
	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name string
		t    Tableau
		args args
		want Tableau
	}{
		{
			name: "Cas normal",
			t:    Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}},
			args: args{cles: []cle.Cle{{P1: 0, P2: 4}, {P1: 0, P2: 25}}},
			want: Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 4}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 25}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.AjoutIteratif(tt.args.cles)

			if !reflect.DeepEqual(tt.t, tt.want) {
				t.Errorf("AjoutIteratif() got = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTableau_EnfDroit(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		ta   Tableau
		args args
		want int
	}{
		{
			name: "Index positif",
			ta: Tableau{
				cle.Cle{P1: 0, P2: 2},
				cle.Cle{P1: 0, P2: 6},
				cle.Cle{P1: 0, P2: 5},
				// Ajouter d'autres éléments du tableau si nécessaire
			},
			args: args{i: 0},
			want: 2, // L'indice correspondant au fils droit de {P1: 0, P2: 5} est 2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ta.EnfDroit(tt.args.i); got != tt.want {
				t.Errorf("EnfDroit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_EnfGauche(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		ta   Tableau
		args args
		want int
	}{
		{
			name: "Index positif",
			ta: Tableau{
				cle.Cle{P1: 0, P2: 2},
				cle.Cle{P1: 0, P2: 6},
				cle.Cle{P1: 0, P2: 5},
				// Ajouter d'autres éléments du tableau si nécessaire
			},
			args: args{i: 0},
			want: 1, // L'indice correspondant à l'élément {P1: 0, P2: 6}
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ta.EnfGauche(tt.args.i); got != tt.want {
				t.Errorf("EnfGauche() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_Parent(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		ta   Tableau
		args args
		want int
	}{
		{
			name: "Index positif",
			ta: Tableau{
				cle.Cle{P1: 0, P2: 2},
				cle.Cle{P1: 0, P2: 6},
				cle.Cle{P1: 0, P2: 5},
				// Ajouter d'autres éléments du tableau si nécessaire
			},
			args: args{i: 2},
			want: 0, // L'indice correspondant au parent de {P1: 2, P2: 2} est 0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ta.Parent(tt.args.i); got != tt.want {
				t.Errorf("Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_String(t *testing.T) {
	tests := []struct {
		name string
		t    Tableau
		want string
	}{
		{
			name: "Cas normal",
			t: Tableau{
				cle.Cle{P1: 0, P2: 2},
				cle.Cle{P1: 0, P2: 6},
				cle.Cle{P1: 0, P2: 5},
				cle.Cle{P1: 0, P2: 10},
				cle.Cle{P1: 0, P2: 13},
				cle.Cle{P1: 0, P2: 7},
				cle.Cle{P1: 0, P2: 8},
				cle.Cle{P1: 0, P2: 12},
				cle.Cle{P1: 0, P2: 15},
				cle.Cle{P1: 0, P2: 14},
			},
			want: "Tas{\n  cle: 2,\n  left: Tas{\n    cle: 6,\n    left: Tas{\n      cle: 10,\n      left: Tas{\n        cle: 12,\n        left: nil,\n        right: nil\n      },\n      right: Tas{\n        cle: 15,\n        left: nil,\n        right: nil\n      }\n    },\n    right: Tas{\n      cle: 13,\n      left: Tas{\n        cle: 14,\n        left: nil,\n        right: nil\n      },\n      right: nil\n    }\n  },\n  right: Tas{\n    cle: 5,\n    left: Tas{\n      cle: 7,\n      left: nil,\n      right: nil\n    },\n    right: Tas{\n      cle: 8,\n      left: nil,\n      right: nil\n    }\n  }\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_SupprMin(t *testing.T) {
	tests := []struct {
		name    string
		t       Tableau
		want    cle.Cle
		wantErr bool
		wantTab Tableau
	}{
		{
			name:    "Cas normal",
			t:       Tableau{cle.Cle{P1: 1, P2: 2}, cle.Cle{P1: 1, P2: 6}, cle.Cle{P1: 1, P2: 5}},
			want:    cle.Cle{P1: 1, P2: 2},
			wantErr: false,
			wantTab: Tableau{cle.Cle{P1: 1, P2: 5}, cle.Cle{P1: 1, P2: 6}},
		},
		{
			name:    "Cas du tableau vide",
			t:       Tableau{},
			want:    cle.Cle{},
			wantErr: true,
			wantTab: Tableau{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.SupprMin()
			if (err != nil) != tt.wantErr {
				t.Errorf("SupprMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SupprMin() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.t, tt.wantTab) {
				t.Errorf("SupprMin() tab = %v, wantTab %v", tt.t, tt.wantTab)
			}
		})
	}
}

func TestTableau_Union(t *testing.T) {
	type args struct {
		o Tableau
	}
	tests := []struct {
		name string
		t    Tableau
		args args
		want Tableau
	}{
		{
			name: "Test Union avec deux tableaux vides",
			t:    Tableau{},
			args: args{o: Tableau{}},
			want: Tableau{},
		},
		{
			name: "Test Union avec des éléments",
			t:    Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 1, P2: 1}},
			args: args{o: Tableau{cle.Cle{P1: 0, P2: 1}, cle.Cle{P1: 0, P2: 3}, cle.Cle{P1: 0, P2: 17}, cle.Cle{P1: 0, P2: 19}, cle.Cle{P1: 0, P2: 36}}},
			want: Tableau{cle.Cle{P1: 0, P2: 1}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 3}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 1, P2: 1}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 17}, cle.Cle{P1: 0, P2: 19}, cle.Cle{P1: 0, P2: 36}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Union(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau__string(t *testing.T) {
	type args struct {
		i      int
		indent int
	}
	tests := []struct {
		name string
		t    Tableau
		args args
		want string
	}{
		{
			name: "Cas normal",
			t: Tableau{
				cle.Cle{P1: 0, P2: 2},
				cle.Cle{P1: 0, P2: 6},
				cle.Cle{P1: 0, P2: 5},
				cle.Cle{P1: 0, P2: 10},
				cle.Cle{P1: 0, P2: 13},
				cle.Cle{P1: 0, P2: 7},
				cle.Cle{P1: 0, P2: 8},
				cle.Cle{P1: 0, P2: 12},
				cle.Cle{P1: 0, P2: 15},
				cle.Cle{P1: 0, P2: 14},
			},
			args: args{i: 0, indent: 0},
			// want: "Tas{\n          cle: 2,\n          left: Tas{\n            cle: 6,\n            left: Tas{\n              cle: 10,\n              left: Tas{\n                cle: 12,\n                left: nil,\n                right: nil\n              },\n              right: Tas{\n                cle: 15,\n                left: nil,\n                right: nil\n              }\n            },\n            right: Tas{\n              cle: 13,\n              left: Tas{\n                cle: 14,\n                left: nil,\n                right: nil\n              },\n              right: nil\n            }\n          },\n          right: Tas{\n            cle: 5,\n            left: Tas{\n              cle: 7,\n              left: nil,\n              right: nil\n            },\n            right: Tas{\n              cle: 8,\n              left: nil,\n              right: nil\n            }\n          }\n        }"},
			want: "Tas{\n  cle: 2,\n  left: Tas{\n    cle: 6,\n    left: Tas{\n      cle: 10,\n      left: Tas{\n        cle: 12,\n        left: nil,\n        right: nil\n      },\n      right: Tas{\n        cle: 15,\n        left: nil,\n        right: nil\n      }\n    },\n    right: Tas{\n      cle: 13,\n      left: Tas{\n        cle: 14,\n        left: nil,\n        right: nil\n      },\n      right: nil\n    }\n  },\n  right: Tas{\n    cle: 5,\n    left: Tas{\n      cle: 7,\n      left: nil,\n      right: nil\n    },\n    right: Tas{\n      cle: 8,\n      left: nil,\n      right: nil\n    }\n  }\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t._string(tt.args.i, tt.args.indent)

			if got != tt.want {
				t.Errorf("_string() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableau_trier(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		t    Tableau
		args args
		want Tableau
	}{
		{
			name: "Cas normal",
			t:    Tableau{cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 13}},
			args: args{index: 0},
			want: Tableau{cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 5}, cle.Cle{P1: 0, P2: 13}},
			// want: Tableau{cle.Cle{P1: 0, P2: 2}, cle.Cle{P1: 0, P2: 6}, cle.Cle{P1: 0, P2: 7}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 13}, cle.Cle{P1: 0, P2: 8}, cle.Cle{P1: 0, P2: 10}, cle.Cle{P1: 0, P2: 12}, cle.Cle{P1: 0, P2: 15}, cle.Cle{P1: 0, P2: 14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.trier(tt.args.index)

			if !reflect.DeepEqual(tt.t, tt.want) {
				t.Errorf("trier() got = %v, want %v", tt.t, tt.want)
			}
		})
	}
}
