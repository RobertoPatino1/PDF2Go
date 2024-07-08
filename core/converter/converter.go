package converter

import (
	"bytes"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/russross/blackfriday/v2"
)

func ConvertMarkdownToPDF(markdown string) ([]byte, error) {
    html := blackfriday.Run([]byte(markdown))

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.SetFont("Arial", "", 12)
    pdf.AddPage()

    processHTML(string(html), pdf)

    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

func processHTML(html string, pdf *gofpdf.Fpdf) {
    lines := strings.Split(html, "\n")
    for _, line := range lines {
        processLine(line, pdf)
    }
}

func processLine(line string, pdf *gofpdf.Fpdf) {
    if strings.HasPrefix(line, "<h1>") {
        pdf.SetFont("Arial", "B", 16)
        pdf.Cell(0, 10, stripTags(line))
        pdf.Ln(12)
    } else if strings.HasPrefix(line, "<h2>") {
        pdf.SetFont("Arial", "B", 14)
        pdf.Cell(0, 10, stripTags(line))
        pdf.Ln(10)
    } else if strings.HasPrefix(line, "<p>") {
        pdf.SetFont("Arial", "", 12)
        pdf.MultiCell(0, 10, stripTags(line), "", "L", false)
        pdf.Ln(-1)
    } else if strings.HasPrefix(line, "<ul>") {
        pdf.SetFont("Arial", "", 12)
        pdf.Cell(0, 10, stripTags(line))
        pdf.Ln(10)
    }
}

func stripTags(input string) string {
    input = strings.ReplaceAll(input, "<h1>", "")
    input = strings.ReplaceAll(input, "</h1>", "")
    input = strings.ReplaceAll(input, "<h2>", "")
    input = strings.ReplaceAll(input, "</h2>", "")
    input = strings.ReplaceAll(input, "<p>", "")
    input = strings.ReplaceAll(input, "</p>", "")
    input = strings.ReplaceAll(input, "<ul>", "")
    input = strings.ReplaceAll(input, "</ul>", "")
    return input
}
