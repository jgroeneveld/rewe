package rewebill

import (
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

type Pdf struct {
	TextPages []string
}

func (pdf Pdf) AllLines() []string {
	var lines []string
	for _, page := range pdf.TextPages {
		lines = append(lines, strings.Split(page, "\n")...)
	}

	return lines
}

func ReadPdf(rs io.ReadSeeker) (Pdf, error) {
	// disable stdout because the library prints stuff and we dont want to clutter stdout
	enableStdout := disableStdout()
	defer enableStdout()

	pdfReader, err := model.NewPdfReader(rs)
	if err != nil {
		return Pdf{}, errors.WithStack(err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return Pdf{}, errors.WithStack(err)
	}

	var textPages []string
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return Pdf{}, errors.WithStack(err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			return Pdf{}, errors.WithStack(err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			return Pdf{}, errors.WithStack(err)
		}

		textPages = append(textPages, text)
	}

	return Pdf{TextPages: textPages}, nil
}

func disableStdout() func() {
	_, w, _ := os.Pipe()

	oldStdout := os.Stdout
	os.Stdout = w

	return func() {
		w.Close()
		os.Stdout = oldStdout
	}
}
