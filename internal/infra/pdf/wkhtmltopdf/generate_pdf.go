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
	templateDir, err := filepath.Abs("../../internal/infra/pdf/template")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tmpl := template.New("")

	templateFiles := []string{
		filepath.Join(templateDir, "template.html"),
		filepath.Join(templateDir, "header/simple_header.html"),
		filepath.Join(templateDir, "body/simple_body.html"),
		filepath.Join(templateDir, "footer/simple_footer.html"),
	}

	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		fmt.Println("Error parsing template files:", err)
		return nil, err
	}

	htmlOutputFile, err := os.CreateTemp("", "output_*.html")
	if err != nil {
		fmt.Println("Error creating html output file:", err)
		return nil, err
	}
	defer os.Remove(htmlOutputFile.Name())

	data := struct {
		ReportName  string
		Date        string
		Header      []string
		Body        []string
		Footer      []string
		TemplateDir string
	}{
		ReportName:  dto.ReportName,
		Date:        dto.Date,
		Header:      dto.Header,
		Body:        dto.Body,
		Footer:      dto.Footer,
		TemplateDir: templateDir,
	}

	err = tmpl.ExecuteTemplate(htmlOutputFile, "template.html", data)
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

	cmd := exec.Command("wkhtmltopdf", "--enable-local-file-access", htmlOutputFile.Name(), pdfOutputFile.Name())
	errorLogFile, err := os.Create("error.log")
	if err != nil {
		fmt.Println("Error creating error log file:", err)
		return nil, err
	}
	defer errorLogFile.Close()

	cmd.Stderr = errorLogFile
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return nil, err
	}
	return pdfOutputFile, nil
}
