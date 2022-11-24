package base62

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encode-0",
			args: args{
				value: 0,
			},
			want: "0",
		},
		{
			name: "Encode-1",
			args: args{
				value: 1,
			},
			want: "1",
		},
		{
			name: "Encode-9",
			args: args{
				value: 9,
			},
			want: "9",
		},
		{
			name: "Encode-10",
			args: args{
				value: 10,
			},
			want: "a",
		},
		{
			name: "Encode-35",
			args: args{
				value: 35,
			},
			want: "z",
		},
		{
			name: "Encode-36",
			args: args{
				value: 36,
			},
			want: "A",
		},
		{
			name: "Encode-61",
			args: args{
				value: 61,
			},
			want: "Z",
		},
		{
			name: "Encode-62",
			args: args{
				value: 62,
			},
			want: "10",
		},
		{
			name: "Encode-18446744073709551615",
			args: args{
				value: 18446744073709551615,
			},
			want: "lYGhA16ahyf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.value); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Decode-0",
			args: args{
				value: "0",
			},
			want: 0,
		},
		{
			name: "Decode-1",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Decode-9",
			args: args{
				value: "9",
			},
			want: 9,
		},
		{
			name: "Decode-a",
			args: args{
				value: "a",
			},
			want: 10,
		},
		{
			name: "Decode-z",
			args: args{
				value: "z",
			},
			want: 35,
		},
		{
			name: "Decode-A",
			args: args{
				value: "A",
			},
			want: 36,
		},
		{
			name: "Decode-Z",
			args: args{
				value: "Z",
			},
			want: 61,
		},
		{
			name: "Decode-10",
			args: args{
				value: "10",
			},
			want: 62,
		},
		{
			name: "Decode-lYGhA16ahyf",
			args: args{
				value: "lYGhA16ahyf",
			},
			want: 18446744073709551615,
		},
		{
			name: "Decode-lYGhA16ahyg",
			args: args{
				value: "lYGhA16ahyg",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("name: %s Decode() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("name: %s Decode() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := Encode(18446744073709551615)
		_ = a
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("lYGhA16ahyflYGhA16ahyf")
	}
}
