package pdf

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/flpgst/go-reportgen/internal/dto"
)

type WKHTMLTOPDF struct {
}

func NewWKHTMLTOPDF() *WKHTMLTOPDF {
	return &WKHTMLTOPDF{}
}

func (wk *WKHTMLTOPDF) GeneratePDF(dto *dto.ReportDTO) (*os.File, error) {
	templateAbsPath, err := filepath.Abs("../../internal/infra/pdf/template.html")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tmpl, err := template.ParseFiles(templateAbsPath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil, err
	}

	htmlOutputFile, err := os.CreateTemp("", "output_*.html")
	if err != nil {
		fmt.Println("Error creating html output file:", err)
		return nil, err
	}
	defer os.Remove(htmlOutputFile.Name())

	err = tmpl.Execute(htmlOutputFile, dto)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return nil, err
	}

	pdfOutputFile, err := os.CreateTemp("", "output_*.pdf")
	if err != nil {
		fmt.Println("Error creating pdf output file:", err)
		return nil, err
	}
	defer os.Remove(pdfOutputFile.Name())

	cmd := exec.Command("wkhtmltopdf", htmlOutputFile.Name(), pdfOutputFile.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return nil, err
	}
	return pdfOutputFile, nil
}
