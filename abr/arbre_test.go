package arbrerecherche

import (
	"projet/cle"
	"reflect"
	"testing"
)

func TestArbreRecherche_Ajout(t *testing.T) {
	type fields struct {
		Elt      *cle.Cle
		Parent   *ArbreRecherche
		SaGauche *ArbreRecherche
		SaDroit  *ArbreRecherche
	}
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArbreRecherche{
				Elt:      tt.fields.Elt,
				Parent:   tt.fields.Parent,
				SaGauche: tt.fields.SaGauche,
				SaDroit:  tt.fields.SaDroit,
			}
			a.Ajout(tt.args.c)
		})
	}
}

func TestArbreRecherche_Contient(t *testing.T) {
	type fields struct {
		Elt      *cle.Cle
		Parent   *ArbreRecherche
		SaGauche *ArbreRecherche
		SaDroit  *ArbreRecherche
	}
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArbreRecherche{
				Elt:      tt.fields.Elt,
				Parent:   tt.fields.Parent,
				SaGauche: tt.fields.SaGauche,
				SaDroit:  tt.fields.SaDroit,
			}
			if got := a.Contient(tt.args.c); got != tt.want {
				t.Errorf("Contient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArbreRecherche_EstArbreVide(t *testing.T) {
	type fields struct {
		Elt      *cle.Cle
		Parent   *ArbreRecherche
		SaGauche *ArbreRecherche
		SaDroit  *ArbreRecherche
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArbreRecherche{
				Elt:      tt.fields.Elt,
				Parent:   tt.fields.Parent,
				SaGauche: tt.fields.SaGauche,
				SaDroit:  tt.fields.SaDroit,
			}
			if got := a.EstArbreVide(); got != tt.want {
				t.Errorf("EstArbreVide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArbreRecherche_Racine(t *testing.T) {
	type fields struct {
		Elt      *cle.Cle
		Parent   *ArbreRecherche
		SaGauche *ArbreRecherche
		SaDroit  *ArbreRecherche
	}
	tests := []struct {
		name   string
		fields fields
		want   ArbreRecherche
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArbreRecherche{
				Elt:      tt.fields.Elt,
				Parent:   tt.fields.Parent,
				SaGauche: tt.fields.SaGauche,
				SaDroit:  tt.fields.SaDroit,
			}
			if got := a.Racine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Racine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArbreRecherche_Suppr(t *testing.T) {
	type fields struct {
		Elt      *cle.Cle
		Parent   *ArbreRecherche
		SaGauche *ArbreRecherche
		SaDroit  *ArbreRecherche
	}
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArbreRecherche{
				Elt:      tt.fields.Elt,
				Parent:   tt.fields.Parent,
				SaGauche: tt.fields.SaGauche,
				SaDroit:  tt.fields.SaDroit,
			}
			a.Suppr(tt.args.c)
		})
	}
}

func TestNewArbreRecherche(t *testing.T) {
	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name string
		args args
		want ArbreRecherche
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArbreRecherche(tt.args.cles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArbreRecherche() = %v, want %v", got, tt.want)
			}
		})
	}
}
