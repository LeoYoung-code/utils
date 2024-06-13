package pdf

import (
    "bytes"
    "io"
    "io/ioutil"
    "os"
    "testing"

    "github.com/jung-kurt/gofpdf"
)

// Creating a test helper function to mock the pdf creation
func createPDF() *gofpdf.Fpdf {
	pdfOut := gofpdf.New("P", "mm", "A4", "")
	pdfOut.AddPage()
	pdfOut.SetFont("Arial", "B", 16)
	pdfOut.Cell(40, 10, "Hello, world")
	return pdfOut
}

func TestMain(m *testing.M) {
	// Redirecting the output to a temp file
	tempFile, _ := ioutil.TempFile("/tmp", "output")
	os.Stdout = tempFile

	// Running the main function
	mainFunc := func() {
		main()
	}

	// Defining the test cases
	tests := []struct {
		name     string
		function func()
	}{
		{
			name:     "TestMain",
			function: mainFunc,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Redirect the output to a buffer
			old := os.Stdout // keep backup of the real stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call function
			test.function()

			outC := make(chan string)
			// copy the output in a separate goroutine so printing can't block indefinitely
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()

			// back to normal state
			w.Close()
			os.Stdout = old // restoring the real stdout
			out := <-outC

			// Comparing the mock and the actual output
			pdfMock := createPDF()
			pdfOut, _ := ioutil.ReadFile(tempFile.Name())

			if !bytes.Equal(pdfOut, pdfMock) {
				t.Errorf("Unexpected output in main() = %v, want %v", out, pdfMock)
			}
		})
	}
}
