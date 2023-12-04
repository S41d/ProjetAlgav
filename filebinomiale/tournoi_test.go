package filebinomiale

import (
	"projet/cle"
	"reflect"
	"testing"
)

func TestTournoiBinomial_Decapite(t *testing.T) {
	type fields struct {
		Cle     *cle.Cle
		Parent  *TournoiBinomial
		Enfants FileBinomiale
		Degre   int
	}
	tests := []struct {
		name   string
		fields fields
		want   FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := TournoiBinomial{
				Cle:     tt.fields.Cle,
				Parent:  tt.fields.Parent,
				Enfants: tt.fields.Enfants,
				Degre:   tt.fields.Degre,
			}
			if got := tb.Decapite(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decapite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournoiBinomial_EstVide(t *testing.T) {
	type fields struct {
		Cle     *cle.Cle
		Parent  *TournoiBinomial
		Enfants FileBinomiale
		Degre   int
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
			tb := TournoiBinomial{
				Cle:     tt.fields.Cle,
				Parent:  tt.fields.Parent,
				Enfants: tt.fields.Enfants,
				Degre:   tt.fields.Degre,
			}
			if got := tb.EstVide(); got != tt.want {
				t.Errorf("EstVide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournoiBinomial_File(t *testing.T) {
	type fields struct {
		Cle     *cle.Cle
		Parent  *TournoiBinomial
		Enfants FileBinomiale
		Degre   int
	}
	tests := []struct {
		name   string
		fields fields
		want   FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := TournoiBinomial{
				Cle:     tt.fields.Cle,
				Parent:  tt.fields.Parent,
				Enfants: tt.fields.Enfants,
				Degre:   tt.fields.Degre,
			}
			if got := tb.File(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTournoiBinomial_Union(t *testing.T) {
	type fields struct {
		Cle     *cle.Cle
		Parent  *TournoiBinomial
		Enfants FileBinomiale
		Degre   int
	}
	type args struct {
		o TournoiBinomial
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   TournoiBinomial
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := TournoiBinomial{
				Cle:     tt.fields.Cle,
				Parent:  tt.fields.Parent,
				Enfants: tt.fields.Enfants,
				Degre:   tt.fields.Degre,
			}
			if got := tb.Union(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assert(t *testing.T) {
	type args struct {
		condition bool
		msg       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(tt.args.condition, tt.args.msg)
		})
	}
}
