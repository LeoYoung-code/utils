package string

import (
	"reflect"
	"testing"
)

func TestGoSanitized(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "hello world",
			},
			want: "hello_world",
		},
		{
			name: "2",
			args: args{
				s: "hello world a",
			},
			want: "hello_world_a",
		},
		{
			name: "3",
			args: args{
				s: "hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoSanitized(tt.args.s); got != tt.want {
				t.Errorf("GoSanitized() = %v, want %v", got, tt.want)
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
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UcFirst(tt.args.str); got != tt.want {
				t.Errorf("UcFirst() = %v, want %v", got, tt.want)
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
