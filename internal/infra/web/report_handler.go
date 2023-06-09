package web

import (
	"io"
	"net/http"

	"github.com/flpgst/go-reportgen/internal/dto"
	"github.com/flpgst/go-reportgen/internal/entity"
	"github.com/flpgst/go-reportgen/internal/infra/pdf"
	"github.com/flpgst/go-reportgen/internal/usecase"
	"github.com/google/uuid"
)

type WebReportHandler struct {
	ReportRepository entity.ReportRepositoryInterface
	PDFBuilder       pdf.PDFBuilderInterface
}

func NewWebReportHandler(
	ReportRepository entity.ReportRepositoryInterface,
	PDFBuilder pdf.PDFBuilderInterface,
) *WebReportHandler {
	return &WebReportHandler{
		ReportRepository: ReportRepository,
		PDFBuilder:       PDFBuilder,
	}
}

func (h *WebReportHandler) Get(w http.ResponseWriter, r *http.Request) {
	reportName := r.URL.Query().Get("reportName")
	date := r.URL.Query().Get("date")
	reportDTO := dto.ReportInputDTO{
		ReportName: reportName,
		Date:       date,
	}

	reportOutput, err := usecase.NewGetReportUseCase(h.ReportRepository).Execute(reportDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pdfFile, err := usecase.NewGeneratePDFUseCase(h.PDFBuilder).Execute(reportOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer pdfFile.Close()

	uniqueID := uuid.New()
	filename := uniqueID.String() + ".pdf"

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	_, err = io.Copy(w, pdfFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
