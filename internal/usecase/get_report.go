package usecase

import (
	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
)

type GetReportUseCase struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewGetReportUseCase(reportRepository entity.ReportRepositoryInterface) *GetReportUseCase {
	return &GetReportUseCase{
		ReportRepository: reportRepository,
	}
}

func (r *GetReportUseCase) Execute(input dto.ReportDTO) (*dto.ReportDTO, error) {
	report, err := r.ReportRepository.GetReport(input.ReportName, input.Date)
	if err != nil {
		return nil, err
	}
	dto := dto.ReportDTO{
		ReportName: report.ReportName,
		Date:       report.Date,
	}
	return &dto, nil
}
