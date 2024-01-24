package encrypt

import (
	"testing"
)

func TestGetMD5(t *testing.T) {
	type testCase struct {
		Name     string
		Input    []byte
		Expected string
	}

	testCases := []testCase{
		{
			Name:     "empty input",
			Input:    []byte(""),
			Expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			Name:     "valid input",
			Input:    []byte("Hello, world!"),
			Expected: "6cd3556deb0da54bca060b4c39479839",
		},
		{
			Name:     "numeric input",
			Input:    []byte("1234567890"),
			Expected: "e807f1fcf82d132f9bb018ca6738a19f",
		},
		{
			Name:     "special character input",
			Input:    []byte("!@#$%^&*()"),
			Expected: "05b28d17a7b6e7024b6e5d8cc43a8bf7",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := getMD5(tc.Input)
			if result != tc.Expected {
				t.Errorf("expected: %s, got: %s", tc.Expected, result)
			}
		})
	}
}
