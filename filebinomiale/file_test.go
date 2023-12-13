package filebinomiale

import (
	"fmt"
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

func TestFileBinomiale_UnionFile(t *testing.T) {
	c1 := &cle.Cle{P1: 0, P2: 3}
	c2 := &cle.Cle{P1: 0, P2: 7}
	c3 := &cle.Cle{P1: 0, P2: 4}
	c4 := &cle.Cle{P1: 0, P2: 11}

	file1 := FileBinomiale{TournoiBinomial{Cle: c1, Degre: 1, Enfants: FileBinomiale{TournoiBinomial{Cle: c2, Degre: 0, Enfants: FileBinomiale{}}}}}
	file2 := FileBinomiale{TournoiBinomial{Cle: c3, Degre: 1, Enfants: FileBinomiale{TournoiBinomial{Cle: c4, Degre: 0, Enfants: FileBinomiale{}}}}}

	tests := []struct {
		name string
		fb   FileBinomiale
		args FileBinomiale
		want FileBinomiale
	}{
		{
			name: "Test UnionFile",
			fb:   file1,
			args: file2,
			want: FileBinomiale{
				TournoiBinomial{Cle: c1, Degre: 2, Enfants: FileBinomiale{
					TournoiBinomial{Cle: c2, Degre: 0, Enfants: FileBinomiale{}},
					TournoiBinomial{Cle: c3, Degre: 1, Enfants: FileBinomiale{
						TournoiBinomial{Cle: c4, Degre: 0, Enfants: FileBinomiale{}},
					}},
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fb.Union(tt.args)
			got.Size()
			if !reflect.DeepEqual(got[0].Cle.P2, 3) {
				t.Errorf("UnionFile() = %v, want 3", got[0].Cle.P2)
			}
			if !reflect.DeepEqual(got[0].Enfants[0].Cle.P2, 7) {
				t.Errorf("UnionFile() = %v, want 7", got[0].Enfants[0].Cle.P2)
			}
			if !reflect.DeepEqual(got[0].Enfants[1].Cle.P2, 4) {
				t.Errorf("UnionFile() = %v, want 4", got[0].Enfants[1].Cle.P2)
			}
			if !reflect.DeepEqual(got[0].Enfants[1].Enfants[0].Cle.P2, 11) {
				t.Errorf("UnionFile() = %v, want 11", got[0].Enfants[1].Enfants[0].Cle.P2)
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
	c1 := &cle.Cle{P1: 0, P2: 3}
	c2 := &cle.Cle{P1: 0, P2: 7}
	c3 := &cle.Cle{P1: 0, P2: 4}
	c4 := &cle.Cle{P1: 0, P2: 11}

	file1 := FileBinomiale{TournoiBinomial{Cle: c1, Degre: 1, Enfants: FileBinomiale{TournoiBinomial{Cle: c2, Degre: 0, Enfants: FileBinomiale{}}}}}
	file2 := FileBinomiale{TournoiBinomial{Cle: c3, Degre: 1, Enfants: FileBinomiale{TournoiBinomial{Cle: c4, Degre: 0, Enfants: FileBinomiale{}}}}}
	union := file1.Union(file2)

	fmt.Println(union.Size())
	union.SupprMin()
	fmt.Println(union.Size())
	union.SupprMin()
	fmt.Println(union.Size())
	union.SupprMin()
	fmt.Println(union.Size())
	union.SupprMin()
	fmt.Println(union.Size())
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
