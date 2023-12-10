package md5

import (
	"crypto/md5"
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	type args struct {
		x uint32
		y uint32
		z uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Case 1", args{0x12345678, 0x87654321, 0xabcdef01}, 0xABEDEB21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := F(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("F() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestG(t *testing.T) {
	type args struct {
		x uint32
		y uint32
		z uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Case 1", args{0x12345678, 0x87654321, 0xabcdef01}, 0x06244620},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := G(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("G() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestH(t *testing.T) {
	type args struct {
		x uint32
		y uint32
		z uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Case 1", args{0x12345678, 0x87654321, 0xabcdef01}, 0x3E9CFA58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := H(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("H() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI(t *testing.T) {
	type args struct {
		x uint32
		y uint32
		z uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Case 1", args{0x12345678, 0x87654321, 0xabcdef01}, 0xD15315DF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeK(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Initialization Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializeK()

			expectedValues := []uint32{
				0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
				0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
				0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
				0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
				0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
				0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
				0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
				0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
				0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
				0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
				0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
				0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
				0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
				0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
				0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
				0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
			}

			for i, value := range k[:64] {
				if value != expectedValues[i] {
					t.Errorf("initializeK() - Value mismatch at index %d. Got: %d, Expected: %d", i, value, expectedValues[i])
				}
			}
		})
	}
}

func Test_leftRotate(t *testing.T) {
	type args struct {
		x uint32
		c int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"Case 1", args{0x12345678, 4}, 0x23456781},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRotate(tt.args.x, tt.args.c); got != tt.want {
				t.Errorf("leftRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_md5(t *testing.T) {
	type args struct {
		message []byte
	}
	tests := []struct {
		name string
		args args
		want [16]byte
	}{
		{name: "Case 1", args: args{[]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.")}, want: md5.Sum([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus."))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5New(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("md5New() = %v, want %v", got, tt.want)
			}
		})
	}
}
