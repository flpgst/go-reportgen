package usecase

import (
	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
)

type CreateReportUseCase struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewCreateReportUseCase(reportRepository entity.ReportRepositoryInterface) *CreateReportUseCase {
	return &CreateReportUseCase{
		ReportRepository: reportRepository,
	}
}

func (c *CreateReportUseCase) Execute(input dto.ReportInputDTO) (dto.ReportOutputDTO, error) {
	report, err := entity.NewReport(input.ReportName)
	if err != nil {
		return dto.ReportOutputDTO{}, err
	}
	if err := c.ReportRepository.Save(report); err != nil {
		return dto.ReportOutputDTO{}, err
	}
	dto := dto.ReportOutputDTO{
		ID:         report.ID,
		ReportName: report.ReportName,
	}
	return dto, nil
}
