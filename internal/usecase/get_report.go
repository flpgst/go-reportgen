package usecase

import (
	"fmt"
	"os"

	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
	pdf "github.com/flpgst/go-reportgen/internal/infra/pdf/wkhtmltopdf"
)

type GetReportUseCase struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewGetReportUseCase(reportRepository entity.ReportRepositoryInterface) *GetReportUseCase {
	return &GetReportUseCase{
		ReportRepository: reportRepository,
	}
}

func (r *GetReportUseCase) Execute(input dto.ReportInputDTO) (*os.File, error) {
	report, err := r.ReportRepository.GetReport(input.ReportName, input.Date)
	if err != nil {
		return nil, err
	}
	dto := dto.ReportOutputDTO{
		ReportName: report.ReportName,
		Date:       report.Date,
	}
	pdfReport := pdf.NewPDFReport(dto.ReportName)
	pdfFile, err := pdfReport.Execute()
	if err != nil {
		return nil, err
	}
	fmt.Println(dto)
	return pdfFile, nil
}
