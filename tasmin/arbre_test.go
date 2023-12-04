package tasmin

import (
	"projet/cle"
	"reflect"
	"testing"
)

func TestArbre_Ajout(t1 *testing.T) {
	type fields struct {
		Size       int
		height     int
		cle        cle.Cle
		parent     *Arbre
		leftChild  *Arbre
		rightChild *Arbre
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Arbre{
				Size:       tt.fields.Size,
				height:     tt.fields.height,
				cle:        tt.fields.cle,
				parent:     tt.fields.parent,
				leftChild:  tt.fields.leftChild,
				rightChild: tt.fields.rightChild,
			}
			t.Ajout(tt.args.c)
		})
	}
}

func TestArbre_AjoutIteratif(t1 *testing.T) {
	type fields struct {
		Size       int
		height     int
		cle        cle.Cle
		parent     *Arbre
		leftChild  *Arbre
		rightChild *Arbre
	}
	type args struct {
		cles []cle.Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
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
			t.AjoutIteratif(tt.args.cles)
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setSigHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setSigHandler()
		})
	}
}
