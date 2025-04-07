package pdf

import (
	"os"
	"testing"
)

func TestGeneratePDF(t *testing.T) {
	err := GeneratePDF()
	if err != nil {
		t.Errorf("Failed to generate PDF: %s", err)
	}
	_, err = os.Stat("example.pdf")
	if os.IsNotExist(err) {
		t.Errorf("Failed to find output file: example.pdf")
	}
}
