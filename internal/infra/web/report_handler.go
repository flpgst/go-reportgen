package web

import (
	"encoding/json"
	"net/http"

	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
	"github.com/flpgst/go-reportgen/internal/usecase"
)

type WebReportHandler struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewWebReportHandler(
	ReportRepository entity.ReportRepositoryInterface,
) *WebReportHandler {
	return &WebReportHandler{
		ReportRepository: ReportRepository,
	}
}

func (h *WebReportHandler) Get(w http.ResponseWriter, r *http.Request) {

	reportName := r.URL.Query().Get("reportName")
	date := r.URL.Query().Get("date")
	reportDTO := dto.ReportInputDTO{
		ReportName: reportName,
		Date:       date,
	}
	reportDTO.ReportName = reportName
	output, err := usecase.NewGetReportUseCase(h.ReportRepository).Execute(reportDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
