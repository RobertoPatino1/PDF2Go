# md_pdf_2_Go

This project allows you to convert Markdown files to PDF using a backend server written in Go. The backend processes the Markdown content and returns a PDF file.

## Requirements

- Go >= 1.16
- Git
- Curl

## Initial Setup

**1. Clone the Repository**

```bash
git clone https://github.com/RobertoPatino1/md_pdf_2_Go.git
cd md_pdf_2_Go
```

**2. Install Dependencies**

```bash
go mod tidy
```

## How to use

Once the previous steps have been followed, make sure to follow these steps in order to use the tool.

**1. Run the server**

```bash
go run main.go
```

**2. Send POST request**

```bash
curl -X POST --data-binary @<YOUR_MARKDOWN_DOCUMENT>.md -H "Content-Type: text/plain" http://localhost:8080/convert -o <OUTPUT_PDF_FILENAME>.pdf
```
