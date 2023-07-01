package usecase

import (
	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
)

type SaveReportUseCase struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewSaveReportUseCase(reportRepository entity.ReportRepositoryInterface) *SaveReportUseCase {
	return &SaveReportUseCase{
		ReportRepository: reportRepository,
	}
}

func (c *SaveReportUseCase) Execute(input *dto.ReportDTO) (dto.ReportDTO, error) {
	report, err := entity.NewReport(input.ReportName, input.Date, input.Header, input.Body, input.Footer, entity.ReportTemplate(input.Template))
	if err != nil {
		return dto.ReportDTO{}, err
	}
	if err := c.ReportRepository.Save(report); err != nil {
		return dto.ReportDTO{}, err
	}
	dto := dto.ReportDTO{
		ReportName: report.ReportName,
		Date:       report.Date,
	}
	return dto, nil
}
