package cle

import (
	"reflect"
	"testing"
)

func TestCle_Inf(t *testing.T) {
	type fields struct {
		p1 uint64
		p2 uint64
	}
	type args struct {
		o Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test inférieure P1",
			fields: fields{
				p1: 5,
				p2: 10,
			},
			args: args{
				o: Cle{
					P1: 7,
					P2: 15,
				},
			},
			want: true,
		},
		{
			name: "Test inférieure P2",
			fields: fields{
				p1: 5,
				p2: 10,
			},
			args: args{
				o: Cle{
					P1: 5,
					P2: 15,
				},
			},
			want: true,
		},
		{
			name: "Test non inférieure P1",
			fields: fields{
				p1: 20,
				p2: 30,
			},
			args: args{
				o: Cle{
					P1: 15,
					P2: 35,
				},
			},
			want: false,
		},
		{
			name: "Test non inférieure P2",
			fields: fields{
				p1: 20,
				p2: 35,
			},
			args: args{
				o: Cle{
					P1: 20,
					P2: 30,
				},
			},
			want: false,
		},
		{
			name: "Test non inférieure P1 P2",
			fields: fields{
				p1: 20,
				p2: 35,
			},
			args: args{
				o: Cle{
					P1: 20,
					P2: 35,
				},
			},
			want: false,
		},
		// Ajouter d'autres cas de test au besoin
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cle{
				P1: tt.fields.p1,
				P2: tt.fields.p2,
			}
			if got := c.Inf(tt.args.o); got != tt.want {
				t.Errorf("Inf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCle_Eg(t *testing.T) {
	type fields struct {
		p1 uint64
		p2 uint64
	}
	type args struct {
		o Cle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test égaux",
			fields: fields{
				p1: 123,
				p2: 456,
			},
			args: args{
				o: Cle{
					P1: 123,
					P2: 456,
				},
			},
			want: true,
		},
		{
			name: "Test non égaux P1",
			fields: fields{
				p1: 789,
				p2: 101,
			},
			args: args{
				o: Cle{
					P1: 111,
					P2: 222,
				},
			},
			want: false,
		},
		{
			name: "Test non égaux P2",
			fields: fields{
				p1: 789,
				p2: 101,
			},
			args: args{
				o: Cle{
					P1: 789,
					P2: 222,
				},
			},
			want: false,
		},
		// Ajouter d'autres cas de test au besoin
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cle{
				P1: tt.fields.p1,
				P2: tt.fields.p2,
			}
			if got := c.Eg(tt.args.o); got != tt.want {
				t.Errorf("Eg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCle_HexaString(t *testing.T) {
	type fields struct {
		p1 uint64
		p2 uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"P1 P2 non Zero", fields{p1: 123, p2: 456}, "7b1c8"},
		{"P1 Zero", fields{p1: 0, p2: 789}, "315"},
		{"P2 Zero", fields{p1: 987, p2: 0}, "3db"},
		{"P1 P2 Zero", fields{p1: 0, p2: 0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cle{
				P1: tt.fields.p1,
				P2: tt.fields.p2,
			}
			if got := c.HexaString(); got != tt.want {
				t.Errorf("HexaString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCle_BinaryString(t *testing.T) {
	type fields struct {
		p1 uint64
		p2 uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test avec des valeurs non nulles",
			fields: fields{p1: 10, p2: 255},
			want:   "10100000000000000000000000000000000000000000000000000000000011111111\n",
		},
		{
			name:   "Test avec des valeurs nulles",
			fields: fields{p1: 0, p2: 0},
			want:   "00000000000000000000000000000000000000000000000000000000000000000\n",
		},
		// Ajouter d'autres cas de test selon vos besoins.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cle{
				P1: tt.fields.p1,
				P2: tt.fields.p2,
			}
			if got := c.BinaryString(); got != tt.want {
				t.Errorf("BinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCle_DecimalString(t *testing.T) {
	type fields struct {
		p1 uint64
		p2 uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test sans zero",
			fields: fields{p1: 1, p2: 123456},
			want:   "18446744073709675072",
		},
	}
	// p1 = 0000000000000000000000000000000000000000000000000000000000000001
	// p2 = 0000000000000000000000000000000000000000000000011110001001000000
	// val = 10000000000000000000000000000000000000000000000011110001001000000 = 18446744073709675072

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cle{
				P1: tt.fields.p1,
				P2: tt.fields.p2,
			}
			if got := c.DecimalString(); got != tt.want {
				t.Errorf("DecimalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToCle(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want Cle
	}{
		{
			name: "Test HexToCle lowercase hexadecimal string",
			args: args{hex: "df6943ba6d51464f6b02157933bdd9ad"},
			want: Cle{P1: 0xdf6943ba6d51464f, P2: 0x6b02157933bdd9ad},
		},
		{
			name: "Test HexToCle mixed case hexadecimal string",
			args: args{hex: "DF6943BA6D51464F6B02157933Bdd9ad"},
			want: Cle{P1: 0xdf6943ba6d51464f, P2: 0x6b02157933bdd9ad},
		},
		{
			name: "Test HexToCle '0x' prefix",
			args: args{hex: "0xdf6943ba6d51464f6b02157933bdd9ad"},
			want: Cle{P1: 0xdf6943ba6d51464f, P2: 0x6b02157933bdd9ad},
		},
		{
			name: "Test HexToCle commence par des 0",
			args: args{hex: "0059529bcc28157b57548bee786a1940"},
			want: Cle{P1: 0x0059529bcc28157b, P2: 0x57548bee786a1940},
		},
		// Ajouter d'autres cas de test au besoin.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToCle(tt.args.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToCle() = %v, want %v", got, tt.want)
			}
		})
	}
}
