package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
)

type PDFReport struct {
	ReportName string
}

func NewPDFReport(reportName string) *PDFReport {
	return &PDFReport{
		ReportName: reportName,
	}
}

func (r *PDFReport) Execute() {
	templateFile := "template.html"
	pdfReport := NewPDFReport("ReportName Teste 1")

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	outputFile := "output.html"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating html output file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, pdfReport)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	pdfFile := "output.pdf"
	cmd := exec.Command("wkhtmltopdf", outputFile, pdfFile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}
	os.Remove(outputFile)
}
