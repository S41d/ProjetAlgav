package filebinomiale

import (
	"projet/cle"
	"reflect"
	"testing"
)

func TestConstruction(t *testing.T) {
	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name string
		args args
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Construction(tt.args.cles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Construction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_Ajout(t *testing.T) {
	type args struct {
		c cle.Cle
	}
	tests := []struct {
		name string
		fb   FileBinomiale
		args args
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.Ajout(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ajout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_AjoutMin(t *testing.T) {
	type args struct {
		tb TournoiBinomial
	}
	tests := []struct {
		name string
		fb   FileBinomiale
		args args
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.AjoutMin(tt.args.tb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AjoutMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_EstVide(t *testing.T) {
	tests := []struct {
		name string
		fb   FileBinomiale
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.EstVide(); got != tt.want {
				t.Errorf("EstVide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_MinDeg(t *testing.T) {
	tests := []struct {
		name string
		fb   FileBinomiale
		want TournoiBinomial
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.MinDeg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinDeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_Reste(t *testing.T) {
	tests := []struct {
		name string
		fb   FileBinomiale
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.Reste(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reste() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_SupprMin(t *testing.T) {
	tests := []struct {
		name string
		fb   FileBinomiale
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.SupprMin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SupprMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileBinomiale_Union(t *testing.T) {
	type args struct {
		o FileBinomiale
	}
	tests := []struct {
		name string
		fb   FileBinomiale
		args args
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fb.Union(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uFret(t *testing.T) {
	type args struct {
		f1 FileBinomiale
		f2 FileBinomiale
		t  TournoiBinomial
	}
	tests := []struct {
		name string
		args args
		want FileBinomiale
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uFret(tt.args.f1, tt.args.f2, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uFret() = %v, want %v", got, tt.want)
			}
		})
	}
}
