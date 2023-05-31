package usecase

import (
	"fmt"

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

func (r *GetReportUseCase) Execute(input dto.ReportInputDTO) (dto.ReportOutputDTO, error) {
	report, err := r.ReportRepository.GetReport(input.ReportName, input.Date)
	if err != nil {
		return dto.ReportOutputDTO{}, err
	}
	dto := dto.ReportOutputDTO{
		ReportName: report.ReportName,
		Date:       report.Date,
	}
	fmt.Println(dto)
	return dto, nil
}
