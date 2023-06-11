package usecase

import (
	"os"

	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/infra/pdf"
)

type GeneratePDFUseCase struct {
	PDFBuilder pdf.PDFBuilderInterface
}

func NewGeneratePDFUseCase(pdfBuilder pdf.PDFBuilderInterface) *GeneratePDFUseCase {
	return &GeneratePDFUseCase{
		PDFBuilder: pdfBuilder,
	}
}

func (r *GeneratePDFUseCase) Execute(dto *dto.ReportDTO) (*os.File, error) {
	pdfFile, err := r.PDFBuilder.GeneratePDF(dto)
	if err != nil {
		return nil, err
	}

	return pdfFile, nil
}
