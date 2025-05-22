package string

import (
	"reflect"
	"testing"
)

func TestGoSanitized(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Simple space",
			input: "hello world",
			want:  "hello_world",
		},
		{
			name:  "Keyword conflict",
			input: "func",
			want:  "_func",
		},
		{
			name:  "Invalid start char",
			input: "123abc",
			want:  "_123abc",
		},
		{
			name:  "Mixed special chars",
			input: "a-b.c@d",
			want:  "a_b_c_d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GoSanitized(tt.input)
			if got != tt.want {
				t.Errorf("GoSanitized(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLcFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "hello world",
			},
			want: "hello world",
		},
		{
			name: "2",
			args: args{
				str: "hello",
			},
			want: "hello",
		},
		{
			name: "3",
			args: args{
				str: "HELLO",
			},
			want: "hELLO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LcFirst(tt.args.str); got != tt.want {
				t.Errorf("LcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUcFirst(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Single lowercase",
			input: "a",
			want:  "A",
		},
		{
			name:  "Single uppercase",
			input: "A",
			want:  "A",
		},
		{
			name:  "Multiple characters",
			input: "hello world",
			want:  "Hello world",
		},
		{
			name:  "Unicode character",
			input: "👋hi",
			want:  "👋hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UcFirst(tt.input)
			if got != tt.want {
				t.Errorf("UcFirst(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func Test_b2s(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				b: []byte("hello world"),
			},
			want: "hello world",
		},
		{
			name: "2",
			args: args{
				b: []byte(""),
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				b: []byte("🏷"),
			},
			want: "🏷",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b2s(tt.args.b); got != tt.want {
				t.Errorf("b2s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_s2b(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				s: "hello world",
			},
			want: []byte("hello world"),
		},
		{
			name: "2",
			args: args{
				s: "",
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				s: "🏷",
			},
			want: []byte("🏷"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := s2b(tt.args.s); !reflect.DeepEqual(gotB, tt.want) {
				t.Errorf("s2b() = %v, want %v", gotB, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				b: []byte("hello world"),
			},
			want: "hello world",
		},
		{
			name: "2",
			args: args{
				b: []byte(""),
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				b: []byte("🏷"),
			},
			want: "🏷",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				s: "hello world",
			},
			want: []byte("hello world"),
		},
		{
			name: "2",
			args: args{
				s: "",
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				s: "🏷",
			},
			want: []byte("🏷"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUcFirst(b *testing.B) {
	str := "hello world"
	for i := 0; i < b.N; i++ {
		UcFirst(str)
	}
}

func BenchmarkLcFirst(b *testing.B) {
	str := "HELLO WORLD"
	for i := 0; i < b.N; i++ {
		LcFirst(str)
	}
}

func BenchmarkGoSanitized(b *testing.B) {
	str := "hello@world.com"
	for i := 0; i < b.N; i++ {
		GoSanitized(str)
	}
}
