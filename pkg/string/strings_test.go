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
		{
			name: "empty_string",
			args: args{
				str: "",
			},
			want: "",
		},
		{
			name: "single_char",
			args: args{
				str: "h",
			},
			want: "H",
		},
		{
			name: "normal_word",
			args: args{
				str: "hello",
			},
			want: "Hello",
		},
		{
			name: "already_uppercase",
			args: args{
				str: "Hello",
			},
			want: "Hello",
		},
		{
			name: "with_spaces",
			args: args{
				str: "hello world",
			},
			want: "Hello world",
		},
		{
			name: "with_special_chars",
			args: args{
				str: "hello!world",
			},
			want: "Hello!world",
		},
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

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "zero_length",
			length: 0,
		},
		{
			name:   "short_string",
			length: 8,
		},
		{
			name:   "medium_string",
			length: 16,
		},
		{
			name:   "long_string",
			length: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomString(tt.length)
			if len(got) != tt.length {
				t.Errorf("GenerateRandomString() length = %v, want %v", len(got), tt.length)
			}

			// 生成两次，验证随机性
			got2 := GenerateRandomString(tt.length)
			if tt.length > 0 && got == got2 {
				t.Errorf("GenerateRandomString() should generate different strings for the same length")
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty string", "", true},
		{"space only", " ", true},
		{"spaces and tabs", "  \t  ", true},
		{"newlines", "\n\r", true},
		{"normal text", "hello", false},
		{"text with spaces", " hello ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty string", "", true},
		{"space only", " ", false},
		{"normal text", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.s); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"simple word", "hello", "olleh"},
		{"unicode", "你好", "好你"},
		{"mixed", "a1b2", "2b1a"},
		{"emoji", "😀😁", "😁😀"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		want   bool
	}{
		{"exact match", "hello", "hello", true},
		{"case insensitive", "Hello", "hello", true},
		{"substring", "Hello World", "world", true},
		{"not found", "hello", "hi", false},
		{"empty substr", "hello", "", true},
		{"empty string", "", "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsIgnoreCase(tt.s, tt.substr); got != tt.want {
				t.Errorf("ContainsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{"no truncation needed", "hello", 10, "hello"},
		{"truncate with ellipsis", "hello world", 8, "hello..."},
		{"truncate short", "hello", 3, "hel"},
		{"zero length", "hello", 0, ""},
		{"negative length", "hello", -1, ""},
		{"unicode", "你好世界", 4, "你好世界"},
		{"unicode truncate", "你好世界", 2, "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.s, tt.maxLen); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padChar rune
		want    string
	}{
		{"no padding needed", "hello", 3, ' ', "hello"},
		{"pad with spaces", "hi", 5, ' ', "   hi"},
		{"pad with zeros", "123", 6, '0', "000123"},
		{"zero length", "hello", 0, ' ', "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.s, tt.length, tt.padChar); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padChar rune
		want    string
	}{
		{"no padding needed", "hello", 3, ' ', "hello"},
		{"pad with spaces", "hi", 5, ' ', "hi   "},
		{"pad with zeros", "123", 6, '0', "123000"},
		{"zero length", "hello", 0, ' ', "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.s, tt.length, tt.padChar); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveSpaces(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"no spaces", "hello", "hello"},
		{"single space", "hel lo", "hello"},
		{"multiple spaces", "h e l l o", "hello"},
		{"leading and trailing", " hello ", "hello"},
		{"only spaces", "   ", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSpaces(tt.s); got != tt.want {
				t.Errorf("RemoveSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"pure digits", "12345", true},
		{"with letters", "123a", false},
		{"with spaces", "1 2 3", false},
		{"empty string", "", false},
		{"single digit", "5", true},
		{"with symbols", "123!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"simple words", "hello world", "helloWorld"},
		{"with underscores", "hello_world_test", "helloWorldTest"},
		{"with hyphens", "hello-world-test", "helloWorldTest"},
		{"mixed separators", "hello_world-test", "helloWorldTest"},
		{"already camel", "helloWorld", "helloworld"},
		{"empty string", "", ""},
		{"single word", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"simple words", "hello world", "HelloWorld"},
		{"with underscores", "hello_world_test", "HelloWorldTest"},
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.s); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"camel case", "helloWorld", "hello_world"},
		{"pascal case", "HelloWorld", "hello_world"},
		{"spaces", "hello world", "hello_world"},
		{"mixed", "helloWorldTest", "hello_world_test"},
		{"already snake", "hello_world", "hello_world"},
		{"empty string", "", ""},
		{"single word", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.s); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"camel case", "helloWorld", "hello-world"},
		{"pascal case", "HelloWorld", "hello-world"},
		{"spaces", "hello world", "hello-world"},
		{"already kebab", "hello-world", "hello-world"},
		{"empty string", "", ""},
		{"single word", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabCase(tt.s); got != tt.want {
				t.Errorf("KebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"simple sentence", "hello world", "Hello World"},
		{"mixed case", "hELLo WoRLd", "Hello World"},
		{"single word", "hello", "Hello"},
		{"empty string", "", ""},
		{"multiple spaces", "hello   world", "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Title(tt.s); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCenter(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		length  int
		padChar rune
		want    string
	}{
		{"even padding", "hi", 6, ' ', "  hi  "},
		{"odd padding", "hi", 5, ' ', " hi  "},
		{"no padding needed", "hello", 3, ' ', "hello"},
		{"with stars", "test", 8, '*', "**test**"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Center(tt.s, tt.length, tt.padChar); got != tt.want {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		substr string
		want   int
	}{
		{"simple count", "hello world", "l", 3},
		{"no matches", "hello", "x", 0},
		{"overlapping", "aaa", "aa", 1},
		{"empty substr", "hello", "", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.s, tt.substr); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPrefixIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		prefix string
		want   bool
	}{
		{"exact match", "hello", "hello", true},
		{"case insensitive", "Hello", "hello", true},
		{"prefix only", "Hello World", "hello", true},
		{"no match", "hello", "hi", false},
		{"empty prefix", "hello", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefixIgnoreCase(tt.s, tt.prefix); got != tt.want {
				t.Errorf("HasPrefixIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSuffixIgnoreCase(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		suffix string
		want   bool
	}{
		{"exact match", "hello", "hello", true},
		{"case insensitive", "Hello", "HELLO", true},
		{"suffix only", "Hello World", "world", true},
		{"no match", "hello", "hi", false},
		{"empty suffix", "hello", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSuffixIgnoreCase(tt.s, tt.suffix); got != tt.want {
				t.Errorf("HasSuffixIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"simple sentence", "hello world", 2},
		{"single word", "hello", 1},
		{"empty string", "", 0},
		{"multiple spaces", "hello   world   test", 3},
		{"with tabs", "hello\tworld", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.s); got != tt.want {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		length int
		want   string
	}{
		{"normal case", "hello world", 5, "hello"},
		{"longer than string", "hi", 5, "hi"},
		{"zero length", "hello", 0, ""},
		{"negative length", "hello", -1, ""},
		{"unicode", "你好世界", 2, "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Left(tt.s, tt.length); got != tt.want {
				t.Errorf("Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		length int
		want   string
	}{
		{"normal case", "hello world", 5, "world"},
		{"longer than string", "hi", 5, "hi"},
		{"zero length", "hello", 0, ""},
		{"negative length", "hello", -1, ""},
		{"unicode", "你好世界", 2, "世界"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Right(tt.s, tt.length); got != tt.want {
				t.Errorf("Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"only letters", "hello", true},
		{"with numbers", "hello123", false},
		{"with spaces", "hello world", false},
		{"with symbols", "hello!", false},
		{"empty string", "", false},
		{"unicode letters", "你好", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.s); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"letters and numbers", "hello123", true},
		{"only letters", "hello", true},
		{"only numbers", "123", true},
		{"with spaces", "hello 123", false},
		{"with symbols", "hello!", false},
		{"empty string", "", false},
		{"unicode", "你好123", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumeric(tt.s); got != tt.want {
				t.Errorf("IsAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
