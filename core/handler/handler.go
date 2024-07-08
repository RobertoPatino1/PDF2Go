package handler

import (
	"io"
	"net/http"

	"github.com/RobertoPatino1/md_pdf_2_Go/core/converter"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error while reading request body", http.StatusBadRequest)
        return
    }

    pdfBytes, err := converter.ConvertMarkdownToPDF(string(body))
    if err != nil {
        http.Error(w, "Error during conversion of Markdown to PDF", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Write(pdfBytes)
}
