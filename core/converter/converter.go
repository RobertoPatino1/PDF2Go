package converter

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"net/http"
	"os"

	pdf "github.com/stephenafamo/goldmark-pdf"
	"github.com/yuin/goldmark"
)
const font = "Noto Sans";
type httpFS struct {
    fs fs.FS
}

type httpFile struct {
    fs.File
}

func (h *httpFS) Open(name string) (http.File, error) {
    file, err := h.fs.Open(name)
    if err != nil {
        return nil, err
    }
    return &httpFile{file}, nil
}

func (f *httpFile) Readdir(count int) ([]fs.FileInfo, error) {
    return nil, fs.ErrInvalid
}

func (f *httpFile) Seek(offset int64, whence int) (int64, error) {
    if seeker, ok := f.File.(io.Seeker); ok {
        return seeker.Seek(offset, whence)
    }
    return 0, fs.ErrInvalid
}

func ConvertMarkdownToPDF(markdown string) ([]byte, error) {
    md := goldmark.New(
        goldmark.WithRenderer(
            pdf.New(
                pdf.WithTraceWriter(os.Stdout),
                pdf.WithContext(context.Background()),
                pdf.WithImageFS(&httpFS{os.DirFS(".")}),
                pdf.WithHeadingFont(pdf.GetTextFont(font, pdf.FontCourier)),
                pdf.WithBodyFont(pdf.GetTextFont(font, pdf.FontCourier)),
                pdf.WithCodeFont(pdf.GetCodeFont(font, pdf.FontCourier)),
            ),
        ),
    )

    var buf bytes.Buffer
    if err := md.Convert([]byte(markdown), &buf); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}
