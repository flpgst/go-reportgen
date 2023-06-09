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

func (c *SaveReportUseCase) Execute(input dto.ReportInputDTO) (dto.ReportOutputDTO, error) {
	report, err := entity.NewReport(input.ReportName, input.Date)
	if err != nil {
		return dto.ReportOutputDTO{}, err
	}
	if err := c.ReportRepository.Save(report); err != nil {
		return dto.ReportOutputDTO{}, err
	}
	dto := dto.ReportOutputDTO{
		ID:         report.ID,
		ReportName: report.ReportName,
		Date:       report.Date,
	}
	return dto, nil
}