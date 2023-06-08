package web

import (
	"io"
	"net/http"

	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
	"github.com/flpgst/go-reportgen/internal/usecase"
	"github.com/google/uuid"
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

	output, err := usecase.NewGetReportUseCase(h.ReportRepository).Execute(reportDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer output.Close()
	uniqueID := uuid.New()
	filename := uniqueID.String() + ".pdf"

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	_, err = io.Copy(w, output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
